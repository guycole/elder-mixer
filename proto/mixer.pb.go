// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: mixer.proto

package proto

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

type EnqueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // requestor
	Command  string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`                   // command
}

func (x *EnqueueRequest) Reset() {
	*x = EnqueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mixer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueRequest) ProtoMessage() {}

func (x *EnqueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mixer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueRequest.ProtoReflect.Descriptor instead.
func (*EnqueueRequest) Descriptor() ([]byte, []int) {
	return file_mixer_proto_rawDescGZIP(), []int{0}
}

func (x *EnqueueRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *EnqueueRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type EnqueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`    // requestor
	ReceiptId string `protobuf:"bytes,2,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"` // response receipt
}

func (x *EnqueueResponse) Reset() {
	*x = EnqueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mixer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnqueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueResponse) ProtoMessage() {}

func (x *EnqueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mixer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueResponse.ProtoReflect.Descriptor instead.
func (*EnqueueResponse) Descriptor() ([]byte, []int) {
	return file_mixer_proto_rawDescGZIP(), []int{1}
}

func (x *EnqueueResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *EnqueueResponse) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

var File_mixer_proto protoreflect.FileDescriptor

var file_mixer_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x4d, 0x0a,
	0x0f, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x32, 0x4a, 0x0a, 0x05,
	0x4d, 0x69, 0x78, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75, 0x79, 0x63, 0x6f, 0x6c, 0x65, 0x2f, 0x65,
	0x6c, 0x64, 0x65, 0x72, 0x2d, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mixer_proto_rawDescOnce sync.Once
	file_mixer_proto_rawDescData = file_mixer_proto_rawDesc
)

func file_mixer_proto_rawDescGZIP() []byte {
	file_mixer_proto_rawDescOnce.Do(func() {
		file_mixer_proto_rawDescData = protoimpl.X.CompressGZIP(file_mixer_proto_rawDescData)
	})
	return file_mixer_proto_rawDescData
}

var file_mixer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mixer_proto_goTypes = []interface{}{
	(*EnqueueRequest)(nil),  // 0: proto.EnqueueRequest
	(*EnqueueResponse)(nil), // 1: proto.EnqueueResponse
}
var file_mixer_proto_depIdxs = []int32{
	0, // 0: proto.Mixer.EnqueueCommand:input_type -> proto.EnqueueRequest
	1, // 1: proto.Mixer.EnqueueCommand:output_type -> proto.EnqueueResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mixer_proto_init() }
func file_mixer_proto_init() {
	if File_mixer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mixer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueRequest); i {
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
		file_mixer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnqueueResponse); i {
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
			RawDescriptor: file_mixer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mixer_proto_goTypes,
		DependencyIndexes: file_mixer_proto_depIdxs,
		MessageInfos:      file_mixer_proto_msgTypes,
	}.Build()
	File_mixer_proto = out.File
	file_mixer_proto_rawDesc = nil
	file_mixer_proto_goTypes = nil
	file_mixer_proto_depIdxs = nil
}
