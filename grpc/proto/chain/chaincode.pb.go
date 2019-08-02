// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto/chain/chaincode.proto

package chain

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

type ChainCodeInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Input                string   `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	Escc                 string   `protobuf:"bytes,5,opt,name=escc,proto3" json:"escc,omitempty"`
	Vscc                 string   `protobuf:"bytes,6,opt,name=vscc,proto3" json:"vscc,omitempty"`
	Id                   []byte   `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainCodeInfo) Reset()         { *m = ChainCodeInfo{} }
func (m *ChainCodeInfo) String() string { return proto.CompactTextString(m) }
func (*ChainCodeInfo) ProtoMessage()    {}
func (*ChainCodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{0}
}

func (m *ChainCodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainCodeInfo.Unmarshal(m, b)
}
func (m *ChainCodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainCodeInfo.Marshal(b, m, deterministic)
}
func (m *ChainCodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainCodeInfo.Merge(m, src)
}
func (m *ChainCodeInfo) XXX_Size() int {
	return xxx_messageInfo_ChainCodeInfo.Size(m)
}
func (m *ChainCodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainCodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChainCodeInfo proto.InternalMessageInfo

func (m *ChainCodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChainCodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChainCodeInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChainCodeInfo) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ChainCodeInfo) GetEscc() string {
	if m != nil {
		return m.Escc
	}
	return ""
}

func (m *ChainCodeInfo) GetVscc() string {
	if m != nil {
		return m.Vscc
	}
	return ""
}

func (m *ChainCodeInfo) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type CCList struct {
	Data                 []*ChainCodeInfo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CCList) Reset()         { *m = CCList{} }
func (m *CCList) String() string { return proto.CompactTextString(m) }
func (*CCList) ProtoMessage()    {}
func (*CCList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{1}
}

func (m *CCList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CCList.Unmarshal(m, b)
}
func (m *CCList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CCList.Marshal(b, m, deterministic)
}
func (m *CCList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CCList.Merge(m, src)
}
func (m *CCList) XXX_Size() int {
	return xxx_messageInfo_CCList.Size(m)
}
func (m *CCList) XXX_DiscardUnknown() {
	xxx_messageInfo_CCList.DiscardUnknown(m)
}

var xxx_messageInfo_CCList proto.InternalMessageInfo

func (m *CCList) GetData() []*ChainCodeInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type Upload struct {
	LedgerName           string   `protobuf:"bytes,1,opt,name=ledgerName,proto3" json:"ledgerName,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Data                 []byte   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Upload) Reset()         { *m = Upload{} }
func (m *Upload) String() string { return proto.CompactTextString(m) }
func (*Upload) ProtoMessage()    {}
func (*Upload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{2}
}

func (m *Upload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Upload.Unmarshal(m, b)
}
func (m *Upload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Upload.Marshal(b, m, deterministic)
}
func (m *Upload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upload.Merge(m, src)
}
func (m *Upload) XXX_Size() int {
	return xxx_messageInfo_Upload.Size(m)
}
func (m *Upload) XXX_DiscardUnknown() {
	xxx_messageInfo_Upload.DiscardUnknown(m)
}

var xxx_messageInfo_Upload proto.InternalMessageInfo

func (m *Upload) GetLedgerName() string {
	if m != nil {
		return m.LedgerName
	}
	return ""
}

func (m *Upload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Upload) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Upload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Install struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,3,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Source               string   `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Version              string   `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	PeerName             string   `protobuf:"bytes,8,opt,name=peerName,proto3" json:"peerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{3}
}

func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (m *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(m, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

func (m *Install) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *Install) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *Install) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *Install) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Install) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Install) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Install) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Install) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

type Installed struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,3,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	PeerName             string   `protobuf:"bytes,4,opt,name=peerName,proto3" json:"peerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Installed) Reset()         { *m = Installed{} }
func (m *Installed) String() string { return proto.CompactTextString(m) }
func (*Installed) ProtoMessage()    {}
func (*Installed) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{4}
}

func (m *Installed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Installed.Unmarshal(m, b)
}
func (m *Installed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Installed.Marshal(b, m, deterministic)
}
func (m *Installed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Installed.Merge(m, src)
}
func (m *Installed) XXX_Size() int {
	return xxx_messageInfo_Installed.Size(m)
}
func (m *Installed) XXX_DiscardUnknown() {
	xxx_messageInfo_Installed.DiscardUnknown(m)
}

var xxx_messageInfo_Installed proto.InternalMessageInfo

func (m *Installed) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *Installed) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *Installed) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *Installed) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

type Instantiate struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,3,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	ChannelID            string   `protobuf:"bytes,4,opt,name=channelID,proto3" json:"channelID,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Version              string   `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	OrgPolicies          []string `protobuf:"bytes,8,rep,name=orgPolicies,proto3" json:"orgPolicies,omitempty"`
	Args                 [][]byte `protobuf:"bytes,9,rep,name=args,proto3" json:"args,omitempty"`
	PeerName             string   `protobuf:"bytes,10,opt,name=peerName,proto3" json:"peerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instantiate) Reset()         { *m = Instantiate{} }
func (m *Instantiate) String() string { return proto.CompactTextString(m) }
func (*Instantiate) ProtoMessage()    {}
func (*Instantiate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{5}
}

func (m *Instantiate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instantiate.Unmarshal(m, b)
}
func (m *Instantiate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instantiate.Marshal(b, m, deterministic)
}
func (m *Instantiate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instantiate.Merge(m, src)
}
func (m *Instantiate) XXX_Size() int {
	return xxx_messageInfo_Instantiate.Size(m)
}
func (m *Instantiate) XXX_DiscardUnknown() {
	xxx_messageInfo_Instantiate.DiscardUnknown(m)
}

var xxx_messageInfo_Instantiate proto.InternalMessageInfo

func (m *Instantiate) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *Instantiate) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *Instantiate) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *Instantiate) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *Instantiate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instantiate) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Instantiate) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Instantiate) GetOrgPolicies() []string {
	if m != nil {
		return m.OrgPolicies
	}
	return nil
}

func (m *Instantiate) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Instantiate) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

type Instantiated struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,3,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	ChannelID            string   `protobuf:"bytes,4,opt,name=channelID,proto3" json:"channelID,omitempty"`
	PeerName             string   `protobuf:"bytes,5,opt,name=peerName,proto3" json:"peerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instantiated) Reset()         { *m = Instantiated{} }
func (m *Instantiated) String() string { return proto.CompactTextString(m) }
func (*Instantiated) ProtoMessage()    {}
func (*Instantiated) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{6}
}

func (m *Instantiated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instantiated.Unmarshal(m, b)
}
func (m *Instantiated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instantiated.Marshal(b, m, deterministic)
}
func (m *Instantiated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instantiated.Merge(m, src)
}
func (m *Instantiated) XXX_Size() int {
	return xxx_messageInfo_Instantiated.Size(m)
}
func (m *Instantiated) XXX_DiscardUnknown() {
	xxx_messageInfo_Instantiated.DiscardUnknown(m)
}

var xxx_messageInfo_Instantiated proto.InternalMessageInfo

func (m *Instantiated) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *Instantiated) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *Instantiated) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *Instantiated) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *Instantiated) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

type Upgrade struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	OrgName              string   `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,3,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	ChannelID            string   `protobuf:"bytes,4,opt,name=channelID,proto3" json:"channelID,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Version              string   `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	OrgPolicies          []string `protobuf:"bytes,8,rep,name=orgPolicies,proto3" json:"orgPolicies,omitempty"`
	Args                 [][]byte `protobuf:"bytes,9,rep,name=args,proto3" json:"args,omitempty"`
	PeerName             string   `protobuf:"bytes,10,opt,name=peerName,proto3" json:"peerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Upgrade) Reset()         { *m = Upgrade{} }
func (m *Upgrade) String() string { return proto.CompactTextString(m) }
func (*Upgrade) ProtoMessage()    {}
func (*Upgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{7}
}

func (m *Upgrade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Upgrade.Unmarshal(m, b)
}
func (m *Upgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Upgrade.Marshal(b, m, deterministic)
}
func (m *Upgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upgrade.Merge(m, src)
}
func (m *Upgrade) XXX_Size() int {
	return xxx_messageInfo_Upgrade.Size(m)
}
func (m *Upgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_Upgrade.DiscardUnknown(m)
}

var xxx_messageInfo_Upgrade proto.InternalMessageInfo

func (m *Upgrade) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *Upgrade) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *Upgrade) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *Upgrade) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *Upgrade) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Upgrade) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Upgrade) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Upgrade) GetOrgPolicies() []string {
	if m != nil {
		return m.OrgPolicies
	}
	return nil
}

func (m *Upgrade) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Upgrade) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

type Invoke struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	ChannelID            string   `protobuf:"bytes,2,opt,name=channelID,proto3" json:"channelID,omitempty"`
	ChainCodeID          string   `protobuf:"bytes,3,opt,name=chainCodeID,proto3" json:"chainCodeID,omitempty"`
	OrgName              string   `protobuf:"bytes,4,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,5,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	Fcn                  string   `protobuf:"bytes,6,opt,name=fcn,proto3" json:"fcn,omitempty"`
	Args                 [][]byte `protobuf:"bytes,7,rep,name=args,proto3" json:"args,omitempty"`
	TargetEndpoints      []string `protobuf:"bytes,8,rep,name=targetEndpoints,proto3" json:"targetEndpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invoke) Reset()         { *m = Invoke{} }
func (m *Invoke) String() string { return proto.CompactTextString(m) }
func (*Invoke) ProtoMessage()    {}
func (*Invoke) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{8}
}

func (m *Invoke) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoke.Unmarshal(m, b)
}
func (m *Invoke) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoke.Marshal(b, m, deterministic)
}
func (m *Invoke) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoke.Merge(m, src)
}
func (m *Invoke) XXX_Size() int {
	return xxx_messageInfo_Invoke.Size(m)
}
func (m *Invoke) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoke.DiscardUnknown(m)
}

var xxx_messageInfo_Invoke proto.InternalMessageInfo

func (m *Invoke) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *Invoke) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *Invoke) GetChainCodeID() string {
	if m != nil {
		return m.ChainCodeID
	}
	return ""
}

func (m *Invoke) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *Invoke) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *Invoke) GetFcn() string {
	if m != nil {
		return m.Fcn
	}
	return ""
}

func (m *Invoke) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Invoke) GetTargetEndpoints() []string {
	if m != nil {
		return m.TargetEndpoints
	}
	return nil
}

type InvokeAsync struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	ChannelID            string   `protobuf:"bytes,2,opt,name=channelID,proto3" json:"channelID,omitempty"`
	ChainCodeID          string   `protobuf:"bytes,3,opt,name=chainCodeID,proto3" json:"chainCodeID,omitempty"`
	OrgName              string   `protobuf:"bytes,4,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,5,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	Fcn                  string   `protobuf:"bytes,6,opt,name=fcn,proto3" json:"fcn,omitempty"`
	Callback             string   `protobuf:"bytes,7,opt,name=callback,proto3" json:"callback,omitempty"`
	Args                 [][]byte `protobuf:"bytes,8,rep,name=args,proto3" json:"args,omitempty"`
	TargetEndpoints      []string `protobuf:"bytes,9,rep,name=targetEndpoints,proto3" json:"targetEndpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeAsync) Reset()         { *m = InvokeAsync{} }
func (m *InvokeAsync) String() string { return proto.CompactTextString(m) }
func (*InvokeAsync) ProtoMessage()    {}
func (*InvokeAsync) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{9}
}

func (m *InvokeAsync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeAsync.Unmarshal(m, b)
}
func (m *InvokeAsync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeAsync.Marshal(b, m, deterministic)
}
func (m *InvokeAsync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeAsync.Merge(m, src)
}
func (m *InvokeAsync) XXX_Size() int {
	return xxx_messageInfo_InvokeAsync.Size(m)
}
func (m *InvokeAsync) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeAsync.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeAsync proto.InternalMessageInfo

func (m *InvokeAsync) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *InvokeAsync) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *InvokeAsync) GetChainCodeID() string {
	if m != nil {
		return m.ChainCodeID
	}
	return ""
}

func (m *InvokeAsync) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *InvokeAsync) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *InvokeAsync) GetFcn() string {
	if m != nil {
		return m.Fcn
	}
	return ""
}

func (m *InvokeAsync) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func (m *InvokeAsync) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *InvokeAsync) GetTargetEndpoints() []string {
	if m != nil {
		return m.TargetEndpoints
	}
	return nil
}

type Query struct {
	ConfigID             string   `protobuf:"bytes,1,opt,name=configID,proto3" json:"configID,omitempty"`
	ChannelID            string   `protobuf:"bytes,2,opt,name=channelID,proto3" json:"channelID,omitempty"`
	ChainCodeID          string   `protobuf:"bytes,3,opt,name=chainCodeID,proto3" json:"chainCodeID,omitempty"`
	OrgName              string   `protobuf:"bytes,4,opt,name=orgName,proto3" json:"orgName,omitempty"`
	OrgUser              string   `protobuf:"bytes,5,opt,name=orgUser,proto3" json:"orgUser,omitempty"`
	Fcn                  string   `protobuf:"bytes,6,opt,name=fcn,proto3" json:"fcn,omitempty"`
	Args                 [][]byte `protobuf:"bytes,7,rep,name=args,proto3" json:"args,omitempty"`
	TargetEndpoints      []string `protobuf:"bytes,8,rep,name=targetEndpoints,proto3" json:"targetEndpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_c776d1056ba94da0, []int{10}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetConfigID() string {
	if m != nil {
		return m.ConfigID
	}
	return ""
}

func (m *Query) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *Query) GetChainCodeID() string {
	if m != nil {
		return m.ChainCodeID
	}
	return ""
}

func (m *Query) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *Query) GetOrgUser() string {
	if m != nil {
		return m.OrgUser
	}
	return ""
}

func (m *Query) GetFcn() string {
	if m != nil {
		return m.Fcn
	}
	return ""
}

func (m *Query) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Query) GetTargetEndpoints() []string {
	if m != nil {
		return m.TargetEndpoints
	}
	return nil
}

func init() {
	proto.RegisterType((*ChainCodeInfo)(nil), "chain.ChainCodeInfo")
	proto.RegisterType((*CCList)(nil), "chain.CCList")
	proto.RegisterType((*Upload)(nil), "chain.Upload")
	proto.RegisterType((*Install)(nil), "chain.Install")
	proto.RegisterType((*Installed)(nil), "chain.Installed")
	proto.RegisterType((*Instantiate)(nil), "chain.Instantiate")
	proto.RegisterType((*Instantiated)(nil), "chain.Instantiated")
	proto.RegisterType((*Upgrade)(nil), "chain.Upgrade")
	proto.RegisterType((*Invoke)(nil), "chain.Invoke")
	proto.RegisterType((*InvokeAsync)(nil), "chain.InvokeAsync")
	proto.RegisterType((*Query)(nil), "chain.Query")
}

func init() { proto.RegisterFile("grpc/proto/chain/chaincode.proto", fileDescriptor_c776d1056ba94da0) }

var fileDescriptor_c776d1056ba94da0 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0x9a, 0x3f, 0xcd, 0xe9, 0xba, 0x2e, 0x61, 0x91, 0xb0, 0x88, 0x84, 0xdc, 0x98,
	0x1b, 0xb3, 0xb0, 0x3e, 0x81, 0xdb, 0x7a, 0x51, 0x94, 0x65, 0x2d, 0xf4, 0xc6, 0xbb, 0xe9, 0x64,
	0x9a, 0x8e, 0x1b, 0x67, 0xc2, 0x64, 0x5a, 0xd9, 0x47, 0x10, 0x2f, 0xbc, 0xf5, 0xce, 0x47, 0xf2,
	0x31, 0xf4, 0x31, 0x64, 0x26, 0x69, 0x3a, 0x8d, 0xeb, 0x82, 0xb0, 0x08, 0xca, 0xde, 0x94, 0x73,
	0xbe, 0x1e, 0x72, 0xbe, 0xef, 0x37, 0x69, 0x13, 0x88, 0x0b, 0x51, 0xe1, 0xd3, 0x4a, 0x70, 0xc9,
	0x4f, 0xf1, 0x0a, 0x51, 0xd6, 0x7c, 0x62, 0x9e, 0x93, 0x4c, 0xab, 0xa1, 0xab, 0x85, 0xe4, 0xab,
	0x05, 0x0f, 0xc6, 0xaa, 0x1a, 0xf3, 0x9c, 0x4c, 0xd9, 0x92, 0x87, 0x21, 0x38, 0x0c, 0xbd, 0x27,
	0x91, 0x15, 0x5b, 0x69, 0x30, 0xd3, 0x75, 0x18, 0x81, 0xbf, 0x21, 0xa2, 0xa6, 0x9c, 0x45, 0xb6,
	0x96, 0xb7, 0xad, 0x9a, 0xae, 0x90, 0x5c, 0x45, 0x83, 0x66, 0x5a, 0xd5, 0xe1, 0x31, 0xb8, 0x94,
	0x55, 0x6b, 0x19, 0x39, 0x5a, 0x6c, 0x1a, 0x35, 0x49, 0x6a, 0x8c, 0x23, 0xb7, 0x99, 0x54, 0xb5,
	0xd2, 0x36, 0x4a, 0xf3, 0x1a, 0x4d, 0xd5, 0xe1, 0x21, 0xd8, 0x34, 0x8f, 0xfc, 0xd8, 0x4a, 0x0f,
	0x66, 0x36, 0xcd, 0x93, 0x33, 0xf0, 0xc6, 0xe3, 0xd7, 0xb4, 0x96, 0x61, 0x0a, 0x4e, 0x8e, 0x24,
	0x8a, 0xac, 0x78, 0x90, 0x8e, 0xce, 0x8e, 0x33, 0x9d, 0x20, 0xdb, 0x73, 0x3f, 0xd3, 0x13, 0xc9,
	0x3b, 0xf0, 0xe6, 0x55, 0xc9, 0x51, 0x1e, 0x3e, 0x01, 0x28, 0x49, 0x5e, 0x10, 0x71, 0xb1, 0xcb,
	0x64, 0x28, 0x5d, 0x5a, 0xfb, 0xe6, 0xb4, 0xee, 0x2f, 0x69, 0xb5, 0x03, 0x4f, 0xbb, 0x6b, 0x76,
	0x7d, 0xb3, 0xc0, 0x9f, 0xb2, 0x5a, 0xa2, 0xb2, 0x0c, 0x4f, 0x60, 0x88, 0x39, 0x5b, 0xd2, 0x62,
	0x3a, 0x69, 0x77, 0x75, 0xbd, 0xba, 0x2a, 0x17, 0xc5, 0xc5, 0x6e, 0xd9, 0xb6, 0x6d, 0xbf, 0x99,
	0xd7, 0x44, 0xb4, 0x18, 0xb7, 0x6d, 0xe7, 0xce, 0x31, 0xdc, 0x3d, 0x02, 0xaf, 0xe6, 0x6b, 0x81,
	0x49, 0x6b, 0xae, 0xed, 0xba, 0x93, 0xf0, 0x8c, 0x93, 0x30, 0x92, 0xf8, 0xfb, 0x49, 0x4e, 0x60,
	0x58, 0x91, 0x96, 0xca, 0xb0, 0x71, 0xba, 0xed, 0x93, 0x0f, 0x10, 0xb4, 0x81, 0x48, 0x7e, 0xe7,
	0x91, 0xcc, 0xc5, 0x4e, 0x6f, 0xf1, 0x67, 0x1b, 0x46, 0x7a, 0x33, 0x93, 0x14, 0x49, 0x72, 0xe7,
	0xbb, 0x1f, 0x43, 0x80, 0x57, 0x88, 0x31, 0x52, 0x4e, 0x27, 0xed, 0xf2, 0x9d, 0xd0, 0xc1, 0x76,
	0x0d, 0xd8, 0x7f, 0x06, 0x35, 0x86, 0x11, 0x17, 0xc5, 0x25, 0x2f, 0x29, 0xa6, 0xa4, 0x8e, 0x86,
	0xf1, 0x20, 0x0d, 0x66, 0xa6, 0xa4, 0xae, 0x87, 0x44, 0x51, 0x47, 0x41, 0x3c, 0x50, 0x37, 0x90,
	0xaa, 0xf7, 0x88, 0x40, 0x8f, 0xc8, 0x17, 0x0b, 0x0e, 0x0c, 0x22, 0xf9, 0x5f, 0x46, 0x62, 0x5a,
	0x73, 0x7b, 0xd6, 0x3e, 0xd9, 0xe0, 0xcf, 0xab, 0x42, 0xa0, 0xfc, 0xfe, 0xa0, 0x48, 0xf2, 0xc3,
	0x02, 0x6f, 0xca, 0x36, 0xfc, 0xea, 0x76, 0x18, 0x7b, 0xc1, 0xec, 0x7e, 0xb0, 0x18, 0x46, 0xb8,
	0xfb, 0x37, 0x9b, 0xb4, 0x50, 0x4c, 0xc9, 0x84, 0xe9, 0xfc, 0x16, 0xa6, 0xbb, 0x0f, 0xf3, 0x08,
	0x06, 0x4b, 0xcc, 0x5a, 0x32, 0xaa, 0xec, 0xc2, 0xf9, 0x46, 0xb8, 0x14, 0x1e, 0x4a, 0x24, 0x0a,
	0x22, 0x5f, 0xb2, 0xbc, 0xe2, 0x94, 0xc9, 0x2d, 0x96, 0xbe, 0x9c, 0x7c, 0xd4, 0xbf, 0x52, 0x15,
	0xf5, 0x45, 0x7d, 0xcd, 0xf0, 0x3f, 0x90, 0x57, 0x39, 0x44, 0x65, 0xb9, 0x40, 0xf8, 0xaa, 0xbd,
	0x13, 0xba, 0xbe, 0x63, 0x31, 0xbc, 0x9d, 0x45, 0x70, 0x33, 0x8b, 0xef, 0x16, 0xb8, 0x6f, 0xd6,
	0x44, 0x5c, 0xff, 0xef, 0xa7, 0x7e, 0xfe, 0x0a, 0x9e, 0x52, 0x9e, 0x11, 0xc6, 0x38, 0xcf, 0x96,
	0x68, 0x21, 0x28, 0x7e, 0x86, 0x4b, 0x4a, 0x98, 0xcc, 0xd4, 0xab, 0x46, 0xf3, 0x52, 0xd1, 0x3c,
	0x91, 0xcf, 0x0f, 0xbb, 0x47, 0xf2, 0xa5, 0x52, 0xdf, 0x1e, 0xf5, 0x5f, 0x46, 0x16, 0x9e, 0x6e,
	0x9e, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x86, 0xc4, 0x2f, 0xa7, 0x08, 0x00, 0x00,
}
