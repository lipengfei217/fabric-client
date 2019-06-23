// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto/geneses/conf.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Generate 请求生成区块链配置对象
type Generate struct {
	LedgerName           string   `protobuf:"bytes,1,opt,name=ledgerName,proto3" json:"ledgerName,omitempty"`
	OrderCount           int32    `protobuf:"varint,2,opt,name=orderCount,proto3" json:"orderCount,omitempty"`
	PeerCount            int32    `protobuf:"varint,3,opt,name=peerCount,proto3" json:"peerCount,omitempty"`
	TemplateCount        int32    `protobuf:"varint,4,opt,name=templateCount,proto3" json:"templateCount,omitempty"`
	UserCount            int32    `protobuf:"varint,5,opt,name=userCount,proto3" json:"userCount,omitempty"`
	BatchTimeout         int32    `protobuf:"varint,6,opt,name=batchTimeout,proto3" json:"batchTimeout,omitempty"`
	MaxMessageCount      int32    `protobuf:"varint,7,opt,name=maxMessageCount,proto3" json:"maxMessageCount,omitempty"`
	Force                bool     `protobuf:"varint,8,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Generate) Reset()         { *m = Generate{} }
func (m *Generate) String() string { return proto.CompactTextString(m) }
func (*Generate) ProtoMessage()    {}
func (*Generate) Descriptor() ([]byte, []int) {
	return fileDescriptor_104dc51ee168266a, []int{0}
}

func (m *Generate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Generate.Unmarshal(m, b)
}
func (m *Generate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Generate.Marshal(b, m, deterministic)
}
func (m *Generate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Generate.Merge(m, src)
}
func (m *Generate) XXX_Size() int {
	return xxx_messageInfo_Generate.Size(m)
}
func (m *Generate) XXX_DiscardUnknown() {
	xxx_messageInfo_Generate.DiscardUnknown(m)
}

var xxx_messageInfo_Generate proto.InternalMessageInfo

func (m *Generate) GetLedgerName() string {
	if m != nil {
		return m.LedgerName
	}
	return ""
}

func (m *Generate) GetOrderCount() int32 {
	if m != nil {
		return m.OrderCount
	}
	return 0
}

func (m *Generate) GetPeerCount() int32 {
	if m != nil {
		return m.PeerCount
	}
	return 0
}

func (m *Generate) GetTemplateCount() int32 {
	if m != nil {
		return m.TemplateCount
	}
	return 0
}

func (m *Generate) GetUserCount() int32 {
	if m != nil {
		return m.UserCount
	}
	return 0
}

func (m *Generate) GetBatchTimeout() int32 {
	if m != nil {
		return m.BatchTimeout
	}
	return 0
}

func (m *Generate) GetMaxMessageCount() int32 {
	if m != nil {
		return m.MaxMessageCount
	}
	return 0
}

func (m *Generate) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

// Crypto 请求生成区块链配置文件集合对象
type Crypto struct {
	LedgerName           string   `protobuf:"bytes,1,opt,name=ledgerName,proto3" json:"ledgerName,omitempty"`
	Force                bool     `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crypto) Reset()         { *m = Crypto{} }
func (m *Crypto) String() string { return proto.CompactTextString(m) }
func (*Crypto) ProtoMessage()    {}
func (*Crypto) Descriptor() ([]byte, []int) {
	return fileDescriptor_104dc51ee168266a, []int{1}
}

func (m *Crypto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crypto.Unmarshal(m, b)
}
func (m *Crypto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crypto.Marshal(b, m, deterministic)
}
func (m *Crypto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crypto.Merge(m, src)
}
func (m *Crypto) XXX_Size() int {
	return xxx_messageInfo_Crypto.Size(m)
}
func (m *Crypto) XXX_DiscardUnknown() {
	xxx_messageInfo_Crypto.DiscardUnknown(m)
}

var xxx_messageInfo_Crypto proto.InternalMessageInfo

func (m *Crypto) GetLedgerName() string {
	if m != nil {
		return m.LedgerName
	}
	return ""
}

func (m *Crypto) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

// ChannelTX 请求生成区块链配置文件集合对象
type ChannelTX struct {
	LedgerName           string   `protobuf:"bytes,1,opt,name=ledgerName,proto3" json:"ledgerName,omitempty"`
	ChannelName          string   `protobuf:"bytes,2,opt,name=channelName,proto3" json:"channelName,omitempty"`
	Force                bool     `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelTX) Reset()         { *m = ChannelTX{} }
func (m *ChannelTX) String() string { return proto.CompactTextString(m) }
func (*ChannelTX) ProtoMessage()    {}
func (*ChannelTX) Descriptor() ([]byte, []int) {
	return fileDescriptor_104dc51ee168266a, []int{2}
}

func (m *ChannelTX) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelTX.Unmarshal(m, b)
}
func (m *ChannelTX) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelTX.Marshal(b, m, deterministic)
}
func (m *ChannelTX) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelTX.Merge(m, src)
}
func (m *ChannelTX) XXX_Size() int {
	return xxx_messageInfo_ChannelTX.Size(m)
}
func (m *ChannelTX) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelTX.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelTX proto.InternalMessageInfo

func (m *ChannelTX) GetLedgerName() string {
	if m != nil {
		return m.LedgerName
	}
	return ""
}

func (m *ChannelTX) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func (m *ChannelTX) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func init() {
	proto.RegisterType((*Generate)(nil), "geneses.Generate")
	proto.RegisterType((*Crypto)(nil), "geneses.Crypto")
	proto.RegisterType((*ChannelTX)(nil), "geneses.ChannelTX")
}

func init() { proto.RegisterFile("grpc/proto/geneses/conf.proto", fileDescriptor_104dc51ee168266a) }

var fileDescriptor_104dc51ee168266a = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3d, 0x4f, 0xc3, 0x30,
	0x18, 0x84, 0x95, 0x94, 0xa4, 0xc9, 0x0b, 0x08, 0xc9, 0x62, 0xc8, 0x00, 0x28, 0x8a, 0x18, 0x32,
	0x91, 0x81, 0x9d, 0x25, 0x03, 0x13, 0x0c, 0x51, 0x07, 0x56, 0xd7, 0xbd, 0xa6, 0x48, 0x89, 0x1d,
	0xd9, 0x8e, 0x04, 0xff, 0x82, 0x9f, 0x8c, 0xea, 0xf4, 0xc3, 0xed, 0xd2, 0x31, 0xcf, 0x3d, 0xef,
	0x9d, 0x14, 0xd3, 0x63, 0xab, 0x07, 0x51, 0x0d, 0x5a, 0x59, 0x55, 0xb5, 0x90, 0x30, 0x30, 0x95,
	0x50, 0x72, 0xfd, 0xe2, 0x10, 0x9b, 0xef, 0x58, 0xf1, 0x17, 0x52, 0xf2, 0x0e, 0x09, 0xcd, 0x2d,
	0xd8, 0x13, 0x51, 0x87, 0x55, 0x0b, 0xfd, 0xc9, 0x7b, 0x64, 0x41, 0x1e, 0x94, 0x69, 0xe3, 0x91,
	0x6d, 0xae, 0xf4, 0x0a, 0xba, 0x56, 0xa3, 0xb4, 0x59, 0x98, 0x07, 0x65, 0xd4, 0x78, 0x84, 0x3d,
	0x50, 0x3a, 0x60, 0x1f, 0xcf, 0x5c, 0x7c, 0x04, 0xec, 0x99, 0x6e, 0x2d, 0xfa, 0xa1, 0xe3, 0x16,
	0x93, 0x71, 0xe5, 0x8c, 0x53, 0xb8, 0xed, 0x18, 0xcd, 0xbe, 0x23, 0x9a, 0x3a, 0x0e, 0x80, 0x15,
	0x74, 0xb3, 0xe4, 0x56, 0x6c, 0x16, 0xdf, 0x3d, 0xd4, 0x68, 0xb3, 0xd8, 0x09, 0x27, 0x8c, 0x95,
	0x74, 0xd7, 0xf3, 0x9f, 0x0f, 0x18, 0xc3, 0xdb, 0xdd, 0xd2, 0xdc, 0x69, 0xe7, 0x98, 0xdd, 0x53,
	0xb4, 0x56, 0x5a, 0x20, 0x4b, 0xf2, 0xa0, 0x4c, 0x9a, 0xe9, 0xa3, 0x78, 0xa3, 0xb8, 0xd6, 0xbf,
	0x83, 0x55, 0x17, 0xff, 0xc7, 0xe1, 0x3e, 0xf4, 0xef, 0x05, 0xa5, 0xf5, 0x86, 0x4b, 0x89, 0x6e,
	0xf1, 0x75, 0xb1, 0x22, 0xa7, 0x6b, 0x31, 0xc9, 0x4e, 0x08, 0x9d, 0xe0, 0xa3, 0xe3, 0xc8, 0xcc,
	0x1b, 0x59, 0xc6, 0xee, 0x1d, 0x5f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x60, 0x45, 0xf0,
	0xe8, 0x01, 0x00, 0x00,
}
