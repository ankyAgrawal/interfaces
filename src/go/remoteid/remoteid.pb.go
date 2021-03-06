// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remoteid/remoteid.proto

package remoteid

import (
	context "context"
	fmt "fmt"
	geo "github.com/airmap/interfaces/src/go/geo"
	ids "github.com/airmap/interfaces/src/go/ids"
	measurements "github.com/airmap/interfaces/src/go/measurements"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Aircraft_Type int32

const (
	Aircraft_UNKNOWN                Aircraft_Type = 0
	Aircraft_AIRPLANE               Aircraft_Type = 1
	Aircraft_ROTORCRAFT             Aircraft_Type = 2
	Aircraft_GYROPLANE              Aircraft_Type = 3
	Aircraft_VTOL                   Aircraft_Type = 4
	Aircraft_ORNITHOPTER            Aircraft_Type = 5
	Aircraft_GLIDER                 Aircraft_Type = 6
	Aircraft_KITE                   Aircraft_Type = 7
	Aircraft_FREE_BALOON            Aircraft_Type = 8
	Aircraft_CAPTIVE_BALLOON        Aircraft_Type = 9
	Aircraft_AIRSHIP                Aircraft_Type = 10
	Aircraft_FREE_FALL_OR_PARACHUTE Aircraft_Type = 11
	Aircraft_ROCKET                 Aircraft_Type = 12
	Aircraft_GROUND_OBSTACLE        Aircraft_Type = 13
	Aircraft_OTHER                  Aircraft_Type = 14
)

var Aircraft_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "AIRPLANE",
	2:  "ROTORCRAFT",
	3:  "GYROPLANE",
	4:  "VTOL",
	5:  "ORNITHOPTER",
	6:  "GLIDER",
	7:  "KITE",
	8:  "FREE_BALOON",
	9:  "CAPTIVE_BALLOON",
	10: "AIRSHIP",
	11: "FREE_FALL_OR_PARACHUTE",
	12: "ROCKET",
	13: "GROUND_OBSTACLE",
	14: "OTHER",
}

var Aircraft_Type_value = map[string]int32{
	"UNKNOWN":                0,
	"AIRPLANE":               1,
	"ROTORCRAFT":             2,
	"GYROPLANE":              3,
	"VTOL":                   4,
	"ORNITHOPTER":            5,
	"GLIDER":                 6,
	"KITE":                   7,
	"FREE_BALOON":            8,
	"CAPTIVE_BALLOON":        9,
	"AIRSHIP":                10,
	"FREE_FALL_OR_PARACHUTE": 11,
	"ROCKET":                 12,
	"GROUND_OBSTACLE":        13,
	"OTHER":                  14,
}

func (x Aircraft_Type) String() string {
	return proto.EnumName(Aircraft_Type_name, int32(x))
}

func (Aircraft_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{4, 0}
}

// MonitorAreaParameters bundles parameters used in the MonitorOperationsArea rpc.
type MonitorAreaParameters struct {
	AreaOfInterest       *geo.BoundingBox `protobuf:"bytes,1,opt,name=area_of_interest,json=areaOfInterest,proto3" json:"area_of_interest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MonitorAreaParameters) Reset()         { *m = MonitorAreaParameters{} }
func (m *MonitorAreaParameters) String() string { return proto.CompactTextString(m) }
func (*MonitorAreaParameters) ProtoMessage()    {}
func (*MonitorAreaParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{0}
}

func (m *MonitorAreaParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorAreaParameters.Unmarshal(m, b)
}
func (m *MonitorAreaParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorAreaParameters.Marshal(b, m, deterministic)
}
func (m *MonitorAreaParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorAreaParameters.Merge(m, src)
}
func (m *MonitorAreaParameters) XXX_Size() int {
	return xxx_messageInfo_MonitorAreaParameters.Size(m)
}
func (m *MonitorAreaParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorAreaParameters.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorAreaParameters proto.InternalMessageInfo

func (m *MonitorAreaParameters) GetAreaOfInterest() *geo.BoundingBox {
	if m != nil {
		return m.AreaOfInterest
	}
	return nil
}

type MonitorAreaResponse struct {
	// details is a discriminated union of response types
	//
	// Types that are valid to be assigned to Details:
	//	*MonitorAreaResponse_Added_
	//	*MonitorAreaResponse_Updated_
	//	*MonitorAreaResponse_Removed_
	Details              isMonitorAreaResponse_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MonitorAreaResponse) Reset()         { *m = MonitorAreaResponse{} }
func (m *MonitorAreaResponse) String() string { return proto.CompactTextString(m) }
func (*MonitorAreaResponse) ProtoMessage()    {}
func (*MonitorAreaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{1}
}

func (m *MonitorAreaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorAreaResponse.Unmarshal(m, b)
}
func (m *MonitorAreaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorAreaResponse.Marshal(b, m, deterministic)
}
func (m *MonitorAreaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorAreaResponse.Merge(m, src)
}
func (m *MonitorAreaResponse) XXX_Size() int {
	return xxx_messageInfo_MonitorAreaResponse.Size(m)
}
func (m *MonitorAreaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorAreaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorAreaResponse proto.InternalMessageInfo

type isMonitorAreaResponse_Details interface {
	isMonitorAreaResponse_Details()
}

type MonitorAreaResponse_Added_ struct {
	Added *MonitorAreaResponse_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}

type MonitorAreaResponse_Updated_ struct {
	Updated *MonitorAreaResponse_Updated `protobuf:"bytes,2,opt,name=updated,proto3,oneof"`
}

type MonitorAreaResponse_Removed_ struct {
	Removed *MonitorAreaResponse_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof"`
}

func (*MonitorAreaResponse_Added_) isMonitorAreaResponse_Details() {}

func (*MonitorAreaResponse_Updated_) isMonitorAreaResponse_Details() {}

func (*MonitorAreaResponse_Removed_) isMonitorAreaResponse_Details() {}

func (m *MonitorAreaResponse) GetDetails() isMonitorAreaResponse_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *MonitorAreaResponse) GetAdded() *MonitorAreaResponse_Added {
	if x, ok := m.GetDetails().(*MonitorAreaResponse_Added_); ok {
		return x.Added
	}
	return nil
}

func (m *MonitorAreaResponse) GetUpdated() *MonitorAreaResponse_Updated {
	if x, ok := m.GetDetails().(*MonitorAreaResponse_Updated_); ok {
		return x.Updated
	}
	return nil
}

func (m *MonitorAreaResponse) GetRemoved() *MonitorAreaResponse_Removed {
	if x, ok := m.GetDetails().(*MonitorAreaResponse_Removed_); ok {
		return x.Removed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MonitorAreaResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MonitorAreaResponse_Added_)(nil),
		(*MonitorAreaResponse_Updated_)(nil),
		(*MonitorAreaResponse_Removed_)(nil),
	}
}

// Added models an operation that has entered the monitored area.
type MonitorAreaResponse_Added struct {
	Operation            *Operation        `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Positions            []*geo.Position4D `protobuf:"bytes,2,rep,name=positions,proto3" json:"positions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MonitorAreaResponse_Added) Reset()         { *m = MonitorAreaResponse_Added{} }
func (m *MonitorAreaResponse_Added) String() string { return proto.CompactTextString(m) }
func (*MonitorAreaResponse_Added) ProtoMessage()    {}
func (*MonitorAreaResponse_Added) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{1, 0}
}

func (m *MonitorAreaResponse_Added) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorAreaResponse_Added.Unmarshal(m, b)
}
func (m *MonitorAreaResponse_Added) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorAreaResponse_Added.Marshal(b, m, deterministic)
}
func (m *MonitorAreaResponse_Added) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorAreaResponse_Added.Merge(m, src)
}
func (m *MonitorAreaResponse_Added) XXX_Size() int {
	return xxx_messageInfo_MonitorAreaResponse_Added.Size(m)
}
func (m *MonitorAreaResponse_Added) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorAreaResponse_Added.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorAreaResponse_Added proto.InternalMessageInfo

func (m *MonitorAreaResponse_Added) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MonitorAreaResponse_Added) GetPositions() []*geo.Position4D {
	if m != nil {
		return m.Positions
	}
	return nil
}

// Updated models an operation that has been updated in the monitored area.
type MonitorAreaResponse_Updated struct {
	// details is a discriminated union of update types
	//
	// Types that are valid to be assigned to Details:
	//	*MonitorAreaResponse_Updated_Operation
	//	*MonitorAreaResponse_Updated_State
	Details              isMonitorAreaResponse_Updated_Details `protobuf_oneof:"details"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MonitorAreaResponse_Updated) Reset()         { *m = MonitorAreaResponse_Updated{} }
func (m *MonitorAreaResponse_Updated) String() string { return proto.CompactTextString(m) }
func (*MonitorAreaResponse_Updated) ProtoMessage()    {}
func (*MonitorAreaResponse_Updated) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{1, 1}
}

func (m *MonitorAreaResponse_Updated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorAreaResponse_Updated.Unmarshal(m, b)
}
func (m *MonitorAreaResponse_Updated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorAreaResponse_Updated.Marshal(b, m, deterministic)
}
func (m *MonitorAreaResponse_Updated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorAreaResponse_Updated.Merge(m, src)
}
func (m *MonitorAreaResponse_Updated) XXX_Size() int {
	return xxx_messageInfo_MonitorAreaResponse_Updated.Size(m)
}
func (m *MonitorAreaResponse_Updated) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorAreaResponse_Updated.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorAreaResponse_Updated proto.InternalMessageInfo

type isMonitorAreaResponse_Updated_Details interface {
	isMonitorAreaResponse_Updated_Details()
}

type MonitorAreaResponse_Updated_Operation struct {
	Operation *Operation `protobuf:"bytes,1,opt,name=operation,proto3,oneof"`
}

type MonitorAreaResponse_Updated_State struct {
	State *Operation_State `protobuf:"bytes,2,opt,name=state,proto3,oneof"`
}

func (*MonitorAreaResponse_Updated_Operation) isMonitorAreaResponse_Updated_Details() {}

func (*MonitorAreaResponse_Updated_State) isMonitorAreaResponse_Updated_Details() {}

func (m *MonitorAreaResponse_Updated) GetDetails() isMonitorAreaResponse_Updated_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *MonitorAreaResponse_Updated) GetOperation() *Operation {
	if x, ok := m.GetDetails().(*MonitorAreaResponse_Updated_Operation); ok {
		return x.Operation
	}
	return nil
}

func (m *MonitorAreaResponse_Updated) GetState() *Operation_State {
	if x, ok := m.GetDetails().(*MonitorAreaResponse_Updated_State); ok {
		return x.State
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MonitorAreaResponse_Updated) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MonitorAreaResponse_Updated_Operation)(nil),
		(*MonitorAreaResponse_Updated_State)(nil),
	}
}

// Removed models an operation that has been removed from the monitored area.
type MonitorAreaResponse_Removed struct {
	Id                   *ids.Operation `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MonitorAreaResponse_Removed) Reset()         { *m = MonitorAreaResponse_Removed{} }
func (m *MonitorAreaResponse_Removed) String() string { return proto.CompactTextString(m) }
func (*MonitorAreaResponse_Removed) ProtoMessage()    {}
func (*MonitorAreaResponse_Removed) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{1, 2}
}

func (m *MonitorAreaResponse_Removed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorAreaResponse_Removed.Unmarshal(m, b)
}
func (m *MonitorAreaResponse_Removed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorAreaResponse_Removed.Marshal(b, m, deterministic)
}
func (m *MonitorAreaResponse_Removed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorAreaResponse_Removed.Merge(m, src)
}
func (m *MonitorAreaResponse_Removed) XXX_Size() int {
	return xxx_messageInfo_MonitorAreaResponse_Removed.Size(m)
}
func (m *MonitorAreaResponse_Removed) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorAreaResponse_Removed.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorAreaResponse_Removed proto.InternalMessageInfo

func (m *MonitorAreaResponse_Removed) GetId() *ids.Operation {
	if m != nil {
		return m.Id
	}
	return nil
}

// Operation models an aerial operation in a given area.
type Operation struct {
	Id                   *ids.Operation `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Operator             *Operator      `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Aircraft             *Aircraft      `protobuf:"bytes,3,opt,name=aircraft,proto3" json:"aircraft,omitempty"`
	Geometry             *geo.Polygon   `protobuf:"bytes,4,opt,name=geometry,proto3" json:"geometry,omitempty"`
	Purpose              string         `protobuf:"bytes,5,opt,name=purpose,proto3" json:"purpose,omitempty"`
	Simulated            bool           `protobuf:"varint,6,opt,name=simulated,proto3" json:"simulated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{2}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetId() *ids.Operation {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Operation) GetOperator() *Operator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *Operation) GetAircraft() *Aircraft {
	if m != nil {
		return m.Aircraft
	}
	return nil
}

func (m *Operation) GetGeometry() *geo.Polygon {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func (m *Operation) GetPurpose() string {
	if m != nil {
		return m.Purpose
	}
	return ""
}

func (m *Operation) GetSimulated() bool {
	if m != nil {
		return m.Simulated
	}
	return false
}

// State models an operation's current state
type Operation_State struct {
	Id                   *ids.Operation                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            *timestamp.Timestamp            `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Position             *measurements.Position_Absolute `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Velocity             *measurements.Velocity_Polar    `protobuf:"bytes,4,opt,name=velocity,proto3" json:"velocity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Operation_State) Reset()         { *m = Operation_State{} }
func (m *Operation_State) String() string { return proto.CompactTextString(m) }
func (*Operation_State) ProtoMessage()    {}
func (*Operation_State) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{2, 0}
}

func (m *Operation_State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation_State.Unmarshal(m, b)
}
func (m *Operation_State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation_State.Marshal(b, m, deterministic)
}
func (m *Operation_State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation_State.Merge(m, src)
}
func (m *Operation_State) XXX_Size() int {
	return xxx_messageInfo_Operation_State.Size(m)
}
func (m *Operation_State) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation_State.DiscardUnknown(m)
}

var xxx_messageInfo_Operation_State proto.InternalMessageInfo

func (m *Operation_State) GetId() *ids.Operation {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Operation_State) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Operation_State) GetPosition() *measurements.Position_Absolute {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Operation_State) GetVelocity() *measurements.Velocity_Polar {
	if m != nil {
		return m.Velocity
	}
	return nil
}

// Operator models the operator of an aircraft
type Operator struct {
	Id                   *ids.Operator     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location             *geo.Coordinate2D `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Operator) Reset()         { *m = Operator{} }
func (m *Operator) String() string { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()    {}
func (*Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{3}
}

func (m *Operator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operator.Unmarshal(m, b)
}
func (m *Operator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operator.Marshal(b, m, deterministic)
}
func (m *Operator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operator.Merge(m, src)
}
func (m *Operator) XXX_Size() int {
	return xxx_messageInfo_Operator.Size(m)
}
func (m *Operator) XXX_DiscardUnknown() {
	xxx_messageInfo_Operator.DiscardUnknown(m)
}

var xxx_messageInfo_Operator proto.InternalMessageInfo

func (m *Operator) GetId() *ids.Operator {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Operator) GetLocation() *geo.Coordinate2D {
	if m != nil {
		return m.Location
	}
	return nil
}

// Aircraft models the vehicle performing an operation.
type Aircraft struct {
	Type                 Aircraft_Type `protobuf:"varint,1,opt,name=type,proto3,enum=remoteid.Aircraft_Type" json:"type,omitempty"`
	Serial               string        `protobuf:"bytes,4,opt,name=serial,proto3" json:"serial,omitempty"`
	Registration         string        `protobuf:"bytes,5,opt,name=registration,proto3" json:"registration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Aircraft) Reset()         { *m = Aircraft{} }
func (m *Aircraft) String() string { return proto.CompactTextString(m) }
func (*Aircraft) ProtoMessage()    {}
func (*Aircraft) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ec874c18515d49f, []int{4}
}

func (m *Aircraft) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Aircraft.Unmarshal(m, b)
}
func (m *Aircraft) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Aircraft.Marshal(b, m, deterministic)
}
func (m *Aircraft) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aircraft.Merge(m, src)
}
func (m *Aircraft) XXX_Size() int {
	return xxx_messageInfo_Aircraft.Size(m)
}
func (m *Aircraft) XXX_DiscardUnknown() {
	xxx_messageInfo_Aircraft.DiscardUnknown(m)
}

var xxx_messageInfo_Aircraft proto.InternalMessageInfo

func (m *Aircraft) GetType() Aircraft_Type {
	if m != nil {
		return m.Type
	}
	return Aircraft_UNKNOWN
}

func (m *Aircraft) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *Aircraft) GetRegistration() string {
	if m != nil {
		return m.Registration
	}
	return ""
}

func init() {
	proto.RegisterEnum("remoteid.Aircraft_Type", Aircraft_Type_name, Aircraft_Type_value)
	proto.RegisterType((*MonitorAreaParameters)(nil), "remoteid.MonitorAreaParameters")
	proto.RegisterType((*MonitorAreaResponse)(nil), "remoteid.MonitorAreaResponse")
	proto.RegisterType((*MonitorAreaResponse_Added)(nil), "remoteid.MonitorAreaResponse.Added")
	proto.RegisterType((*MonitorAreaResponse_Updated)(nil), "remoteid.MonitorAreaResponse.Updated")
	proto.RegisterType((*MonitorAreaResponse_Removed)(nil), "remoteid.MonitorAreaResponse.Removed")
	proto.RegisterType((*Operation)(nil), "remoteid.Operation")
	proto.RegisterType((*Operation_State)(nil), "remoteid.Operation.State")
	proto.RegisterType((*Operator)(nil), "remoteid.Operator")
	proto.RegisterType((*Aircraft)(nil), "remoteid.Aircraft")
}

func init() { proto.RegisterFile("remoteid/remoteid.proto", fileDescriptor_6ec874c18515d49f) }

var fileDescriptor_6ec874c18515d49f = []byte{
	// 893 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xcf, 0x6f, 0xdb, 0x36,
	0x14, 0xc7, 0xe3, 0xdf, 0xd2, 0x73, 0xec, 0x68, 0x0c, 0xd6, 0x7a, 0x46, 0xbb, 0x04, 0x1e, 0x06,
	0x78, 0xd8, 0x2a, 0xad, 0xee, 0x0e, 0xc5, 0x7a, 0x92, 0x6d, 0x25, 0x36, 0xe2, 0x59, 0x06, 0xa3,
	0x64, 0xbf, 0x0e, 0x06, 0x63, 0xd1, 0x1a, 0x01, 0xcb, 0x14, 0x48, 0xba, 0x58, 0x80, 0xfd, 0x81,
	0xfb, 0x13, 0x76, 0xde, 0x79, 0xf7, 0x5d, 0x07, 0xfd, 0xb4, 0xb3, 0x06, 0x6d, 0x6f, 0x22, 0xdf,
	0xe7, 0xfb, 0xc8, 0xef, 0xe3, 0x23, 0x05, 0x4f, 0x05, 0x0d, 0xb9, 0xa2, 0xcc, 0xb7, 0xf2, 0x0f,
	0x33, 0x12, 0x5c, 0x71, 0xa4, 0xe5, 0xe3, 0x6e, 0x8b, 0xf9, 0xd2, 0x62, 0xbe, 0x4c, 0x03, 0xdd,
	0x56, 0x40, 0xb9, 0x15, 0x50, 0x9e, 0x0d, 0xcf, 0x42, 0x4a, 0xe4, 0x4e, 0xd0, 0x90, 0x6e, 0x95,
	0xb4, 0x0e, 0x07, 0x39, 0x10, 0x70, 0x1e, 0x6c, 0xa8, 0x95, 0x8c, 0xee, 0x76, 0x6b, 0x4b, 0xb1,
	0x90, 0x4a, 0x45, 0xc2, 0x28, 0x05, 0x7a, 0xd7, 0xf0, 0xe9, 0x0f, 0x7c, 0xcb, 0x14, 0x17, 0xb6,
	0xa0, 0x64, 0x41, 0x04, 0x09, 0xa9, 0xa2, 0x42, 0xa2, 0xef, 0xc1, 0x20, 0x82, 0x92, 0x25, 0x5f,
	0x2f, 0xd9, 0x56, 0x51, 0x41, 0xa5, 0xea, 0x94, 0xce, 0x4b, 0xfd, 0xe6, 0xc0, 0x30, 0xe3, 0x0d,
	0x0c, 0xf9, 0x6e, 0xeb, 0xb3, 0x6d, 0x30, 0xe4, 0xbf, 0xe3, 0x76, 0x4c, 0xba, 0xeb, 0x69, 0xc6,
	0xf5, 0xfe, 0xad, 0xc0, 0xe9, 0x41, 0x56, 0x4c, 0x65, 0xc4, 0xb7, 0x92, 0xa2, 0x37, 0x50, 0x23,
	0xbe, 0x4f, 0xfd, 0x2c, 0xd1, 0x17, 0x66, 0x61, 0xfb, 0x11, 0xda, 0xb4, 0x63, 0x74, 0x72, 0x84,
	0x53, 0x0d, 0xb2, 0xa1, 0xb1, 0x8b, 0x7c, 0xa2, 0xa8, 0xdf, 0x29, 0x27, 0xf2, 0x2f, 0xdf, 0x2f,
	0xbf, 0x49, 0xe1, 0xc9, 0x11, 0xce, 0x75, 0x71, 0x8a, 0x58, 0xf2, 0x96, 0xfa, 0x9d, 0xca, 0xc7,
	0xa4, 0xc0, 0x29, 0x1c, 0xa7, 0xc8, 0x74, 0x5d, 0x06, 0xb5, 0x64, 0x5f, 0xe8, 0x25, 0xe8, 0x3c,
	0xa2, 0x82, 0x28, 0xc6, 0xb7, 0x99, 0x9f, 0xd3, 0x7d, 0x36, 0x37, 0x0f, 0xe1, 0x3d, 0x85, 0x5e,
	0x80, 0x1e, 0x71, 0xc9, 0xe2, 0x6f, 0xd9, 0x29, 0x9f, 0x57, 0xfa, 0xcd, 0xc1, 0x49, 0x52, 0xcb,
	0x45, 0x36, 0xfb, 0xdd, 0x18, 0xef, 0x89, 0xee, 0x1f, 0xd0, 0xc8, 0x3c, 0xa0, 0x57, 0x1f, 0xb7,
	0xd8, 0xe4, 0xe8, 0x70, 0xb9, 0x97, 0x50, 0x93, 0x8a, 0x28, 0x9a, 0x95, 0xeb, 0xb3, 0x47, 0x04,
	0xe6, 0x75, 0x0c, 0xc4, 0x35, 0x4e, 0xc8, 0xa1, 0x0e, 0x0d, 0x9f, 0x2a, 0xc2, 0x36, 0xb2, 0xfb,
	0x15, 0x34, 0x32, 0xfb, 0xe8, 0x73, 0x28, 0xb3, 0xfc, 0xcc, 0xda, 0x66, 0xdc, 0x8c, 0x7b, 0x7b,
	0x65, 0xe6, 0x1f, 0xa8, 0x7a, 0x7f, 0x56, 0x40, 0x2f, 0x82, 0x1f, 0x12, 0x22, 0x13, 0xb4, 0x74,
	0xbb, 0x5c, 0x64, 0x9b, 0x44, 0xff, 0xdf, 0x24, 0x17, 0xb8, 0x60, 0x62, 0x9e, 0x30, 0xb1, 0x12,
	0x64, 0xad, 0xb2, 0x03, 0x3c, 0xe0, 0xed, 0x2c, 0x82, 0x0b, 0x06, 0xf5, 0x41, 0x0b, 0x28, 0x0f,
	0xa9, 0x12, 0xf7, 0x9d, 0x6a, 0xc2, 0x1f, 0x67, 0xf5, 0xde, 0xdc, 0x07, 0x7c, 0x8b, 0x8b, 0x28,
	0xea, 0x40, 0x23, 0xda, 0x89, 0x88, 0x4b, 0xda, 0xa9, 0x9d, 0x97, 0xfa, 0x3a, 0xce, 0x87, 0xe8,
	0x19, 0xe8, 0x92, 0x85, 0xbb, 0x4d, 0xd2, 0x78, 0xf5, 0xf3, 0x52, 0x5f, 0xc3, 0xfb, 0x89, 0xee,
	0x5f, 0x25, 0xa8, 0x25, 0x35, 0xfc, 0xa0, 0xd7, 0xd7, 0xa0, 0x17, 0x77, 0x2f, 0x33, 0xdb, 0x35,
	0xd3, 0xdb, 0x69, 0xe6, 0xb7, 0xd3, 0xf4, 0x72, 0x02, 0xef, 0x61, 0xf4, 0x06, 0xb4, 0xbc, 0x29,
	0x32, 0xd7, 0x67, 0xe6, 0x83, 0xab, 0x9e, 0xb7, 0x8f, 0x69, 0xdf, 0x49, 0xbe, 0xd9, 0x29, 0x8a,
	0x0b, 0x01, 0x7a, 0x0d, 0xda, 0x5b, 0xba, 0xe1, 0x2b, 0xa6, 0xf2, 0x12, 0x3c, 0x7b, 0x28, 0xbe,
	0xcd, 0xa2, 0x71, 0x51, 0x88, 0xc0, 0x05, 0xdd, 0xfb, 0x09, 0xb4, 0xfc, 0x08, 0xd0, 0xf3, 0x03,
	0x73, 0xad, 0x03, 0x73, 0x5c, 0x24, 0xde, 0x5e, 0x80, 0xb6, 0xe1, 0xab, 0xb4, 0x3b, 0x53, 0x6b,
	0x9f, 0x24, 0x75, 0x1e, 0x71, 0x2e, 0x7c, 0xb6, 0x25, 0x8a, 0x0e, 0xc6, 0xb8, 0x40, 0x7a, 0x7f,
	0x97, 0x41, 0xcb, 0x4f, 0x0b, 0x7d, 0x0d, 0x55, 0x75, 0x1f, 0xd1, 0x24, 0x79, 0x7b, 0xf0, 0xf4,
	0xdd, 0xf3, 0x34, 0xbd, 0xfb, 0x88, 0xe2, 0x04, 0x42, 0x4f, 0xa0, 0x2e, 0xa9, 0x60, 0x64, 0x93,
	0x78, 0xd1, 0x71, 0x36, 0x42, 0x3d, 0x38, 0x16, 0x34, 0x60, 0x52, 0x65, 0x57, 0x24, 0x3d, 0xc3,
	0x07, 0x73, 0xbd, 0x7f, 0x4a, 0x50, 0x8d, 0x53, 0xa1, 0x26, 0x34, 0x6e, 0xe6, 0x57, 0x73, 0xf7,
	0xc7, 0xb9, 0x71, 0x84, 0x8e, 0x41, 0xb3, 0xa7, 0x78, 0x31, 0xb3, 0xe7, 0x8e, 0x51, 0x42, 0x6d,
	0x00, 0xec, 0x7a, 0x2e, 0x1e, 0x61, 0xfb, 0xc2, 0x33, 0xca, 0xa8, 0x05, 0xfa, 0xe5, 0xcf, 0xd8,
	0x4d, 0xc3, 0x15, 0xa4, 0x41, 0xf5, 0xd6, 0x73, 0x67, 0x46, 0x15, 0x9d, 0x40, 0xd3, 0xc5, 0xf3,
	0xa9, 0x37, 0x71, 0x17, 0x9e, 0x83, 0x8d, 0x1a, 0x02, 0xa8, 0x5f, 0xce, 0xa6, 0x63, 0x07, 0x1b,
	0xf5, 0x18, 0xbb, 0x9a, 0x7a, 0x8e, 0xd1, 0x88, 0xb1, 0x0b, 0xec, 0x38, 0xcb, 0xa1, 0x3d, 0x73,
	0xdd, 0xb9, 0xa1, 0xa1, 0x53, 0x38, 0x19, 0xd9, 0x0b, 0x6f, 0x7a, 0x9b, 0xcc, 0x25, 0x93, 0x7a,
	0xbc, 0x21, 0x7b, 0x8a, 0xaf, 0x27, 0xd3, 0x85, 0x01, 0xa8, 0x0b, 0x4f, 0x12, 0xc9, 0x85, 0x3d,
	0x9b, 0x2d, 0x5d, 0xbc, 0x5c, 0xd8, 0xd8, 0x1e, 0x4d, 0x6e, 0x3c, 0xc7, 0x68, 0xc6, 0x8b, 0x60,
	0x77, 0x74, 0xe5, 0x78, 0xc6, 0x71, 0x9c, 0xe9, 0x12, 0xbb, 0x37, 0xf3, 0xf1, 0xd2, 0x1d, 0x5e,
	0x7b, 0xf6, 0x68, 0xe6, 0x18, 0x2d, 0xa4, 0x43, 0xcd, 0xf5, 0x26, 0x0e, 0x36, 0xda, 0x83, 0x5f,
	0x41, 0xc3, 0x49, 0x29, 0xa7, 0x63, 0xe4, 0x42, 0xf3, 0xe0, 0x79, 0x43, 0x67, 0x8f, 0xbe, 0x7a,
	0xfb, 0xb7, 0xbf, 0xfb, 0xfc, 0xbd, 0xcf, 0xe2, 0xb7, 0xa5, 0xa1, 0xf9, 0xcb, 0x37, 0x01, 0x53,
	0xbf, 0xed, 0xee, 0xcc, 0x15, 0x0f, 0x2d, 0xc2, 0x44, 0x48, 0x22, 0x2b, 0xf9, 0x4d, 0xac, 0xc9,
	0x8a, 0x4a, 0x4b, 0x8a, 0x95, 0x15, 0xf0, 0xe2, 0xaf, 0x76, 0x57, 0x4f, 0x3a, 0xfc, 0xd5, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x66, 0x72, 0xff, 0xf1, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemoteIDClient is the client API for RemoteID service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteIDClient interface {
	// ALPHA: API subject to change
	// MonitorArea monitors a given area for active operations.
	MonitorArea(ctx context.Context, in *MonitorAreaParameters, opts ...grpc.CallOption) (RemoteID_MonitorAreaClient, error)
}

type remoteIDClient struct {
	cc *grpc.ClientConn
}

func NewRemoteIDClient(cc *grpc.ClientConn) RemoteIDClient {
	return &remoteIDClient{cc}
}

func (c *remoteIDClient) MonitorArea(ctx context.Context, in *MonitorAreaParameters, opts ...grpc.CallOption) (RemoteID_MonitorAreaClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RemoteID_serviceDesc.Streams[0], "/remoteid.RemoteID/MonitorArea", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteIDMonitorAreaClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteID_MonitorAreaClient interface {
	Recv() (*MonitorAreaResponse, error)
	grpc.ClientStream
}

type remoteIDMonitorAreaClient struct {
	grpc.ClientStream
}

func (x *remoteIDMonitorAreaClient) Recv() (*MonitorAreaResponse, error) {
	m := new(MonitorAreaResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RemoteIDServer is the server API for RemoteID service.
type RemoteIDServer interface {
	// ALPHA: API subject to change
	// MonitorArea monitors a given area for active operations.
	MonitorArea(*MonitorAreaParameters, RemoteID_MonitorAreaServer) error
}

// UnimplementedRemoteIDServer can be embedded to have forward compatible implementations.
type UnimplementedRemoteIDServer struct {
}

func (*UnimplementedRemoteIDServer) MonitorArea(req *MonitorAreaParameters, srv RemoteID_MonitorAreaServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorArea not implemented")
}

func RegisterRemoteIDServer(s *grpc.Server, srv RemoteIDServer) {
	s.RegisterService(&_RemoteID_serviceDesc, srv)
}

func _RemoteID_MonitorArea_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorAreaParameters)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteIDServer).MonitorArea(m, &remoteIDMonitorAreaServer{stream})
}

type RemoteID_MonitorAreaServer interface {
	Send(*MonitorAreaResponse) error
	grpc.ServerStream
}

type remoteIDMonitorAreaServer struct {
	grpc.ServerStream
}

func (x *remoteIDMonitorAreaServer) Send(m *MonitorAreaResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RemoteID_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remoteid.RemoteID",
	HandlerType: (*RemoteIDServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorArea",
			Handler:       _RemoteID_MonitorArea_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remoteid/remoteid.proto",
}
