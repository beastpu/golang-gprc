// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: laptop.proto

package pc

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storage  []*Storage `protobuf:"bytes,7,rep,name=storage,proto3" json:"storage,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight        `protobuf_oneof:"weight"`
	PriceUsd             float64                `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32                 `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_28a7e4886f546705, []int{0}
}
func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof" json:"weight_kg,omitempty"`
}
type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof" json:"weight_lb,omitempty"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}
func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorage() []*Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func init() {
	proto.RegisterType((*Laptop)(nil), "pc.Laptop")
}

func init() { proto.RegisterFile("laptop.proto", fileDescriptor_28a7e4886f546705) }

var fileDescriptor_28a7e4886f546705 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x41, 0x8f, 0x94, 0x30,
	0x14, 0xc7, 0x05, 0x66, 0x19, 0x78, 0x30, 0x63, 0xd2, 0x78, 0xa8, 0xb3, 0x1a, 0x71, 0x13, 0x13,
	0x4e, 0x6c, 0xa2, 0x27, 0x8f, 0xea, 0x41, 0x93, 0x5d, 0x93, 0x4d, 0xd7, 0x39, 0x78, 0x22, 0x85,
	0x3e, 0x91, 0x2c, 0x4c, 0x9b, 0xb6, 0xc4, 0xcc, 0xb7, 0xf4, 0x23, 0x19, 0x4a, 0x51, 0xf7, 0xc6,
	0xff, 0xf7, 0xfb, 0x93, 0xd7, 0xf6, 0x41, 0x3e, 0x70, 0x65, 0xa5, 0xaa, 0x94, 0x96, 0x56, 0x92,
	0x50, 0xb5, 0x87, 0x74, 0xc4, 0x71, 0x89, 0x87, 0xa7, 0x4a, 0xcb, 0x16, 0x8d, 0x91, 0xda, 0x83,
	0x9d, 0xb1, 0x52, 0xf3, 0x0e, 0x7d, 0xcc, 0x4d, 0xab, 0x11, 0x4f, 0x3e, 0xed, 0x1f, 0xf0, 0xdc,
	0x48, 0xae, 0x85, 0xcf, 0xaf, 0x3a, 0x29, 0xbb, 0x01, 0xaf, 0x5d, 0x6a, 0xa6, 0x1f, 0xd7, 0xb6,
	0x1f, 0xd1, 0x58, 0x3e, 0xfa, 0x69, 0x57, 0xbf, 0x23, 0x88, 0x6f, 0xdd, 0x78, 0xb2, 0x87, 0xb0,
	0x17, 0x34, 0x28, 0x82, 0x32, 0x65, 0x61, 0x2f, 0xc8, 0x33, 0xb8, 0x68, 0x34, 0x3f, 0x09, 0x1a,
	0x3a, 0xb4, 0x04, 0x42, 0x60, 0x73, 0xe2, 0x23, 0xd2, 0xc8, 0x41, 0xf7, 0x4d, 0x9e, 0x43, 0xd4,
	0xaa, 0x89, 0x6e, 0x8a, 0xa0, 0xcc, 0xde, 0x6e, 0x2b, 0xd5, 0x56, 0x9f, 0xee, 0x8e, 0x6c, 0x66,
	0xe4, 0x05, 0x44, 0x9a, 0x8f, 0xf4, 0xc2, 0x29, 0x98, 0xd5, 0x57, 0x1c, 0xa5, 0x3e, 0xb3, 0x19,
	0x93, 0x4b, 0xd8, 0x74, 0x6a, 0x32, 0x34, 0x2e, 0xa2, 0xf5, 0xcf, 0xcf, 0x77, 0x47, 0xe6, 0x20,
	0x79, 0x03, 0x5b, 0x7f, 0x55, 0xba, 0x75, 0x3e, 0x9b, 0xfd, 0xfd, 0x82, 0xd8, 0xea, 0xc8, 0x15,
	0xc4, 0xcb, 0x13, 0xd0, 0xe4, 0xdf, 0x90, 0x7b, 0x47, 0x98, 0x37, 0xa4, 0x84, 0x64, 0x7d, 0x18,
	0x9a, 0xba, 0x56, 0x3e, 0xb7, 0x6e, 0x3c, 0x63, 0x7f, 0x2d, 0x79, 0x09, 0xe9, 0x2f, 0xec, 0xbb,
	0x9f, 0xb6, 0x7e, 0xe8, 0x28, 0x14, 0x41, 0x19, 0x7c, 0x79, 0xc2, 0x92, 0x05, 0xdd, 0x74, 0xff,
	0xe9, 0xa1, 0xa1, 0xd9, 0x63, 0x7d, 0xdb, 0x90, 0x4b, 0x48, 0x95, 0xee, 0x5b, 0xac, 0x27, 0x23,
	0x68, 0x3e, 0x6b, 0x96, 0x38, 0x70, 0x34, 0x82, 0xbc, 0x86, 0x5c, 0xe3, 0x80, 0xdc, 0x60, 0x7d,
	0x46, 0xae, 0xe9, 0xae, 0x08, 0xca, 0x1d, 0xcb, 0x3c, 0xfb, 0x8e, 0x5c, 0x93, 0xf7, 0x00, 0x93,
	0x12, 0xdc, 0xa2, 0xa8, 0xb9, 0xa5, 0x7b, 0x77, 0xd2, 0x43, 0xb5, 0xec, 0xb0, 0x5a, 0x77, 0x58,
	0x7d, 0x5b, 0x77, 0xc8, 0x52, 0xdf, 0xfe, 0x60, 0x3f, 0x26, 0x10, 0x2f, 0xc7, 0x68, 0x62, 0x57,
	0x7c, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x66, 0xc9, 0x7a, 0x19, 0x57, 0x02, 0x00, 0x00,
}
