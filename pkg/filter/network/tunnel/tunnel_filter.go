/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tunnel

import (
	"context"

	"mosn.io/api"
	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/log"
	"mosn.io/mosn/pkg/network"
	"mosn.io/mosn/pkg/types"
	"mosn.io/mosn/pkg/upstream/cluster"
	"mosn.io/mosn/pkg/upstream/tunnel"
	"mosn.io/mosn/pkg/upstream/tunnel/ext"
)

var _ api.ReadFilter = (*tunnelFilter)(nil)

// tunnelFilter handles the reverse connection sent by the edge node，
// mainly used in scenarios where cloud nodes and edge nodes may be located in separate and isolated network region.
type tunnelFilter struct {
	clusterManager types.ClusterManager
	readCallbacks  api.ReadFilterCallbacks
	// connInitialized is used as a flag to be processed only once
	connInitialized bool
}

func (t *tunnelFilter) OnData(buffer api.IoBuffer) api.FilterStatus {
	if t.connInitialized {
		// hand it over to the next filter directly
		return api.Continue
	}
	if log.DefaultLogger.GetLogLevel() >= log.DEBUG {
		log.DefaultLogger.Debugf("[tunnel server] [ondata] read data , len: %v", buffer.Len())
	}
	data, err := tunnel.DecodeFromBuffer(buffer)
	conn := t.readCallbacks.Connection()
	if data == nil && err == nil {
		// Not enough data was read.
		return api.Stop
	}
	if err != nil {
		log.DefaultLogger.Errorf("[tunnel server] [ondata] failed to read from buffer, close connection err: %+v", err)
		writeConnectResponse(tunnel.ConnectUnknownFailed, conn)
		return api.Stop
	}
	// Now it can only be ConnectionInitInfo
	info, ok := data.(tunnel.ConnectionInitInfo)
	if !ok {
		log.DefaultLogger.Errorf("[tunnel server] [ondata] decode failed, data error")
		writeConnectResponse(tunnel.ConnectUnknownFailed, conn)
		return api.Stop
	}
	// Auth the connection
	if info.CredentialPolicy != "" {
		validator := ext.GetConnectionValidator(info.CredentialPolicy)
		if validator == nil {
			writeConnectResponse(tunnel.ConnectValidatorNotFound, conn)
			return api.Stop
		}
		res := validator.Validate(info.Credential, info.HostName, info.ClusterName)
		if !res {
			writeConnectResponse(tunnel.ConnectAuthFailed, conn)
			return api.Stop
		}
	}
	if !t.clusterManager.ClusterExist(info.ClusterName) {
		writeConnectResponse(tunnel.ConnectClusterNotExist, conn)
		return api.Stop
	}
	// Set the flag that has been initialized, subsequent data processing skips this filter
	err = writeConnectResponse(tunnel.ConnectSuccess, conn)
	if err != nil {
		return api.Stop
	}
	conn.AddConnectionEventListener(NewHostRemover(conn.RemoteAddr().String(), info.ClusterName))
	tunnelHostMutex.Lock()
	defer tunnelHostMutex.Unlock()
	snapshot := t.clusterManager.GetClusterSnapshot(context.Background(), info.ClusterName)
	_ = t.clusterManager.AppendClusterTypesHosts(info.ClusterName, []types.Host{cluster.NewTunnelHost(v2.Host{
		HostConfig: v2.HostConfig{
			Address:    conn.RemoteAddr().String(),
			Hostname:   info.HostName,
			Weight:     uint32(info.Weight),
			TLSDisable: false,
		}}, snapshot.ClusterInfo(), network.CreateTunnelAgentConnection(conn))})
	t.connInitialized = true
	return api.Stop
}

func (t *tunnelFilter) OnNewConnection() api.FilterStatus {
	return api.Continue
}

func (t *tunnelFilter) InitializeReadFilterCallbacks(cb api.ReadFilterCallbacks) {
	t.readCallbacks = cb
}

func writeConnectResponse(status tunnel.ConnectStatus, conn api.Connection) error {
	log.DefaultLogger.Debugf("[tunnel server] try to write response, connect status: %v", status)
	buffer, err := tunnel.Encode(&tunnel.ConnectionInitResponse{Status: status})
	if err != nil {
		log.DefaultLogger.Errorf("[tunnel server] failed to encode response, err: %+v", err)
		conn.Close(api.NoFlush, api.LocalClose)
		return err
	}
	err = conn.Write(buffer)
	if err != nil {
		log.DefaultLogger.Errorf("[tunnel server] failed to write response, err: %+v", err)
		conn.Close(api.NoFlush, api.OnWriteErrClose)
		return err
	}
	return nil
}