// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audit/v1/audit.proto

package auditv1

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	v1 "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

type TimeRange struct {
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_16d867259c094f3a, []int{0}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type GetEventsRequest struct {
	// Types that are valid to be assigned to Window:
	//	*GetEventsRequest_Range
	//	*GetEventsRequest_Since
	Window               isGetEventsRequest_Window `protobuf_oneof:"window"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetEventsRequest) Reset()         { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()    {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_16d867259c094f3a, []int{1}
}

func (m *GetEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsRequest.Unmarshal(m, b)
}
func (m *GetEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsRequest.Marshal(b, m, deterministic)
}
func (m *GetEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsRequest.Merge(m, src)
}
func (m *GetEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventsRequest.Size(m)
}
func (m *GetEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsRequest proto.InternalMessageInfo

type isGetEventsRequest_Window interface {
	isGetEventsRequest_Window()
}

type GetEventsRequest_Range struct {
	Range *TimeRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type GetEventsRequest_Since struct {
	Since *duration.Duration `protobuf:"bytes,2,opt,name=since,proto3,oneof"`
}

func (*GetEventsRequest_Range) isGetEventsRequest_Window() {}

func (*GetEventsRequest_Since) isGetEventsRequest_Window() {}

func (m *GetEventsRequest) GetWindow() isGetEventsRequest_Window {
	if m != nil {
		return m.Window
	}
	return nil
}

func (m *GetEventsRequest) GetRange() *TimeRange {
	if x, ok := m.GetWindow().(*GetEventsRequest_Range); ok {
		return x.Range
	}
	return nil
}

func (m *GetEventsRequest) GetSince() *duration.Duration {
	if x, ok := m.GetWindow().(*GetEventsRequest_Since); ok {
		return x.Since
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetEventsRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetEventsRequest_Range)(nil),
		(*GetEventsRequest_Since)(nil),
	}
}

type Resource struct {
	TypeUrl              string   `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_16d867259c094f3a, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RequestEvent struct {
	// What attempted the action.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// The service performing the operation.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The method being called.
	MethodName string `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// The type of operation being performed.
	Type v1.ActionType `protobuf:"varint,4,opt,name=type,proto3,enum=clutch.api.v1.ActionType" json:"type,omitempty"`
	// The status of the overall operation.
	Status *status.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// The resources touched during the event.
	Resources            []*Resource `protobuf:"bytes,6,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RequestEvent) Reset()         { *m = RequestEvent{} }
func (m *RequestEvent) String() string { return proto.CompactTextString(m) }
func (*RequestEvent) ProtoMessage()    {}
func (*RequestEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_16d867259c094f3a, []int{3}
}

func (m *RequestEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestEvent.Unmarshal(m, b)
}
func (m *RequestEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestEvent.Marshal(b, m, deterministic)
}
func (m *RequestEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEvent.Merge(m, src)
}
func (m *RequestEvent) XXX_Size() int {
	return xxx_messageInfo_RequestEvent.Size(m)
}
func (m *RequestEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEvent proto.InternalMessageInfo

func (m *RequestEvent) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RequestEvent) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RequestEvent) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *RequestEvent) GetType() v1.ActionType {
	if m != nil {
		return m.Type
	}
	return v1.ActionType_UNSPECIFIED
}

func (m *RequestEvent) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RequestEvent) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type Event struct {
	// When the event happened.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,1,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// Types that are valid to be assigned to EventType:
	//	*Event_Event
	EventType            isEvent_EventType `protobuf_oneof:"event_type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_16d867259c094f3a, []int{4}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

type isEvent_EventType interface {
	isEvent_EventType()
}

type Event_Event struct {
	Event *RequestEvent `protobuf:"bytes,2,opt,name=event,proto3,oneof"`
}

func (*Event_Event) isEvent_EventType() {}

func (m *Event) GetEventType() isEvent_EventType {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (m *Event) GetEvent() *RequestEvent {
	if x, ok := m.GetEventType().(*Event_Event); ok {
		return x.Event
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_Event)(nil),
	}
}

type GetEventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsResponse) Reset()         { *m = GetEventsResponse{} }
func (m *GetEventsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventsResponse) ProtoMessage()    {}
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16d867259c094f3a, []int{5}
}

func (m *GetEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsResponse.Unmarshal(m, b)
}
func (m *GetEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsResponse.Marshal(b, m, deterministic)
}
func (m *GetEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsResponse.Merge(m, src)
}
func (m *GetEventsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventsResponse.Size(m)
}
func (m *GetEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsResponse proto.InternalMessageInfo

func (m *GetEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeRange)(nil), "clutch.audit.v1.TimeRange")
	proto.RegisterType((*GetEventsRequest)(nil), "clutch.audit.v1.GetEventsRequest")
	proto.RegisterType((*Resource)(nil), "clutch.audit.v1.Resource")
	proto.RegisterType((*RequestEvent)(nil), "clutch.audit.v1.RequestEvent")
	proto.RegisterType((*Event)(nil), "clutch.audit.v1.Event")
	proto.RegisterType((*GetEventsResponse)(nil), "clutch.audit.v1.GetEventsResponse")
}

func init() {
	proto.RegisterFile("audit/v1/audit.proto", fileDescriptor_16d867259c094f3a)
}

var fileDescriptor_16d867259c094f3a = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0xff, 0xc9, 0xd6, 0x36, 0x79, 0x3b, 0x6d, 0xfb, 0x7b, 0x88, 0x65, 0x55, 0x81, 0x11,
	0x10, 0xaa, 0x26, 0x94, 0xa8, 0x45, 0xd3, 0xa4, 0x71, 0x6a, 0x06, 0x02, 0x2e, 0x08, 0x99, 0x71,
	0xe1, 0x52, 0x79, 0x89, 0xe9, 0x22, 0xb5, 0x49, 0xb0, 0x9d, 0x8e, 0x09, 0x71, 0x19, 0x48, 0x88,
	0x33, 0xdf, 0x62, 0xc7, 0x9d, 0xf8, 0x1e, 0x7c, 0x82, 0x4a, 0x7c, 0x8a, 0x9d, 0x90, 0x1d, 0x27,
	0xdb, 0x5a, 0x60, 0xb7, 0xd8, 0xcf, 0xef, 0x7d, 0xfd, 0xf8, 0xf1, 0xdb, 0xc2, 0x0d, 0x92, 0x47,
	0xb1, 0xf0, 0x27, 0x5d, 0x5f, 0x7d, 0x78, 0x19, 0x4b, 0x45, 0x8a, 0x56, 0xc2, 0x51, 0x2e, 0xc2,
	0x43, 0xaf, 0xd8, 0x9b, 0x74, 0x5b, 0xed, 0x61, 0x9a, 0x0e, 0x47, 0xd4, 0x27, 0x59, 0xec, 0x93,
	0x24, 0x49, 0x05, 0x11, 0x71, 0x9a, 0xf0, 0x02, 0x6f, 0xdd, 0xd6, 0xaa, 0x5a, 0x1d, 0xe4, 0xef,
	0xfc, 0x28, 0x67, 0x0a, 0xd0, 0xfa, 0x9d, 0x59, 0x5d, 0xc4, 0x63, 0xca, 0x05, 0x19, 0x67, 0x1a,
	0x58, 0xd7, 0x00, 0xcb, 0x42, 0x9f, 0x0b, 0x22, 0xf2, 0xb2, 0xf3, 0xfa, 0x84, 0x8c, 0xe2, 0x88,
	0x08, 0xea, 0x97, 0x1f, 0x5a, 0x70, 0xa4, 0x13, 0xe9, 0x7a, 0xce, 0xcc, 0x9a, 0x56, 0x78, 0x78,
	0x48, 0xc7, 0xa4, 0xd8, 0x74, 0xbf, 0x1a, 0x60, 0xef, 0xc7, 0x63, 0x8a, 0x49, 0x32, 0xa4, 0x68,
	0x0f, 0x80, 0x0b, 0xc2, 0xc4, 0x40, 0xfa, 0x70, 0x8c, 0x4d, 0xa3, 0xd3, 0xec, 0xb5, 0xbc, 0xc2,
	0x83, 0x57, 0x9a, 0xf4, 0xf6, 0x4b, 0x93, 0x81, 0x75, 0x1e, 0xd4, 0xce, 0x0c, 0xd3, 0x32, 0xb0,
	0xad, 0xea, 0xa4, 0x82, 0xb6, 0xc1, 0xa2, 0x49, 0x54, 0xb4, 0x30, 0xaf, 0x6b, 0x81, 0x1b, 0x34,
	0x89, 0xe4, 0xca, 0xfd, 0x6c, 0xc0, 0xea, 0x33, 0x2a, 0x9e, 0x4e, 0x68, 0x22, 0x38, 0xa6, 0xef,
	0x73, 0xca, 0x05, 0xea, 0x41, 0x8d, 0x49, 0x67, 0x95, 0x97, 0x99, 0xfc, 0xbd, 0xca, 0xfb, 0xf3,
	0xff, 0x70, 0x81, 0xa2, 0x2e, 0xd4, 0x78, 0x9c, 0x84, 0xe5, 0xe1, 0x1b, 0x73, 0x87, 0x3f, 0xd1,
	0x8f, 0x20, 0x4b, 0x14, 0x19, 0x58, 0x50, 0x3f, 0x8a, 0x93, 0x28, 0x3d, 0x72, 0x43, 0xb0, 0x30,
	0xe5, 0x69, 0xce, 0x42, 0x8a, 0x36, 0xc0, 0x12, 0xc7, 0x19, 0x1d, 0xe4, 0x6c, 0xa4, 0xce, 0xb7,
	0x71, 0x43, 0xae, 0xdf, 0xb0, 0x11, 0x5a, 0x06, 0x33, 0x8e, 0xd4, 0x01, 0x36, 0x36, 0xe3, 0x68,
	0xf7, 0xe1, 0xd9, 0xb4, 0xdd, 0x81, 0x07, 0xe0, 0xcc, 0xfa, 0xab, 0x9a, 0xc1, 0xc7, 0xb2, 0xd7,
	0x27, 0xf7, 0x87, 0x09, 0x4b, 0xfa, 0x86, 0xea, 0xba, 0xe8, 0x1e, 0x58, 0x39, 0xa7, 0x2c, 0x21,
	0x3a, 0x75, 0x3b, 0x68, 0x9c, 0x07, 0x8b, 0xcc, 0x5c, 0x35, 0x70, 0x25, 0xa0, 0x2d, 0x58, 0xe2,
	0x94, 0x4d, 0xe2, 0x90, 0x0e, 0x14, 0x68, 0x5e, 0x05, 0x9b, 0x5a, 0x7c, 0x29, 0xd9, 0x0e, 0x34,
	0xc7, 0x54, 0x1c, 0xa6, 0x51, 0x81, 0x2e, 0x5c, 0x45, 0xa1, 0xd0, 0x14, 0xb9, 0x03, 0x8b, 0xd2,
	0x98, 0xb3, 0xb8, 0x69, 0x74, 0x96, 0x7b, 0x1b, 0x55, 0xc0, 0x59, 0x2c, 0xed, 0xf7, 0x43, 0x19,
	0xd4, 0xfe, 0x71, 0x46, 0xd5, 0x5b, 0x9f, 0x18, 0xb2, 0x5c, 0x15, 0xa0, 0x2d, 0xa8, 0x17, 0x13,
	0xe9, 0xd4, 0x54, 0xce, 0xa8, 0xcc, 0x99, 0x65, 0xa1, 0xf7, 0x5a, 0x29, 0x58, 0x13, 0x68, 0x07,
	0x6c, 0xa6, 0x83, 0xe0, 0x4e, 0x7d, 0x73, 0x41, 0x3d, 0xcb, 0xdf, 0xa2, 0xc2, 0x17, 0xec, 0xee,
	0xca, 0xe9, 0xb4, 0xdd, 0xbc, 0x54, 0xec, 0x7e, 0x33, 0xa0, 0x56, 0x64, 0xf6, 0x18, 0x9a, 0x69,
	0x18, 0xe6, 0x8c, 0xd1, 0x68, 0x40, 0xc4, 0xf5, 0xc3, 0x8a, 0xa1, 0xc4, 0xfb, 0x02, 0x6d, 0x43,
	0x8d, 0xca, 0x2e, 0x7a, 0x46, 0x6e, 0xfd, 0xc1, 0xcc, 0xc5, 0xf3, 0xc8, 0x39, 0x51, 0x74, 0xb0,
	0x04, 0xa0, 0x3e, 0x06, 0x32, 0x01, 0x77, 0x0f, 0xfe, 0xbf, 0x34, 0xb0, 0x3c, 0x4b, 0x13, 0x4e,
	0x91, 0x07, 0x75, 0x85, 0x70, 0xc7, 0x50, 0xf7, 0xbc, 0x39, 0xd7, 0x5a, 0x15, 0x60, 0x4d, 0xf5,
	0xbe, 0x18, 0x60, 0xf5, 0xa5, 0xd4, 0x7f, 0xf5, 0x02, 0x7d, 0x00, 0xbb, 0xea, 0x88, 0xee, 0xce,
	0x55, 0xce, 0xfe, 0x3c, 0x5a, 0xee, 0xbf, 0x90, 0xc2, 0x90, 0x7b, 0xff, 0x74, 0xda, 0x36, 0x2d,
	0xf3, 0xe4, 0xe7, 0xaf, 0xef, 0xa6, 0xe3, 0xae, 0x55, 0xff, 0x68, 0xfe, 0xb0, 0x44, 0x77, 0x8d,
	0xad, 0xc0, 0x7e, 0xdb, 0x50, 0xbb, 0x93, 0xee, 0x41, 0x5d, 0x65, 0xf7, 0xe8, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xba, 0x57, 0x83, 0x75, 0x02, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuditAPIClient is the client API for AuditAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuditAPIClient interface {
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
}

type auditAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditAPIClient(cc grpc.ClientConnInterface) AuditAPIClient {
	return &auditAPIClient{cc}
}

func (c *auditAPIClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/clutch.audit.v1.AuditAPI/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditAPIServer is the server API for AuditAPI service.
type AuditAPIServer interface {
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
}

// UnimplementedAuditAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAuditAPIServer struct {
}

func (*UnimplementedAuditAPIServer) GetEvents(ctx context.Context, req *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}

func RegisterAuditAPIServer(s *grpc.Server, srv AuditAPIServer) {
	s.RegisterService(&_AuditAPI_serviceDesc, srv)
}

func _AuditAPI_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditAPIServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.audit.v1.AuditAPI/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditAPIServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuditAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.audit.v1.AuditAPI",
	HandlerType: (*AuditAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _AuditAPI_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "audit/v1/audit.proto",
}