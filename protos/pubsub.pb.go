// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protos/pubsub.proto

package pubsub

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubscribeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubscribeInput) Reset() {
	*x = SubscribeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeInput) ProtoMessage() {}

func (x *SubscribeInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeInput.ProtoReflect.Descriptor instead.
func (*SubscribeInput) Descriptor() ([]byte, []int) {
	return file_protos_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PushInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"` //* Your payload
}

func (x *PushInput) Reset() {
	*x = PushInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pubsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushInput) ProtoMessage() {}

func (x *PushInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pubsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushInput.ProtoReflect.Descriptor instead.
func (*PushInput) Descriptor() ([]byte, []int) {
	return file_protos_pubsub_proto_rawDescGZIP(), []int{1}
}

func (x *PushInput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` //* Your payload
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_pubsub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protos_pubsub_proto_msgTypes[2]
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
	return file_protos_pubsub_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protos_pubsub_proto protoreflect.FileDescriptor

var file_protos_pubsub_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x61, 0x77, 0x6e, 0x79, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x09, 0x50,
	0x75, 0x73, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x7e, 0x0a, 0x0d, 0x50,
	0x75, 0x62, 0x53, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x15, 0x2e, 0x74, 0x61, 0x77, 0x6e,
	0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x0e, 0x2e, 0x74, 0x61, 0x77, 0x6e, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12,
	0x10, 0x2e, 0x74, 0x61, 0x77, 0x6e, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x3b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_pubsub_proto_rawDescOnce sync.Once
	file_protos_pubsub_proto_rawDescData = file_protos_pubsub_proto_rawDesc
)

func file_protos_pubsub_proto_rawDescGZIP() []byte {
	file_protos_pubsub_proto_rawDescOnce.Do(func() {
		file_protos_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_pubsub_proto_rawDescData)
	})
	return file_protos_pubsub_proto_rawDescData
}

var file_protos_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_pubsub_proto_goTypes = []interface{}{
	(*SubscribeInput)(nil), // 0: tawny.SubscribeInput
	(*PushInput)(nil),      // 1: tawny.PushInput
	(*Message)(nil),        // 2: tawny.Message
	(*emptypb.Empty)(nil),  // 3: google.protobuf.Empty
}
var file_protos_pubsub_proto_depIdxs = []int32{
	0, // 0: tawny.PubSubService.Subscribe:input_type -> tawny.SubscribeInput
	1, // 1: tawny.PubSubService.Publish:input_type -> tawny.PushInput
	2, // 2: tawny.PubSubService.Subscribe:output_type -> tawny.Message
	3, // 3: tawny.PubSubService.Publish:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_pubsub_proto_init() }
func file_protos_pubsub_proto_init() {
	if File_protos_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeInput); i {
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
		file_protos_pubsub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushInput); i {
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
		file_protos_pubsub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_pubsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_pubsub_proto_goTypes,
		DependencyIndexes: file_protos_pubsub_proto_depIdxs,
		MessageInfos:      file_protos_pubsub_proto_msgTypes,
	}.Build()
	File_protos_pubsub_proto = out.File
	file_protos_pubsub_proto_rawDesc = nil
	file_protos_pubsub_proto_goTypes = nil
	file_protos_pubsub_proto_depIdxs = nil
}
