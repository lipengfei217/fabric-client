// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto/geneses/result.proto

package geneses

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type String struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa888651797a6f10, []int{0}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CMDOut struct {
	Line                 int32    `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Prints               []string `protobuf:"bytes,2,rep,name=prints,proto3" json:"prints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CMDOut) Reset()         { *m = CMDOut{} }
func (m *CMDOut) String() string { return proto.CompactTextString(m) }
func (*CMDOut) ProtoMessage()    {}
func (*CMDOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa888651797a6f10, []int{1}
}

func (m *CMDOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMDOut.Unmarshal(m, b)
}
func (m *CMDOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMDOut.Marshal(b, m, deterministic)
}
func (m *CMDOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMDOut.Merge(m, src)
}
func (m *CMDOut) XXX_Size() int {
	return xxx_messageInfo_CMDOut.Size(m)
}
func (m *CMDOut) XXX_DiscardUnknown() {
	xxx_messageInfo_CMDOut.DiscardUnknown(m)
}

var xxx_messageInfo_CMDOut proto.InternalMessageInfo

func (m *CMDOut) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *CMDOut) GetPrints() []string {
	if m != nil {
		return m.Prints
	}
	return nil
}

func init() {
	proto.RegisterType((*String)(nil), "geneses.String")
	proto.RegisterType((*CMDOut)(nil), "geneses.CMDOut")
}

func init() { proto.RegisterFile("grpc/proto/geneses/result.proto", fileDescriptor_fa888651797a6f10) }

var fileDescriptor_fa888651797a6f10 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2f, 0x2a, 0x48,
	0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0x4e, 0x2d, 0xd6, 0x2f,
	0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0xd1, 0x03, 0x0b, 0x0a, 0xb1, 0x43, 0x45, 0x95, 0x64, 0xb8, 0xd8,
	0x82, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x85, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x13, 0x2e, 0x36, 0x67, 0x5f, 0x17, 0xff, 0xd2,
	0x12, 0x90, 0x6c, 0x4e, 0x66, 0x5e, 0x2a, 0x58, 0x96, 0x35, 0x08, 0xcc, 0x16, 0x12, 0xe3, 0x62,
	0x2b, 0x28, 0xca, 0xcc, 0x2b, 0x29, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0xf2, 0x92,
	0xd8, 0xc0, 0x76, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x97, 0xdb, 0xf7, 0x86, 0x00,
	0x00, 0x00,
}
