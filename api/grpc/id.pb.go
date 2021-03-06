// Code generated by protoc-gen-go. DO NOT EDIT.
// source: id.proto

package gid

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

type ReqId struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqId) Reset()         { *m = ReqId{} }
func (m *ReqId) String() string { return proto.CompactTextString(m) }
func (*ReqId) ProtoMessage()    {}
func (*ReqId) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{0}
}

func (m *ReqId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqId.Unmarshal(m, b)
}
func (m *ReqId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqId.Marshal(b, m, deterministic)
}
func (m *ReqId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqId.Merge(m, src)
}
func (m *ReqId) XXX_Size() int {
	return xxx_messageInfo_ReqId.Size(m)
}
func (m *ReqId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqId.DiscardUnknown(m)
}

var xxx_messageInfo_ReqId proto.InternalMessageInfo

func (m *ReqId) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ResId struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResId) Reset()         { *m = ResId{} }
func (m *ResId) String() string { return proto.CompactTextString(m) }
func (*ResId) ProtoMessage()    {}
func (*ResId) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{1}
}

func (m *ResId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResId.Unmarshal(m, b)
}
func (m *ResId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResId.Marshal(b, m, deterministic)
}
func (m *ResId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResId.Merge(m, src)
}
func (m *ResId) XXX_Size() int {
	return xxx_messageInfo_ResId.Size(m)
}
func (m *ResId) XXX_DiscardUnknown() {
	xxx_messageInfo_ResId.DiscardUnknown(m)
}

var xxx_messageInfo_ResId proto.InternalMessageInfo

func (m *ResId) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ResId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReqPing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqPing) Reset()         { *m = ReqPing{} }
func (m *ReqPing) String() string { return proto.CompactTextString(m) }
func (*ReqPing) ProtoMessage()    {}
func (*ReqPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{2}
}

func (m *ReqPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqPing.Unmarshal(m, b)
}
func (m *ReqPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqPing.Marshal(b, m, deterministic)
}
func (m *ReqPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqPing.Merge(m, src)
}
func (m *ReqPing) XXX_Size() int {
	return xxx_messageInfo_ReqPing.Size(m)
}
func (m *ReqPing) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqPing.DiscardUnknown(m)
}

var xxx_messageInfo_ReqPing proto.InternalMessageInfo

type ResPong struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResPong) Reset()         { *m = ResPong{} }
func (m *ResPong) String() string { return proto.CompactTextString(m) }
func (*ResPong) ProtoMessage()    {}
func (*ResPong) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{3}
}

func (m *ResPong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResPong.Unmarshal(m, b)
}
func (m *ResPong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResPong.Marshal(b, m, deterministic)
}
func (m *ResPong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResPong.Merge(m, src)
}
func (m *ResPong) XXX_Size() int {
	return xxx_messageInfo_ResPong.Size(m)
}
func (m *ResPong) XXX_DiscardUnknown() {
	xxx_messageInfo_ResPong.DiscardUnknown(m)
}

var xxx_messageInfo_ResPong proto.InternalMessageInfo

func (m *ResPong) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ResPong) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ReqTagCreate struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	MaxId                int64    `protobuf:"varint,2,opt,name=maxId,proto3" json:"maxId,omitempty"`
	Step                 int64    `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTagCreate) Reset()         { *m = ReqTagCreate{} }
func (m *ReqTagCreate) String() string { return proto.CompactTextString(m) }
func (*ReqTagCreate) ProtoMessage()    {}
func (*ReqTagCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{4}
}

func (m *ReqTagCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTagCreate.Unmarshal(m, b)
}
func (m *ReqTagCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTagCreate.Marshal(b, m, deterministic)
}
func (m *ReqTagCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTagCreate.Merge(m, src)
}
func (m *ReqTagCreate) XXX_Size() int {
	return xxx_messageInfo_ReqTagCreate.Size(m)
}
func (m *ReqTagCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTagCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTagCreate proto.InternalMessageInfo

func (m *ReqTagCreate) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ReqTagCreate) GetMaxId() int64 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

func (m *ReqTagCreate) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *ReqTagCreate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type ResTagCreate struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResTagCreate) Reset()         { *m = ResTagCreate{} }
func (m *ResTagCreate) String() string { return proto.CompactTextString(m) }
func (*ResTagCreate) ProtoMessage()    {}
func (*ResTagCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{5}
}

func (m *ResTagCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResTagCreate.Unmarshal(m, b)
}
func (m *ResTagCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResTagCreate.Marshal(b, m, deterministic)
}
func (m *ResTagCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResTagCreate.Merge(m, src)
}
func (m *ResTagCreate) XXX_Size() int {
	return xxx_messageInfo_ResTagCreate.Size(m)
}
func (m *ResTagCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ResTagCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ResTagCreate proto.InternalMessageInfo

func (m *ResTagCreate) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqRandId struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRandId) Reset()         { *m = ReqRandId{} }
func (m *ReqRandId) String() string { return proto.CompactTextString(m) }
func (*ReqRandId) ProtoMessage()    {}
func (*ReqRandId) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{6}
}

func (m *ReqRandId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRandId.Unmarshal(m, b)
}
func (m *ReqRandId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRandId.Marshal(b, m, deterministic)
}
func (m *ReqRandId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRandId.Merge(m, src)
}
func (m *ReqRandId) XXX_Size() int {
	return xxx_messageInfo_ReqRandId.Size(m)
}
func (m *ReqRandId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRandId.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRandId proto.InternalMessageInfo

func (m *ReqRandId) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ResRandId struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResRandId) Reset()         { *m = ResRandId{} }
func (m *ResRandId) String() string { return proto.CompactTextString(m) }
func (*ResRandId) ProtoMessage()    {}
func (*ResRandId) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3ad0c1fc883139, []int{7}
}

func (m *ResRandId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResRandId.Unmarshal(m, b)
}
func (m *ResRandId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResRandId.Marshal(b, m, deterministic)
}
func (m *ResRandId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResRandId.Merge(m, src)
}
func (m *ResRandId) XXX_Size() int {
	return xxx_messageInfo_ResRandId.Size(m)
}
func (m *ResRandId) XXX_DiscardUnknown() {
	xxx_messageInfo_ResRandId.DiscardUnknown(m)
}

var xxx_messageInfo_ResRandId proto.InternalMessageInfo

func (m *ResRandId) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ResRandId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ReqId)(nil), "gid.ReqId")
	proto.RegisterType((*ResId)(nil), "gid.ResId")
	proto.RegisterType((*ReqPing)(nil), "gid.ReqPing")
	proto.RegisterType((*ResPong)(nil), "gid.ResPong")
	proto.RegisterType((*ReqTagCreate)(nil), "gid.ReqTagCreate")
	proto.RegisterType((*ResTagCreate)(nil), "gid.ResTagCreate")
	proto.RegisterType((*ReqRandId)(nil), "gid.ReqRandId")
	proto.RegisterType((*ResRandId)(nil), "gid.ResRandId")
}

func init() { proto.RegisterFile("id.proto", fileDescriptor_4b3ad0c1fc883139) }

var fileDescriptor_4b3ad0c1fc883139 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0xd2, 0x46, 0x33, 0x15, 0x91, 0x45, 0x24, 0x0a, 0x42, 0x59, 0x2f, 0x3d, 0xe5,
	0x60, 0xaf, 0x1e, 0x44, 0x4f, 0xb9, 0x95, 0xd1, 0x3f, 0x30, 0x61, 0x86, 0x65, 0x91, 0x26, 0xcd,
	0xee, 0x0a, 0xfe, 0x7c, 0xd9, 0x6d, 0xc4, 0x83, 0x1e, 0x42, 0x6f, 0x6f, 0xde, 0xc7, 0xbe, 0x79,
	0xcc, 0xc2, 0xb9, 0xe5, 0xe6, 0xe0, 0x86, 0x30, 0xa8, 0xc2, 0x58, 0xbe, 0x83, 0x8e, 0xbc, 0x1c,
	0x0d, 0x7d, 0x0b, 0x4b, 0x94, 0xb1, 0x65, 0x75, 0x05, 0x45, 0x20, 0x53, 0x67, 0xeb, 0x6c, 0x53,
	0x61, 0x94, 0xfa, 0x29, 0x22, 0xdf, 0xb2, 0x7a, 0x80, 0xd2, 0x07, 0x0a, 0x9f, 0x3e, 0xd1, 0xd5,
	0xe3, 0xaa, 0x31, 0x96, 0x9b, 0xb7, 0x64, 0xe1, 0x84, 0xd4, 0x25, 0xe4, 0x96, 0xeb, 0x7c, 0x9d,
	0x6d, 0x0a, 0xcc, 0x2d, 0xeb, 0x0a, 0xce, 0x50, 0xc6, 0x9d, 0xed, 0x8d, 0x7e, 0x89, 0xd2, 0xef,
	0x86, 0xde, 0xcc, 0x8b, 0x52, 0xb0, 0x60, 0x0a, 0x94, 0xc2, 0x2a, 0x4c, 0x5a, 0x77, 0x70, 0x81,
	0x32, 0xbe, 0x93, 0x79, 0x75, 0x42, 0x41, 0xfe, 0xd6, 0x55, 0xd7, 0xb0, 0xdc, 0xd3, 0x57, 0xfb,
	0xd3, 0xe1, 0x38, 0xc4, 0x2c, 0x1f, 0xe4, 0x50, 0x17, 0xc9, 0x4c, 0x5a, 0xdd, 0x40, 0xe9, 0x64,
	0x4f, 0xee, 0xa3, 0x5e, 0xa4, 0xe7, 0xd3, 0xa4, 0xb7, 0x71, 0x87, 0xff, 0xdd, 0x31, 0xa7, 0xac,
	0xbe, 0x87, 0x0a, 0x65, 0x44, 0xea, 0xf9, 0xdf, 0x23, 0x3e, 0x47, 0xec, 0x27, 0x7c, 0xca, 0x21,
	0xbb, 0x32, 0x7d, 0xd4, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x44, 0x48, 0xdc, 0x42, 0xc5, 0x01,
	0x00, 0x00,
}
