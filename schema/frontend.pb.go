// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frontend.proto

package schema

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
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

// GetServicesRequest is a request body for GetServices method
type GetServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesRequest) Reset()         { *m = GetServicesRequest{} }
func (m *GetServicesRequest) String() string { return proto.CompactTextString(m) }
func (*GetServicesRequest) ProtoMessage()    {}
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0}
}

func (m *GetServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesRequest.Unmarshal(m, b)
}
func (m *GetServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesRequest.Marshal(b, m, deterministic)
}
func (m *GetServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesRequest.Merge(m, src)
}
func (m *GetServicesRequest) XXX_Size() int {
	return xxx_messageInfo_GetServicesRequest.Size(m)
}
func (m *GetServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesRequest proto.InternalMessageInfo

// GetServicesResponse is a response body for GetServices method
type GetServicesResponse struct {
	// service_types - services that ever pushed data to this memprofiler
	ServiceTypes         []string `protobuf:"bytes,1,rep,name=service_types,json=serviceTypes,proto3" json:"service_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesResponse) Reset()         { *m = GetServicesResponse{} }
func (m *GetServicesResponse) String() string { return proto.CompactTextString(m) }
func (*GetServicesResponse) ProtoMessage()    {}
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{1}
}

func (m *GetServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesResponse.Unmarshal(m, b)
}
func (m *GetServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesResponse.Marshal(b, m, deterministic)
}
func (m *GetServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesResponse.Merge(m, src)
}
func (m *GetServicesResponse) XXX_Size() int {
	return xxx_messageInfo_GetServicesResponse.Size(m)
}
func (m *GetServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesResponse proto.InternalMessageInfo

func (m *GetServicesResponse) GetServiceTypes() []string {
	if m != nil {
		return m.ServiceTypes
	}
	return nil
}

// GetInstancesRequest is a request body for GetInstances method
type GetInstancesRequest struct {
	// service_type - identifier for a group of similar services
	ServiceType          string   `protobuf:"bytes,1,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInstancesRequest) Reset()         { *m = GetInstancesRequest{} }
func (m *GetInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstancesRequest) ProtoMessage()    {}
func (*GetInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{2}
}

func (m *GetInstancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstancesRequest.Unmarshal(m, b)
}
func (m *GetInstancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstancesRequest.Marshal(b, m, deterministic)
}
func (m *GetInstancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstancesRequest.Merge(m, src)
}
func (m *GetInstancesRequest) XXX_Size() int {
	return xxx_messageInfo_GetInstancesRequest.Size(m)
}
func (m *GetInstancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstancesRequest proto.InternalMessageInfo

func (m *GetInstancesRequest) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

// GetInstancesResponse is a response body for GetInstances method
type GetInstancesResponse struct {
	// service_instances - list of a particular kind of services instances
	ServiceInstances     []*ServiceDescription `protobuf:"bytes,1,rep,name=service_instances,json=serviceInstances,proto3" json:"service_instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetInstancesResponse) Reset()         { *m = GetInstancesResponse{} }
func (m *GetInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*GetInstancesResponse) ProtoMessage()    {}
func (*GetInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{3}
}

func (m *GetInstancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstancesResponse.Unmarshal(m, b)
}
func (m *GetInstancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstancesResponse.Marshal(b, m, deterministic)
}
func (m *GetInstancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstancesResponse.Merge(m, src)
}
func (m *GetInstancesResponse) XXX_Size() int {
	return xxx_messageInfo_GetInstancesResponse.Size(m)
}
func (m *GetInstancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstancesResponse proto.InternalMessageInfo

func (m *GetInstancesResponse) GetServiceInstances() []*ServiceDescription {
	if m != nil {
		return m.ServiceInstances
	}
	return nil
}

// GetSessionsRequest is a request body for GetSessions method
type GetSessionsRequest struct {
	// service_description - service instance information
	ServiceDescription   *ServiceDescription `protobuf:"bytes,1,opt,name=service_description,json=serviceDescription,proto3" json:"service_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetSessionsRequest) Reset()         { *m = GetSessionsRequest{} }
func (m *GetSessionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionsRequest) ProtoMessage()    {}
func (*GetSessionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{4}
}

func (m *GetSessionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionsRequest.Unmarshal(m, b)
}
func (m *GetSessionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionsRequest.Marshal(b, m, deterministic)
}
func (m *GetSessionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionsRequest.Merge(m, src)
}
func (m *GetSessionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionsRequest.Size(m)
}
func (m *GetSessionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionsRequest proto.InternalMessageInfo

func (m *GetSessionsRequest) GetServiceDescription() *ServiceDescription {
	if m != nil {
		return m.ServiceDescription
	}
	return nil
}

// GetSessionsResponse is a response body for GetSessions method
type GetSessionsResponse struct {
	// sessions is a list of  information about available sessions
	Sessions             []*Session `protobuf:"bytes,1,rep,name=sessions,proto3" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetSessionsResponse) Reset()         { *m = GetSessionsResponse{} }
func (m *GetSessionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSessionsResponse) ProtoMessage()    {}
func (*GetSessionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{5}
}

func (m *GetSessionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionsResponse.Unmarshal(m, b)
}
func (m *GetSessionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionsResponse.Marshal(b, m, deterministic)
}
func (m *GetSessionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionsResponse.Merge(m, src)
}
func (m *GetSessionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetSessionsResponse.Size(m)
}
func (m *GetSessionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionsResponse proto.InternalMessageInfo

func (m *GetSessionsResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

// SubscribeForSessionRequest is a request body for SubscribeForSession request
type SubscribeForSessionRequest struct {
	SessionDescription   *SessionDescription `protobuf:"bytes,1,opt,name=session_description,json=sessionDescription,proto3" json:"session_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SubscribeForSessionRequest) Reset()         { *m = SubscribeForSessionRequest{} }
func (m *SubscribeForSessionRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeForSessionRequest) ProtoMessage()    {}
func (*SubscribeForSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{6}
}

func (m *SubscribeForSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeForSessionRequest.Unmarshal(m, b)
}
func (m *SubscribeForSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeForSessionRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeForSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeForSessionRequest.Merge(m, src)
}
func (m *SubscribeForSessionRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeForSessionRequest.Size(m)
}
func (m *SubscribeForSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeForSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeForSessionRequest proto.InternalMessageInfo

func (m *SubscribeForSessionRequest) GetSessionDescription() *SessionDescription {
	if m != nil {
		return m.SessionDescription
	}
	return nil
}

// MemoryUtilizationRate is a collection of rate values for memory consumption indicators.
// Formally, the rate (or velocity) is the first time derivative of any memory consumption indicator.
// For Bytes rate units are bytes per second, for Objects rate units are units per second
type MemoryUtilizationRate struct {
	// span is a time span that is used to compute rates
	Span *duration.Duration `protobuf:"bytes,1,opt,name=span,proto3" json:"span,omitempty"`
	// values contains actual rates for a specified time span
	Values               *MemoryUtilizationRate_Values `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MemoryUtilizationRate) Reset()         { *m = MemoryUtilizationRate{} }
func (m *MemoryUtilizationRate) String() string { return proto.CompactTextString(m) }
func (*MemoryUtilizationRate) ProtoMessage()    {}
func (*MemoryUtilizationRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{7}
}

func (m *MemoryUtilizationRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryUtilizationRate.Unmarshal(m, b)
}
func (m *MemoryUtilizationRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryUtilizationRate.Marshal(b, m, deterministic)
}
func (m *MemoryUtilizationRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryUtilizationRate.Merge(m, src)
}
func (m *MemoryUtilizationRate) XXX_Size() int {
	return xxx_messageInfo_MemoryUtilizationRate.Size(m)
}
func (m *MemoryUtilizationRate) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryUtilizationRate.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryUtilizationRate proto.InternalMessageInfo

func (m *MemoryUtilizationRate) GetSpan() *duration.Duration {
	if m != nil {
		return m.Span
	}
	return nil
}

func (m *MemoryUtilizationRate) GetValues() *MemoryUtilizationRate_Values {
	if m != nil {
		return m.Values
	}
	return nil
}

// Values is a set of rate values
type MemoryUtilizationRate_Values struct {
	AllocObjects         float64  `protobuf:"fixed64,1,opt,name=alloc_objects,json=allocObjects,proto3" json:"alloc_objects,omitempty"`
	AllocBytes           float64  `protobuf:"fixed64,2,opt,name=alloc_bytes,json=allocBytes,proto3" json:"alloc_bytes,omitempty"`
	FreeObjects          float64  `protobuf:"fixed64,3,opt,name=free_objects,json=freeObjects,proto3" json:"free_objects,omitempty"`
	FreeBytes            float64  `protobuf:"fixed64,4,opt,name=free_bytes,json=freeBytes,proto3" json:"free_bytes,omitempty"`
	InUseObjects         float64  `protobuf:"fixed64,5,opt,name=in_use_objects,json=inUseObjects,proto3" json:"in_use_objects,omitempty"`
	InUseBytes           float64  `protobuf:"fixed64,6,opt,name=in_use_bytes,json=inUseBytes,proto3" json:"in_use_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemoryUtilizationRate_Values) Reset()         { *m = MemoryUtilizationRate_Values{} }
func (m *MemoryUtilizationRate_Values) String() string { return proto.CompactTextString(m) }
func (*MemoryUtilizationRate_Values) ProtoMessage()    {}
func (*MemoryUtilizationRate_Values) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{7, 0}
}

func (m *MemoryUtilizationRate_Values) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryUtilizationRate_Values.Unmarshal(m, b)
}
func (m *MemoryUtilizationRate_Values) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryUtilizationRate_Values.Marshal(b, m, deterministic)
}
func (m *MemoryUtilizationRate_Values) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryUtilizationRate_Values.Merge(m, src)
}
func (m *MemoryUtilizationRate_Values) XXX_Size() int {
	return xxx_messageInfo_MemoryUtilizationRate_Values.Size(m)
}
func (m *MemoryUtilizationRate_Values) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryUtilizationRate_Values.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryUtilizationRate_Values proto.InternalMessageInfo

func (m *MemoryUtilizationRate_Values) GetAllocObjects() float64 {
	if m != nil {
		return m.AllocObjects
	}
	return 0
}

func (m *MemoryUtilizationRate_Values) GetAllocBytes() float64 {
	if m != nil {
		return m.AllocBytes
	}
	return 0
}

func (m *MemoryUtilizationRate_Values) GetFreeObjects() float64 {
	if m != nil {
		return m.FreeObjects
	}
	return 0
}

func (m *MemoryUtilizationRate_Values) GetFreeBytes() float64 {
	if m != nil {
		return m.FreeBytes
	}
	return 0
}

func (m *MemoryUtilizationRate_Values) GetInUseObjects() float64 {
	if m != nil {
		return m.InUseObjects
	}
	return 0
}

func (m *MemoryUtilizationRate_Values) GetInUseBytes() float64 {
	if m != nil {
		return m.InUseBytes
	}
	return 0
}

// LocationMetrics is a set of memory allocation statistics
// that happened on a particular line of source code
type LocationMetrics struct {
	// rates represents memory consumption rates estimated
	// for some averaging window defined by server
	Rates []*MemoryUtilizationRate `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty"`
	// callstack describes location in code where the allocation happened
	Callstack            *Callstack `protobuf:"bytes,2,opt,name=callstack,proto3" json:"callstack,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LocationMetrics) Reset()         { *m = LocationMetrics{} }
func (m *LocationMetrics) String() string { return proto.CompactTextString(m) }
func (*LocationMetrics) ProtoMessage()    {}
func (*LocationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{8}
}

func (m *LocationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationMetrics.Unmarshal(m, b)
}
func (m *LocationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationMetrics.Marshal(b, m, deterministic)
}
func (m *LocationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationMetrics.Merge(m, src)
}
func (m *LocationMetrics) XXX_Size() int {
	return xxx_messageInfo_LocationMetrics.Size(m)
}
func (m *LocationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_LocationMetrics proto.InternalMessageInfo

func (m *LocationMetrics) GetRates() []*MemoryUtilizationRate {
	if m != nil {
		return m.Rates
	}
	return nil
}

func (m *LocationMetrics) GetCallstack() *Callstack {
	if m != nil {
		return m.Callstack
	}
	return nil
}

// SessionMetrics contains list of heap allocation metrics per every location
type SessionMetrics struct {
	Locations            []*LocationMetrics `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SessionMetrics) Reset()         { *m = SessionMetrics{} }
func (m *SessionMetrics) String() string { return proto.CompactTextString(m) }
func (*SessionMetrics) ProtoMessage()    {}
func (*SessionMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{9}
}

func (m *SessionMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionMetrics.Unmarshal(m, b)
}
func (m *SessionMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionMetrics.Marshal(b, m, deterministic)
}
func (m *SessionMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionMetrics.Merge(m, src)
}
func (m *SessionMetrics) XXX_Size() int {
	return xxx_messageInfo_SessionMetrics.Size(m)
}
func (m *SessionMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_SessionMetrics proto.InternalMessageInfo

func (m *SessionMetrics) GetLocations() []*LocationMetrics {
	if m != nil {
		return m.Locations
	}
	return nil
}

func init() {
	proto.RegisterType((*GetServicesRequest)(nil), "schema.GetServicesRequest")
	proto.RegisterType((*GetServicesResponse)(nil), "schema.GetServicesResponse")
	proto.RegisterType((*GetInstancesRequest)(nil), "schema.GetInstancesRequest")
	proto.RegisterType((*GetInstancesResponse)(nil), "schema.GetInstancesResponse")
	proto.RegisterType((*GetSessionsRequest)(nil), "schema.GetSessionsRequest")
	proto.RegisterType((*GetSessionsResponse)(nil), "schema.GetSessionsResponse")
	proto.RegisterType((*SubscribeForSessionRequest)(nil), "schema.SubscribeForSessionRequest")
	proto.RegisterType((*MemoryUtilizationRate)(nil), "schema.MemoryUtilizationRate")
	proto.RegisterType((*MemoryUtilizationRate_Values)(nil), "schema.MemoryUtilizationRate.Values")
	proto.RegisterType((*LocationMetrics)(nil), "schema.LocationMetrics")
	proto.RegisterType((*SessionMetrics)(nil), "schema.SessionMetrics")
}

func init() { proto.RegisterFile("frontend.proto", fileDescriptor_eca3873955a29cfe) }

var fileDescriptor_eca3873955a29cfe = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x4f, 0xd4, 0x40,
	0x14, 0xa5, 0x7c, 0x6c, 0xdc, 0xbb, 0x0b, 0xc8, 0x2c, 0x2a, 0x56, 0x51, 0x1c, 0x79, 0x20, 0x31,
	0x16, 0x03, 0x31, 0x31, 0xc6, 0x27, 0x24, 0xa0, 0x41, 0x62, 0x32, 0x88, 0xaf, 0x9b, 0x69, 0x99,
	0xc5, 0x91, 0x6e, 0xa7, 0xce, 0x4c, 0x31, 0xf8, 0x1f, 0x7d, 0xf7, 0xcd, 0xbf, 0x62, 0x3a, 0x1f,
	0x6d, 0xb7, 0xac, 0xfb, 0xd8, 0xd3, 0x73, 0xce, 0x3d, 0x33, 0xb9, 0x67, 0x60, 0x65, 0x24, 0x45,
	0xa6, 0x59, 0x76, 0x11, 0xe5, 0x52, 0x68, 0x81, 0x3a, 0x2a, 0xf9, 0xc6, 0xc6, 0x34, 0xec, 0x27,
	0x62, 0x3c, 0x16, 0x99, 0x45, 0xc3, 0xe5, 0x98, 0x26, 0x57, 0x15, 0x29, 0x7c, 0x72, 0x29, 0xc4,
	0x65, 0xca, 0x76, 0xcd, 0x57, 0x5c, 0x8c, 0x76, 0x2f, 0x0a, 0x49, 0x35, 0xf7, 0x74, 0xbc, 0x0e,
	0xe8, 0x98, 0xe9, 0x33, 0x26, 0xaf, 0x79, 0xc2, 0x14, 0x61, 0x3f, 0x0a, 0xa6, 0x34, 0x7e, 0x0b,
	0x83, 0x09, 0x54, 0xe5, 0x22, 0x53, 0x0c, 0x3d, 0x87, 0x65, 0x65, 0xb1, 0xa1, 0xbe, 0xc9, 0x99,
	0xda, 0x08, 0xb6, 0x16, 0x76, 0xba, 0xa4, 0xef, 0xc0, 0x2f, 0x25, 0x86, 0xdf, 0x18, 0xed, 0xc7,
	0x4c, 0x69, 0x9a, 0xd5, 0x96, 0xe8, 0x19, 0xf4, 0x9b, 0xda, 0x8d, 0x60, 0x2b, 0xd8, 0xe9, 0x92,
	0x5e, 0x43, 0x8a, 0x87, 0xb0, 0x3e, 0xa9, 0x74, 0x63, 0x8f, 0x61, 0xcd, 0x4b, 0xb9, 0xff, 0x69,
	0x46, 0xf7, 0xf6, 0xc2, 0xc8, 0x5e, 0x42, 0xe4, 0xb2, 0x1e, 0x32, 0x95, 0x48, 0x9e, 0x97, 0x07,
	0x24, 0x77, 0x9d, 0xa8, 0x32, 0xc4, 0xd4, 0x1d, 0x56, 0x29, 0x2e, 0xb2, 0x2a, 0xd9, 0x09, 0x0c,
	0xbc, 0xfd, 0x45, 0x2d, 0x37, 0x01, 0x67, 0x0f, 0x40, 0xea, 0x16, 0x86, 0x0f, 0xdc, 0xcd, 0xf9,
	0x11, 0xee, 0x08, 0x2f, 0xe0, 0x8e, 0x72, 0x98, 0x4b, 0xbe, 0x5a, 0x1b, 0x1b, 0x9c, 0x54, 0x04,
	0xcc, 0x21, 0x3c, 0x2b, 0xe2, 0xd2, 0x33, 0x66, 0x47, 0x42, 0x7a, 0x42, 0x33, 0xae, 0x41, 0x66,
	0xc7, 0x35, 0x94, 0x56, 0xdc, 0x36, 0x86, 0xff, 0xce, 0xc3, 0xbd, 0x53, 0x36, 0x16, 0xf2, 0xe6,
	0x5c, 0xf3, 0x94, 0xff, 0x32, 0xab, 0x41, 0xa8, 0x66, 0xe8, 0x25, 0x2c, 0xaa, 0x9c, 0x7a, 0xdf,
	0x87, 0x91, 0xdd, 0xa3, 0xc8, 0xef, 0x51, 0x74, 0xe8, 0xf6, 0x88, 0x18, 0x1a, 0x7a, 0x07, 0x9d,
	0x6b, 0x9a, 0x16, 0x4c, 0x6d, 0xcc, 0x1b, 0xc1, 0xb6, 0x0f, 0x32, 0xd5, 0x3d, 0xfa, 0x6a, 0xb8,
	0xc4, 0x69, 0xc2, 0x3f, 0x01, 0x74, 0x2c, 0x54, 0xee, 0x18, 0x4d, 0x53, 0x91, 0x0c, 0x45, 0xfc,
	0x9d, 0x25, 0x5a, 0x99, 0x00, 0x01, 0xe9, 0x1b, 0xf0, 0xb3, 0xc5, 0xd0, 0x53, 0xe8, 0x59, 0x52,
	0x7c, 0xa3, 0xdd, 0xc8, 0x80, 0x80, 0x81, 0x0e, 0x4a, 0xa4, 0xdc, 0xb6, 0x91, 0x64, 0xac, 0x32,
	0x59, 0x30, 0x8c, 0x5e, 0x89, 0x79, 0x8f, 0x4d, 0x00, 0x43, 0xb1, 0x16, 0x8b, 0x86, 0xd0, 0x2d,
	0x11, 0xeb, 0xb0, 0x0d, 0x2b, 0x3c, 0x1b, 0x16, 0xaa, 0xf6, 0x58, 0xb2, 0x41, 0x78, 0x76, 0xae,
	0x2a, 0x93, 0x2d, 0xe8, 0x3b, 0x96, 0xb5, 0xe9, 0xd8, 0x24, 0x86, 0x63, 0x7c, 0xf0, 0x4f, 0x58,
	0xfd, 0x24, 0x12, 0x73, 0xf2, 0x53, 0xa6, 0x25, 0x4f, 0x14, 0xda, 0x87, 0x25, 0x49, 0x75, 0xb5,
	0xc3, 0x9b, 0x33, 0xaf, 0x8a, 0x58, 0x2e, 0xda, 0x85, 0x6e, 0x42, 0xd3, 0x54, 0x69, 0x9a, 0x5c,
	0xb9, 0x3b, 0x5e, 0xf3, 0xc2, 0xf7, 0xfe, 0x07, 0xa9, 0x39, 0xf8, 0x18, 0x56, 0xdc, 0x12, 0xf8,
	0xb9, 0xaf, 0xa1, 0x9b, 0xba, 0x28, 0x7e, 0xf6, 0x03, 0x6f, 0xd1, 0xca, 0x48, 0x6a, 0xe6, 0xde,
	0xef, 0x79, 0x18, 0x9c, 0xb2, 0x71, 0x2e, 0xc5, 0x88, 0xa7, 0x4c, 0x1e, 0xb9, 0x57, 0x08, 0x7d,
	0x80, 0x5e, 0xe3, 0x91, 0x40, 0xd5, 0xea, 0xdd, 0x7e, 0x4f, 0xc2, 0x47, 0x53, 0xff, 0xd9, 0x6e,
	0xe0, 0x39, 0x74, 0x02, 0xfd, 0x66, 0xf1, 0x51, 0x93, 0xde, 0x7e, 0x48, 0xc2, 0xc7, 0xd3, 0x7f,
	0x56, 0x66, 0x3e, 0x96, 0x2d, 0x53, 0x2b, 0xd6, 0x44, 0xf3, 0x5b, 0xb1, 0x26, 0x2b, 0x8b, 0xe7,
	0xd0, 0x39, 0x0c, 0xa6, 0xf4, 0x10, 0xe1, 0xaa, 0x63, 0xff, 0x2d, 0x69, 0x78, 0xbf, 0xd5, 0x43,
	0x77, 0xad, 0x78, 0xee, 0x55, 0x10, 0x77, 0x4c, 0x87, 0xf6, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x78, 0xaa, 0x79, 0xa7, 0xd0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MemprofilerFrontendClient is the client Backend for MemprofilerFrontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemprofilerFrontendClient interface {
	// GetServices returns the list of registered services
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	// GetInstances returns the list of service instances
	GetInstances(ctx context.Context, in *GetInstancesRequest, opts ...grpc.CallOption) (*GetInstancesResponse, error)
	// GetSessions returns the list of available profiling sessions for a service instance
	GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error)
	// SubscribeForSession returns the stream of session updates with fresh trend values, if it's still alive
	SubscribeForSession(ctx context.Context, in *SubscribeForSessionRequest, opts ...grpc.CallOption) (MemprofilerFrontend_SubscribeForSessionClient, error)
}

type memprofilerFrontendClient struct {
	cc *grpc.ClientConn
}

func NewMemprofilerFrontendClient(cc *grpc.ClientConn) MemprofilerFrontendClient {
	return &memprofilerFrontendClient{cc}
}

func (c *memprofilerFrontendClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, "/schema.MemprofilerFrontend/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memprofilerFrontendClient) GetInstances(ctx context.Context, in *GetInstancesRequest, opts ...grpc.CallOption) (*GetInstancesResponse, error) {
	out := new(GetInstancesResponse)
	err := c.cc.Invoke(ctx, "/schema.MemprofilerFrontend/GetInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memprofilerFrontendClient) GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error) {
	out := new(GetSessionsResponse)
	err := c.cc.Invoke(ctx, "/schema.MemprofilerFrontend/GetSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memprofilerFrontendClient) SubscribeForSession(ctx context.Context, in *SubscribeForSessionRequest, opts ...grpc.CallOption) (MemprofilerFrontend_SubscribeForSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MemprofilerFrontend_serviceDesc.Streams[0], "/schema.MemprofilerFrontend/SubscribeForSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &memprofilerFrontendSubscribeForSessionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MemprofilerFrontend_SubscribeForSessionClient interface {
	Recv() (*SessionMetrics, error)
	grpc.ClientStream
}

type memprofilerFrontendSubscribeForSessionClient struct {
	grpc.ClientStream
}

func (x *memprofilerFrontendSubscribeForSessionClient) Recv() (*SessionMetrics, error) {
	m := new(SessionMetrics)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MemprofilerFrontendServer is the server Backend for MemprofilerFrontend service.
type MemprofilerFrontendServer interface {
	// GetServices returns the list of registered services
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	// GetInstances returns the list of service instances
	GetInstances(context.Context, *GetInstancesRequest) (*GetInstancesResponse, error)
	// GetSessions returns the list of available profiling sessions for a service instance
	GetSessions(context.Context, *GetSessionsRequest) (*GetSessionsResponse, error)
	// SubscribeForSession returns the stream of session updates with fresh trend values, if it's still alive
	SubscribeForSession(*SubscribeForSessionRequest, MemprofilerFrontend_SubscribeForSessionServer) error
}

func RegisterMemprofilerFrontendServer(s *grpc.Server, srv MemprofilerFrontendServer) {
	s.RegisterService(&_MemprofilerFrontend_serviceDesc, srv)
}

func _MemprofilerFrontend_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemprofilerFrontendServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.MemprofilerFrontend/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemprofilerFrontendServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemprofilerFrontend_GetInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemprofilerFrontendServer).GetInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.MemprofilerFrontend/GetInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemprofilerFrontendServer).GetInstances(ctx, req.(*GetInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemprofilerFrontend_GetSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemprofilerFrontendServer).GetSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.MemprofilerFrontend/GetSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemprofilerFrontendServer).GetSessions(ctx, req.(*GetSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemprofilerFrontend_SubscribeForSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeForSessionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MemprofilerFrontendServer).SubscribeForSession(m, &memprofilerFrontendSubscribeForSessionServer{stream})
}

type MemprofilerFrontend_SubscribeForSessionServer interface {
	Send(*SessionMetrics) error
	grpc.ServerStream
}

type memprofilerFrontendSubscribeForSessionServer struct {
	grpc.ServerStream
}

func (x *memprofilerFrontendSubscribeForSessionServer) Send(m *SessionMetrics) error {
	return x.ServerStream.SendMsg(m)
}

var _MemprofilerFrontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.MemprofilerFrontend",
	HandlerType: (*MemprofilerFrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _MemprofilerFrontend_GetServices_Handler,
		},
		{
			MethodName: "GetInstances",
			Handler:    _MemprofilerFrontend_GetInstances_Handler,
		},
		{
			MethodName: "GetSessions",
			Handler:    _MemprofilerFrontend_GetSessions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeForSession",
			Handler:       _MemprofilerFrontend_SubscribeForSession_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "frontend.proto",
}
