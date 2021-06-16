// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

package pc

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Storage_Driver int32

const (
	Storage_UNKNOWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

var Storage_Driver_name = map[int32]string{
	0: "UNKNOWN",
	1: "HDD",
	2: "SSD",
}

var Storage_Driver_value = map[string]int32{
	"UNKNOWN": 0,
	"HDD":     1,
	"SSD":     2,
}

func (x Storage_Driver) String() string {
	return proto.EnumName(Storage_Driver_name, int32(x))
}

func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0, 0}
}

type Storage struct {
	Driver               Storage_Driver `protobuf:"varint,1,opt,name=driver,proto3,enum=pc.Storage_Driver" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}
func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_Driver {
	if m != nil {
		return m.Driver
	}
	return Storage_UNKNOWN
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("pc.Storage_Driver", Storage_Driver_name, Storage_Driver_value)
	proto.RegisterType((*Storage)(nil), "pc.Storage")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x96, 0xe2, 0xcc,
	0x4d, 0xcd, 0x85, 0x70, 0x95, 0x9a, 0x18, 0xb9, 0xd8, 0x83, 0x21, 0x0a, 0x84, 0xb4, 0xb8, 0xd8,
	0x52, 0x8a, 0x32, 0xcb, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x84, 0xf4, 0x0a,
	0x92, 0xf5, 0xa0, 0x92, 0x7a, 0x2e, 0x60, 0x99, 0x20, 0xa8, 0x0a, 0x21, 0x25, 0x2e, 0xb6, 0xdc,
	0xd4, 0xdc, 0xfc, 0xa2, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x2e, 0x90, 0x5a, 0x5f,
	0xb0, 0x48, 0x10, 0x54, 0x46, 0x49, 0x9d, 0x8b, 0x0d, 0xa2, 0x4b, 0x88, 0x9b, 0x8b, 0x3d, 0xd4,
	0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80, 0x41, 0x88, 0x9d, 0x8b, 0xd9, 0xc3, 0xc5, 0x45, 0x80,
	0x11, 0xc4, 0x08, 0x0e, 0x76, 0x11, 0x60, 0x4a, 0x62, 0x03, 0xbb, 0xc5, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x3e, 0x14, 0xa4, 0xab, 0x00, 0x00, 0x00,
}