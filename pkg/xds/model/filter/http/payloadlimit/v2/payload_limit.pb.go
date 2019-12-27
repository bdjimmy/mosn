// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payload_limit.proto

package v2

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PayloadLimit struct {
	MaxEntitySize int32 `protobuf:"varint,1,opt,name=max_entity_size,json=maxEntitySize,proto3" json:"max_entity_size,omitempty"`
	// Types that are valid to be assigned to ErrorType:
	//	*PayloadLimit_HttpStatus
	ErrorType            isPayloadLimit_ErrorType `protobuf_oneof:"error_type"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PayloadLimit) Reset()         { *m = PayloadLimit{} }
func (m *PayloadLimit) String() string { return proto.CompactTextString(m) }
func (*PayloadLimit) ProtoMessage()    {}
func (*PayloadLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcec3c28efae8940, []int{0}
}
func (m *PayloadLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadLimit.Unmarshal(m, b)
}
func (m *PayloadLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadLimit.Marshal(b, m, deterministic)
}
func (m *PayloadLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadLimit.Merge(m, src)
}
func (m *PayloadLimit) XXX_Size() int {
	return xxx_messageInfo_PayloadLimit.Size(m)
}
func (m *PayloadLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadLimit.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadLimit proto.InternalMessageInfo

type isPayloadLimit_ErrorType interface {
	isPayloadLimit_ErrorType()
}

type PayloadLimit_HttpStatus struct {
	HttpStatus int32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*PayloadLimit_HttpStatus) isPayloadLimit_ErrorType() {}

func (m *PayloadLimit) GetErrorType() isPayloadLimit_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *PayloadLimit) GetMaxEntitySize() int32 {
	if m != nil {
		return m.MaxEntitySize
	}
	return 0
}

func (m *PayloadLimit) GetHttpStatus() int32 {
	if x, ok := m.GetErrorType().(*PayloadLimit_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PayloadLimit) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PayloadLimit_OneofMarshaler, _PayloadLimit_OneofUnmarshaler, _PayloadLimit_OneofSizer, []interface{}{
		(*PayloadLimit_HttpStatus)(nil),
	}
}

func _PayloadLimit_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PayloadLimit)
	// error_type
	switch x := m.ErrorType.(type) {
	case *PayloadLimit_HttpStatus:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		return fmt.Errorf("PayloadLimit.ErrorType has unexpected type %T", x)
	}
	return nil
}

func _PayloadLimit_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PayloadLimit)
	switch tag {
	case 2: // error_type.http_status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ErrorType = &PayloadLimit_HttpStatus{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _PayloadLimit_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PayloadLimit)
	// error_type
	switch x := m.ErrorType.(type) {
	case *PayloadLimit_HttpStatus:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*PayloadLimit)(nil), "envoy.config.filter.http.payloadlimit.v2.PayloadLimit")
}

func init() { proto.RegisterFile("payload_limit.proto", fileDescriptor_fcec3c28efae8940) }

var fileDescriptor_fcec3c28efae8940 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0xac, 0xcc,
	0xc9, 0x4f, 0x4c, 0x89, 0xcf, 0xc9, 0xcc, 0xcd, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xd2, 0x48, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4, 0x4b, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4b, 0xcb,
	0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xcb, 0x28, 0x29, 0x29, 0xd0, 0x83, 0xaa, 0x86, 0x28, 0x2e, 0x33,
	0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x46, 0x28,
	0xd5, 0x72, 0xf1, 0x04, 0x40, 0xd4, 0xfa, 0x80, 0xd4, 0x0a, 0xa9, 0x71, 0xf1, 0xe7, 0x26, 0x56,
	0xc4, 0xa7, 0xe6, 0x95, 0x64, 0x96, 0x54, 0xc6, 0x17, 0x67, 0x56, 0xa5, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x06, 0xf1, 0xe6, 0x26, 0x56, 0xb8, 0x82, 0x45, 0x83, 0x33, 0xab, 0x52, 0x85, 0x0c,
	0xb8, 0xb8, 0x41, 0x16, 0xc5, 0x17, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b, 0x30, 0x81, 0xd4, 0x38,
	0xf1, 0xee, 0x7a, 0x79, 0x80, 0x99, 0x43, 0x8a, 0x4d, 0xe0, 0x06, 0x8b, 0xc6, 0x09, 0x46, 0x0f,
	0x86, 0x20, 0x2e, 0x90, 0x9a, 0x60, 0xb0, 0x12, 0x27, 0x61, 0x2e, 0xae, 0xd4, 0xa2, 0xa2, 0xfc,
	0xa2, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x21, 0xd6, 0x1d, 0x2f, 0x0f, 0x30, 0x33, 0x3a, 0x05, 0x72,
	0x99, 0x65, 0xe6, 0xeb, 0x81, 0xbd, 0x51, 0x50, 0x94, 0x5f, 0x51, 0xa9, 0x47, 0xac, 0x8f, 0x9c,
	0x04, 0x91, 0x9d, 0x1d, 0x00, 0xf2, 0x4b, 0x00, 0x63, 0x14, 0x53, 0x99, 0x51, 0x12, 0x1b, 0xd8,
	0x63, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xe7, 0x3a, 0xd5, 0x32, 0x01, 0x00, 0x00,
}
