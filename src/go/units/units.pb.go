// Code generated by protoc-gen-go. DO NOT EDIT.
// source: units/units.proto

package units // import "github.com/airmap/interfaces/src/go/units"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Degrees models an angular quantity in degrees [°].
type Degrees struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Degrees) Reset()         { *m = Degrees{} }
func (m *Degrees) String() string { return proto.CompactTextString(m) }
func (*Degrees) ProtoMessage()    {}
func (*Degrees) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_a1471f0fe313ad8a, []int{0}
}
func (m *Degrees) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Degrees.Unmarshal(m, b)
}
func (m *Degrees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Degrees.Marshal(b, m, deterministic)
}
func (dst *Degrees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Degrees.Merge(dst, src)
}
func (m *Degrees) XXX_Size() int {
	return xxx_messageInfo_Degrees.Size(m)
}
func (m *Degrees) XXX_DiscardUnknown() {
	xxx_messageInfo_Degrees.DiscardUnknown(m)
}

var xxx_messageInfo_Degrees proto.InternalMessageInfo

func (m *Degrees) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Celsius models a quantity in degrees Celsius [°C].
type Celsius struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Celsius) Reset()         { *m = Celsius{} }
func (m *Celsius) String() string { return proto.CompactTextString(m) }
func (*Celsius) ProtoMessage()    {}
func (*Celsius) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_a1471f0fe313ad8a, []int{1}
}
func (m *Celsius) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Celsius.Unmarshal(m, b)
}
func (m *Celsius) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Celsius.Marshal(b, m, deterministic)
}
func (dst *Celsius) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Celsius.Merge(dst, src)
}
func (m *Celsius) XXX_Size() int {
	return xxx_messageInfo_Celsius.Size(m)
}
func (m *Celsius) XXX_DiscardUnknown() {
	xxx_messageInfo_Celsius.DiscardUnknown(m)
}

var xxx_messageInfo_Celsius proto.InternalMessageInfo

func (m *Celsius) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Meters models a quantity in meters.
type Meters struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meters) Reset()         { *m = Meters{} }
func (m *Meters) String() string { return proto.CompactTextString(m) }
func (*Meters) ProtoMessage()    {}
func (*Meters) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_a1471f0fe313ad8a, []int{2}
}
func (m *Meters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meters.Unmarshal(m, b)
}
func (m *Meters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meters.Marshal(b, m, deterministic)
}
func (dst *Meters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meters.Merge(dst, src)
}
func (m *Meters) XXX_Size() int {
	return xxx_messageInfo_Meters.Size(m)
}
func (m *Meters) XXX_DiscardUnknown() {
	xxx_messageInfo_Meters.DiscardUnknown(m)
}

var xxx_messageInfo_Meters proto.InternalMessageInfo

func (m *Meters) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Foot models a quantity in feet.
type Foot struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foot) Reset()         { *m = Foot{} }
func (m *Foot) String() string { return proto.CompactTextString(m) }
func (*Foot) ProtoMessage()    {}
func (*Foot) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_a1471f0fe313ad8a, []int{3}
}
func (m *Foot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foot.Unmarshal(m, b)
}
func (m *Foot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foot.Marshal(b, m, deterministic)
}
func (dst *Foot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foot.Merge(dst, src)
}
func (m *Foot) XXX_Size() int {
	return xxx_messageInfo_Foot.Size(m)
}
func (m *Foot) XXX_DiscardUnknown() {
	xxx_messageInfo_Foot.DiscardUnknown(m)
}

var xxx_messageInfo_Foot proto.InternalMessageInfo

func (m *Foot) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// MetersPerSecond models a quantity in meters per second [m/s].
type MetersPerSecond struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetersPerSecond) Reset()         { *m = MetersPerSecond{} }
func (m *MetersPerSecond) String() string { return proto.CompactTextString(m) }
func (*MetersPerSecond) ProtoMessage()    {}
func (*MetersPerSecond) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_a1471f0fe313ad8a, []int{4}
}
func (m *MetersPerSecond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetersPerSecond.Unmarshal(m, b)
}
func (m *MetersPerSecond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetersPerSecond.Marshal(b, m, deterministic)
}
func (dst *MetersPerSecond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetersPerSecond.Merge(dst, src)
}
func (m *MetersPerSecond) XXX_Size() int {
	return xxx_messageInfo_MetersPerSecond.Size(m)
}
func (m *MetersPerSecond) XXX_DiscardUnknown() {
	xxx_messageInfo_MetersPerSecond.DiscardUnknown(m)
}

var xxx_messageInfo_MetersPerSecond proto.InternalMessageInfo

func (m *MetersPerSecond) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Pascal models a quantity in Pascals [Pa].
type Pascal struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pascal) Reset()         { *m = Pascal{} }
func (m *Pascal) String() string { return proto.CompactTextString(m) }
func (*Pascal) ProtoMessage()    {}
func (*Pascal) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_a1471f0fe313ad8a, []int{5}
}
func (m *Pascal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pascal.Unmarshal(m, b)
}
func (m *Pascal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pascal.Marshal(b, m, deterministic)
}
func (dst *Pascal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pascal.Merge(dst, src)
}
func (m *Pascal) XXX_Size() int {
	return xxx_messageInfo_Pascal.Size(m)
}
func (m *Pascal) XXX_DiscardUnknown() {
	xxx_messageInfo_Pascal.DiscardUnknown(m)
}

var xxx_messageInfo_Pascal proto.InternalMessageInfo

func (m *Pascal) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Percent models a quantity as a percentage [%].
type Percent struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Percent) Reset()         { *m = Percent{} }
func (m *Percent) String() string { return proto.CompactTextString(m) }
func (*Percent) ProtoMessage()    {}
func (*Percent) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_a1471f0fe313ad8a, []int{6}
}
func (m *Percent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Percent.Unmarshal(m, b)
}
func (m *Percent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Percent.Marshal(b, m, deterministic)
}
func (dst *Percent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Percent.Merge(dst, src)
}
func (m *Percent) XXX_Size() int {
	return xxx_messageInfo_Percent.Size(m)
}
func (m *Percent) XXX_DiscardUnknown() {
	xxx_messageInfo_Percent.DiscardUnknown(m)
}

var xxx_messageInfo_Percent proto.InternalMessageInfo

func (m *Percent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Degrees)(nil), "units.Degrees")
	proto.RegisterType((*Celsius)(nil), "units.Celsius")
	proto.RegisterType((*Meters)(nil), "units.Meters")
	proto.RegisterType((*Foot)(nil), "units.Foot")
	proto.RegisterType((*MetersPerSecond)(nil), "units.MetersPerSecond")
	proto.RegisterType((*Pascal)(nil), "units.Pascal")
	proto.RegisterType((*Percent)(nil), "units.Percent")
}

func init() { proto.RegisterFile("units/units.proto", fileDescriptor_units_a1471f0fe313ad8a) }

var fileDescriptor_units_a1471f0fe313ad8a = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x0b, 0xc2, 0x30,
	0x14, 0x84, 0x29, 0xd8, 0x0a, 0x59, 0xc4, 0xe2, 0xe0, 0x20, 0x2a, 0x2e, 0x2a, 0x42, 0x33, 0xf8,
	0x0f, 0x54, 0xdc, 0x84, 0xa2, 0x9b, 0x5b, 0x1a, 0xcf, 0x1a, 0x68, 0xf3, 0xca, 0x4b, 0xe2, 0xef,
	0x17, 0xdb, 0x39, 0xcb, 0x83, 0xe3, 0xee, 0x71, 0xc7, 0x27, 0xa6, 0xc1, 0x1a, 0xef, 0x64, 0x7f,
	0x8b, 0x8e, 0xc9, 0x53, 0x9e, 0xf6, 0x62, 0xb3, 0x12, 0xe3, 0x0b, 0x6a, 0x06, 0x5c, 0x3e, 0x13,
	0xe9, 0x57, 0x35, 0x01, 0xf3, 0x64, 0x9d, 0xec, 0x92, 0xfb, 0x20, 0xfe, 0x81, 0x33, 0x1a, 0x67,
	0x42, 0x2c, 0xb0, 0x14, 0xd9, 0x0d, 0x1e, 0x1c, 0xf3, 0x17, 0x62, 0x74, 0x25, 0xf2, 0x11, 0x77,
	0x2b, 0x26, 0xc3, 0x77, 0x09, 0x7e, 0x40, 0x93, 0x7d, 0xc5, 0x6b, 0x4a, 0xe5, 0xb4, 0x6a, 0xe2,
	0x3b, 0x4b, 0xb0, 0x86, 0x8d, 0x34, 0x9d, 0x0e, 0xcf, 0x7d, 0x6d, 0xfc, 0x27, 0x54, 0x85, 0xa6,
	0x56, 0x2a, 0xc3, 0xad, 0xea, 0xa4, 0xb1, 0x1e, 0xfc, 0x56, 0x1a, 0x4e, 0x3a, 0xd6, 0xb2, 0xa6,
	0x81, 0x51, 0x95, 0xf5, 0x90, 0x8e, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xfe, 0x99, 0x5c,
	0x39, 0x01, 0x00, 0x00,
}
