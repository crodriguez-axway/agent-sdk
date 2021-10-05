// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: watch.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event_Type int32

const (
	Event_CREATED Event_Type = 0
	Event_UPDATED Event_Type = 1
	Event_DELETED Event_Type = 2
)

// Enum value maps for Event_Type.
var (
	Event_Type_name = map[int32]string{
		0: "CREATED",
		1: "UPDATED",
		2: "DELETED",
	}
	Event_Type_value = map[string]int32{
		"CREATED": 0,
		"UPDATED": 1,
		"DELETED": 2,
	}
)

func (x Event_Type) Enum() *Event_Type {
	p := new(Event_Type)
	*p = x
	return p
}

func (x Event_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_watch_proto_enumTypes[0].Descriptor()
}

func (Event_Type) Type() protoreflect.EnumType {
	return &file_watch_proto_enumTypes[0]
}

func (x Event_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_Type.Descriptor instead.
func (Event_Type) EnumDescriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{2, 0}
}

type Trigger_Type int32

const (
	Trigger_CREATED Trigger_Type = 0
	Trigger_UPDATED Trigger_Type = 1
	Trigger_DELETED Trigger_Type = 2
)

// Enum value maps for Trigger_Type.
var (
	Trigger_Type_name = map[int32]string{
		0: "CREATED",
		1: "UPDATED",
		2: "DELETED",
	}
	Trigger_Type_value = map[string]int32{
		"CREATED": 0,
		"UPDATED": 1,
		"DELETED": 2,
	}
)

func (x Trigger_Type) Enum() *Trigger_Type {
	p := new(Trigger_Type)
	*p = x
	return p
}

func (x Trigger_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Trigger_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_watch_proto_enumTypes[1].Descriptor()
}

func (Trigger_Type) Type() protoreflect.EnumType {
	return &file_watch_proto_enumTypes[1]
}

func (x Trigger_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Trigger_Type.Descriptor instead.
func (Trigger_Type) EnumDescriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{4, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Triggers []*Trigger `protobuf:"bytes,1,rep,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_watch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetTriggers() []*Trigger {
	if x != nil {
		return x.Triggers
	}
	return nil
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfLink string `protobuf:"bytes,1,opt,name=selfLink,proto3" json:"selfLink,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_watch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{1}
}

func (x *Topic) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time          string            `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Version       string            `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Product       string            `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	CorrelationId string            `protobuf:"bytes,5,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	Organization  *Organization     `protobuf:"bytes,6,opt,name=organization,proto3" json:"organization,omitempty"`
	Type          Event_Type        `protobuf:"varint,7,opt,name=type,proto3,enum=apis.v1.Event_Type" json:"type,omitempty"`
	Payload       *ResourceInstance `protobuf:"bytes,8,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_watch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Event) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Event) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Event) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *Event) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *Event) GetType() Event_Type {
	if x != nil {
		return x.Type
	}
	return Event_CREATED
}

func (x *Event) GetPayload() *ResourceInstance {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_watch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{3}
}

func (x *Organization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string         `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Kind  string         `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Name  string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Scope *Trigger_Scope `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	Type  []Trigger_Type `protobuf:"varint,5,rep,packed,name=type,proto3,enum=apis.v1.Trigger_Type" json:"type,omitempty"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_watch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{4}
}

func (x *Trigger) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Trigger) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Trigger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trigger) GetScope() *Trigger_Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Trigger) GetType() []Trigger_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

type Trigger_Scope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Trigger_Scope) Reset() {
	*x = Trigger_Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger_Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger_Scope) ProtoMessage() {}

func (x *Trigger_Scope) ProtoReflect() protoreflect.Message {
	mi := &file_watch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger_Scope.ProtoReflect.Descriptor instead.
func (*Trigger_Scope) Descriptor() ([]byte, []int) {
	return file_watch_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Trigger_Scope) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Trigger_Scope) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_watch_proto protoreflect.FileDescriptor

var file_watch_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x10, 0x61, 0x70, 0x69, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x22, 0x23, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0xde, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x1e, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x80, 0x02, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x2f, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x7b, 0x0a, 0x0c, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x36, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x1a, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x78, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_watch_proto_rawDescOnce sync.Once
	file_watch_proto_rawDescData = file_watch_proto_rawDesc
)

func file_watch_proto_rawDescGZIP() []byte {
	file_watch_proto_rawDescOnce.Do(func() {
		file_watch_proto_rawDescData = protoimpl.X.CompressGZIP(file_watch_proto_rawDescData)
	})
	return file_watch_proto_rawDescData
}

var file_watch_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_watch_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_watch_proto_goTypes = []interface{}{
	(Event_Type)(0),          // 0: apis.v1.Event.Type
	(Trigger_Type)(0),        // 1: apis.v1.Trigger.Type
	(*Request)(nil),          // 2: apis.v1.Request
	(*Topic)(nil),            // 3: apis.v1.Topic
	(*Event)(nil),            // 4: apis.v1.Event
	(*Organization)(nil),     // 5: apis.v1.Organization
	(*Trigger)(nil),          // 6: apis.v1.Trigger
	(*Trigger_Scope)(nil),    // 7: apis.v1.Trigger.Scope
	(*ResourceInstance)(nil), // 8: amplifycentral.datamodel.ResourceInstance
}
var file_watch_proto_depIdxs = []int32{
	6, // 0: apis.v1.Request.triggers:type_name -> apis.v1.Trigger
	5, // 1: apis.v1.Event.organization:type_name -> apis.v1.Organization
	0, // 2: apis.v1.Event.type:type_name -> apis.v1.Event.Type
	8, // 3: apis.v1.Event.payload:type_name -> amplifycentral.datamodel.ResourceInstance
	7, // 4: apis.v1.Trigger.scope:type_name -> apis.v1.Trigger.Scope
	1, // 5: apis.v1.Trigger.type:type_name -> apis.v1.Trigger.Type
	2, // 6: apis.v1.WatchService.CreateWatch:input_type -> apis.v1.Request
	3, // 7: apis.v1.WatchService.SubscribeToTopic:input_type -> apis.v1.Topic
	4, // 8: apis.v1.WatchService.CreateWatch:output_type -> apis.v1.Event
	4, // 9: apis.v1.WatchService.SubscribeToTopic:output_type -> apis.v1.Event
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_watch_proto_init() }
func file_watch_proto_init() {
	if File_watch_proto != nil {
		return
	}
	file_apicentral_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_watch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_watch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topic); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_watch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_watch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_watch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_watch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger_Scope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_watch_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_watch_proto_goTypes,
		DependencyIndexes: file_watch_proto_depIdxs,
		EnumInfos:         file_watch_proto_enumTypes,
		MessageInfos:      file_watch_proto_msgTypes,
	}.Build()
	File_watch_proto = out.File
	file_watch_proto_rawDesc = nil
	file_watch_proto_goTypes = nil
	file_watch_proto_depIdxs = nil
}