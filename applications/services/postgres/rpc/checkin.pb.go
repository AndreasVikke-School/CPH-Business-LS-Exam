// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkin.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type CheckIns struct {
	CheckIn              []*CheckIn `protobuf:"bytes,1,rep,name=checkIn,proto3" json:"checkIn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CheckIns) Reset()         { *m = CheckIns{} }
func (m *CheckIns) String() string { return proto.CompactTextString(m) }
func (*CheckIns) ProtoMessage()    {}
func (*CheckIns) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{0}
}

func (m *CheckIns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckIns.Unmarshal(m, b)
}
func (m *CheckIns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckIns.Marshal(b, m, deterministic)
}
func (m *CheckIns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckIns.Merge(m, src)
}
func (m *CheckIns) XXX_Size() int {
	return xxx_messageInfo_CheckIns.Size(m)
}
func (m *CheckIns) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckIns.DiscardUnknown(m)
}

var xxx_messageInfo_CheckIns proto.InternalMessageInfo

func (m *CheckIns) GetCheckIn() []*CheckIn {
	if m != nil {
		return m.CheckIn
	}
	return nil
}

type CheckIn struct {
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AttendanceCode       int64    `protobuf:"varint,1,opt,name=attendanceCode,proto3" json:"attendanceCode,omitempty"`
	StudentId            int64    `protobuf:"varint,3,opt,name=studentId,proto3" json:"studentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckIn) Reset()         { *m = CheckIn{} }
func (m *CheckIn) String() string { return proto.CompactTextString(m) }
func (*CheckIn) ProtoMessage()    {}
func (*CheckIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_072e71e6019dc001, []int{1}
}

func (m *CheckIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckIn.Unmarshal(m, b)
}
func (m *CheckIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckIn.Marshal(b, m, deterministic)
}
func (m *CheckIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckIn.Merge(m, src)
}
func (m *CheckIn) XXX_Size() int {
	return xxx_messageInfo_CheckIn.Size(m)
}
func (m *CheckIn) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckIn.DiscardUnknown(m)
}

var xxx_messageInfo_CheckIn proto.InternalMessageInfo

func (m *CheckIn) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CheckIn) GetAttendanceCode() int64 {
	if m != nil {
		return m.AttendanceCode
	}
	return 0
}

func (m *CheckIn) GetStudentId() int64 {
	if m != nil {
		return m.StudentId
	}
	return 0
}

func init() {
	proto.RegisterType((*CheckIns)(nil), "rpc.CheckIns")
	proto.RegisterType((*CheckIn)(nil), "rpc.CheckIn")
}

func init() {
	proto.RegisterFile("checkin.proto", fileDescriptor_072e71e6019dc001)
}

var fileDescriptor_072e71e6019dc001 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0x4d,
	0xce, 0xce, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0x92,
	0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x17, 0x25,
	0x16, 0x14, 0xa4, 0x16, 0x15, 0x43, 0x14, 0x49, 0x49, 0xa3, 0xcb, 0xa7, 0xe6, 0x16, 0x94, 0x54,
	0x42, 0x24, 0x95, 0x8c, 0xb8, 0x38, 0x9c, 0x41, 0x46, 0x7a, 0xe6, 0x15, 0x0b, 0xa9, 0x71, 0xb1,
	0x27, 0x43, 0xd8, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x3c, 0x7a, 0x45, 0x05, 0xc9, 0x7a,
	0x50, 0xf9, 0x20, 0x98, 0xa4, 0x52, 0x3c, 0x17, 0x3b, 0x54, 0x4c, 0x88, 0x8f, 0x8b, 0x29, 0x33,
	0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x88, 0x29, 0x33, 0x45, 0x48, 0x8d, 0x8b, 0x2f, 0xb1,
	0xa4, 0x24, 0x35, 0x2f, 0x25, 0x31, 0x2f, 0x39, 0xd5, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x11, 0x2c,
	0x87, 0x26, 0x2a, 0x24, 0xc3, 0xc5, 0x59, 0x5c, 0x52, 0x9a, 0x92, 0x9a, 0x57, 0xe2, 0x99, 0x22,
	0xc1, 0x0c, 0x56, 0x82, 0x10, 0x30, 0x6a, 0x61, 0xe4, 0xe2, 0x81, 0xda, 0x10, 0x00, 0xf6, 0xa7,
	0x35, 0x17, 0x9f, 0x7b, 0x6a, 0x09, 0x54, 0xc8, 0xa9, 0xd2, 0x33, 0x45, 0x48, 0x5a, 0x0f, 0xe2,
	0x2b, 0x3d, 0x98, 0xaf, 0xf4, 0x3c, 0xf3, 0x4a, 0xcc, 0x4c, 0xc2, 0x12, 0x73, 0x4a, 0x53, 0xa5,
	0x50, 0xdc, 0x2d, 0x64, 0x0e, 0xd6, 0xec, 0x98, 0x93, 0x03, 0xf7, 0xa8, 0x18, 0x86, 0x66, 0x57,
	0x50, 0x90, 0x48, 0xf1, 0x22, 0xeb, 0x2b, 0x76, 0xe2, 0x5c, 0xc5, 0xc4, 0x06, 0xb6, 0xbf, 0x38,
	0x89, 0x0d, 0xac, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x19, 0x77, 0x6c, 0x3d, 0x80, 0x01,
	0x00, 0x00,
}
