// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/messenger.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Void models an empty type, with no data associated.
type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messenger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_api_messenger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_api_messenger_proto_rawDescGZIP(), []int{0}
}

// UUID is a wrapper over a universally unique identifier, represented as
// a string.
type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messenger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_api_messenger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_api_messenger_proto_rawDescGZIP(), []int{1}
}

func (x *UUID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// MessageBody models a message body, with the sender's username and the
// message body.
type MessageBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUsername string `protobuf:"bytes,1,opt,name=senderUsername,proto3" json:"senderUsername,omitempty"`
	Body           string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MessageBody) Reset() {
	*x = MessageBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messenger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBody) ProtoMessage() {}

func (x *MessageBody) ProtoReflect() protoreflect.Message {
	mi := &file_api_messenger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBody.ProtoReflect.Descriptor instead.
func (*MessageBody) Descriptor() ([]byte, []int) {
	return file_api_messenger_proto_rawDescGZIP(), []int{2}
}

func (x *MessageBody) GetSenderUsername() string {
	if x != nil {
		return x.SenderUsername
	}
	return ""
}

func (x *MessageBody) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// Message models a message, with a UUID, body, and timestamp. Messages can
// be uniquely identified by their UUID.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body   *MessageBody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	SentAt *Timestamp   `protobuf:"bytes,3,opt,name=sentAt,proto3" json:"sentAt,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messenger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_messenger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_messenger_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetBody() *MessageBody {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Message) GetSentAt() *Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

// Timestamp models a timestamp, with the seconds and nanoseconds associated
// to an instance in time.
type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messenger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_api_messenger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_api_messenger_proto_rawDescGZIP(), []int{4}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_api_messenger_proto protoreflect.FileDescriptor

var file_api_messenger_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x16, 0x0a,
	0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x5f, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x22, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41,
	0x74, 0x22, 0x3b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x32, 0x7a,
	0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x05, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x05, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x08, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x67, 0x72,
	0x70, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_messenger_proto_rawDescOnce sync.Once
	file_api_messenger_proto_rawDescData = file_api_messenger_proto_rawDesc
)

func file_api_messenger_proto_rawDescGZIP() []byte {
	file_api_messenger_proto_rawDescOnce.Do(func() {
		file_api_messenger_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_messenger_proto_rawDescData)
	})
	return file_api_messenger_proto_rawDescData
}

var file_api_messenger_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_messenger_proto_goTypes = []interface{}{
	(*Void)(nil),        // 0: Void
	(*UUID)(nil),        // 1: UUID
	(*MessageBody)(nil), // 2: MessageBody
	(*Message)(nil),     // 3: Message
	(*Timestamp)(nil),   // 4: Timestamp
}
var file_api_messenger_proto_depIdxs = []int32{
	2, // 0: Message.body:type_name -> MessageBody
	4, // 1: Message.sentAt:type_name -> Timestamp
	2, // 2: Messenger.PublishMessage:input_type -> MessageBody
	1, // 3: Messenger.DeleteMessage:input_type -> UUID
	0, // 4: Messenger.ListMessages:input_type -> Void
	1, // 5: Messenger.PublishMessage:output_type -> UUID
	0, // 6: Messenger.DeleteMessage:output_type -> Void
	3, // 7: Messenger.ListMessages:output_type -> Message
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_messenger_proto_init() }
func file_api_messenger_proto_init() {
	if File_api_messenger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_messenger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_api_messenger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_api_messenger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageBody); i {
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
		file_api_messenger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_api_messenger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
			RawDescriptor: file_api_messenger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_messenger_proto_goTypes,
		DependencyIndexes: file_api_messenger_proto_depIdxs,
		MessageInfos:      file_api_messenger_proto_msgTypes,
	}.Build()
	File_api_messenger_proto = out.File
	file_api_messenger_proto_rawDesc = nil
	file_api_messenger_proto_goTypes = nil
	file_api_messenger_proto_depIdxs = nil
}
