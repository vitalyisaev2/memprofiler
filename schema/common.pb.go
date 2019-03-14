// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package schema

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// ServiceDescription describes a service whose memory stats is being tracked
type ServiceDescription struct {
	// type - general service description (kind, role and so on)
	ServiceType string `protobuf:"bytes,1,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	// instance - unique id of a particular service instance (it may be IP or node ID, and so on)
	ServiceInstance      string   `protobuf:"bytes,2,opt,name=service_instance,json=serviceInstance,proto3" json:"service_instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceDescription) Reset()         { *m = ServiceDescription{} }
func (m *ServiceDescription) String() string { return proto.CompactTextString(m) }
func (*ServiceDescription) ProtoMessage()    {}
func (*ServiceDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *ServiceDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceDescription.Unmarshal(m, b)
}
func (m *ServiceDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceDescription.Marshal(b, m, deterministic)
}
func (m *ServiceDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceDescription.Merge(m, src)
}
func (m *ServiceDescription) XXX_Size() int {
	return xxx_messageInfo_ServiceDescription.Size(m)
}
func (m *ServiceDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceDescription proto.InternalMessageInfo

func (m *ServiceDescription) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *ServiceDescription) GetServiceInstance() string {
	if m != nil {
		return m.ServiceInstance
	}
	return ""
}

// SessionDescription identifies a single memory tracking session for
// a particular service instance. Typically every service restart
// terminates current session (if there is any) and starts new one.
type SessionDescription struct {
	// type - general service description (kind, role and so on)
	ServiceType string `protobuf:"bytes,1,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	// instance - description of a particular service instance
	ServiceInstance string `protobuf:"bytes,2,opt,name=service_instance,json=serviceInstance,proto3" json:"service_instance,omitempty"`
	// session_id - unique session identifier
	SessionId            uint32   `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionDescription) Reset()         { *m = SessionDescription{} }
func (m *SessionDescription) String() string { return proto.CompactTextString(m) }
func (*SessionDescription) ProtoMessage()    {}
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *SessionDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDescription.Unmarshal(m, b)
}
func (m *SessionDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDescription.Marshal(b, m, deterministic)
}
func (m *SessionDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDescription.Merge(m, src)
}
func (m *SessionDescription) XXX_Size() int {
	return xxx_messageInfo_SessionDescription.Size(m)
}
func (m *SessionDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDescription proto.InternalMessageInfo

func (m *SessionDescription) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *SessionDescription) GetServiceInstance() string {
	if m != nil {
		return m.ServiceInstance
	}
	return ""
}

func (m *SessionDescription) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

// SessionMetadata contains server-side session metadata
type SessionMetadata struct {
	// started_at - time when session has started (UTC)
	StartedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// finished_at - time when session has stopped (UTC), may be empty if session is still alive
	FinishedAt           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SessionMetadata) Reset()         { *m = SessionMetadata{} }
func (m *SessionMetadata) String() string { return proto.CompactTextString(m) }
func (*SessionMetadata) ProtoMessage()    {}
func (*SessionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *SessionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionMetadata.Unmarshal(m, b)
}
func (m *SessionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionMetadata.Marshal(b, m, deterministic)
}
func (m *SessionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionMetadata.Merge(m, src)
}
func (m *SessionMetadata) XXX_Size() int {
	return xxx_messageInfo_SessionMetadata.Size(m)
}
func (m *SessionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_SessionMetadata proto.InternalMessageInfo

func (m *SessionMetadata) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *SessionMetadata) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

// Session combines all available information about a memory tracking session
type Session struct {
	Description          *SessionDescription `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Metadata             *SessionMetadata    `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetDescription() *SessionDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Session) GetMetadata() *SessionMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceDescription)(nil), "schema.ServiceDescription")
	proto.RegisterType((*SessionDescription)(nil), "schema.SessionDescription")
	proto.RegisterType((*SessionMetadata)(nil), "schema.SessionMetadata")
	proto.RegisterType((*Session)(nil), "schema.Session")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xbf, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x7e, 0x14, 0x72, 0x29, 0x2a, 0xf2, 0x42, 0x54, 0x09, 0x11, 0x32, 0x85, 0x25,
	0x95, 0xda, 0x09, 0xc1, 0x52, 0x89, 0xa5, 0x03, 0x4b, 0xe8, 0x1e, 0x39, 0xc9, 0xb5, 0xb5, 0x84,
	0xed, 0x28, 0x3e, 0x90, 0x2a, 0xb1, 0x30, 0xf2, 0x5f, 0xa3, 0xc6, 0x0e, 0x20, 0x18, 0x98, 0xba,
	0x3e, 0x7d, 0xe7, 0x7b, 0xf7, 0x19, 0x86, 0x95, 0x51, 0xca, 0xe8, 0xac, 0x69, 0x0d, 0x19, 0x3e,
	0xb0, 0xd5, 0x06, 0x95, 0x18, 0x5f, 0xad, 0x8d, 0x59, 0x3f, 0xe3, 0xa4, 0x4b, 0xcb, 0x97, 0xd5,
	0x84, 0xa4, 0x42, 0x4b, 0x42, 0x35, 0x0e, 0x4c, 0x4a, 0xe0, 0x4f, 0xd8, 0xbe, 0xca, 0x0a, 0x1f,
	0xd0, 0x56, 0xad, 0x6c, 0x48, 0x1a, 0xcd, 0xaf, 0x61, 0x68, 0x5d, 0x5a, 0xd0, 0xb6, 0xc1, 0x88,
	0xc5, 0x2c, 0x0d, 0xf2, 0xd0, 0x67, 0xcb, 0x6d, 0x83, 0xfc, 0x06, 0xce, 0x7b, 0x44, 0x6a, 0x4b,
	0x42, 0x57, 0x18, 0x1d, 0x74, 0xd8, 0xc8, 0xe7, 0x0b, 0x1f, 0x27, 0xef, 0x6c, 0xb7, 0xc4, 0x5a,
	0x69, 0xf4, 0xde, 0x96, 0xf0, 0x4b, 0x00, 0xeb, 0x76, 0x14, 0xb2, 0x8e, 0x0e, 0x63, 0x96, 0x9e,
	0xe5, 0x81, 0x4f, 0x16, 0x75, 0xf2, 0xc1, 0x60, 0xe4, 0x3b, 0x3c, 0x22, 0x89, 0x5a, 0x90, 0xe0,
	0xb7, 0x00, 0x96, 0x44, 0x4b, 0x58, 0x17, 0x82, 0xa2, 0xa3, 0x98, 0xa5, 0xe1, 0x74, 0x9c, 0x39,
	0x63, 0x59, 0x6f, 0x2c, 0x5b, 0xf6, 0xc6, 0xf2, 0xc0, 0xd3, 0x73, 0xe2, 0x77, 0x10, 0xae, 0xa4,
	0x96, 0x76, 0xe3, 0x66, 0x8f, 0xff, 0x9d, 0x85, 0x1e, 0x9f, 0x53, 0xf2, 0x06, 0x27, 0xbe, 0x0a,
	0xbf, 0x87, 0xb0, 0xfe, 0x56, 0xd2, 0x29, 0xd8, 0xbd, 0xe3, 0x7e, 0x2f, 0xfb, 0x2b, 0x2d, 0xff,
	0x89, 0xf3, 0x19, 0x9c, 0x2a, 0x7f, 0x4c, 0xa7, 0x25, 0x9c, 0x5e, 0xfc, 0x1a, 0xed, 0x6f, 0xcd,
	0xbf, 0xc0, 0x72, 0xd0, 0xb5, 0x9b, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xf2, 0x17, 0x66,
	0x31, 0x02, 0x00, 0x00,
}
