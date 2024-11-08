// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: nvidia.proto

package nvidia

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

type NVIDIARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt string `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Model  string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *NVIDIARequest) Reset() {
	*x = NVIDIARequest{}
	mi := &file_nvidia_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NVIDIARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NVIDIARequest) ProtoMessage() {}

func (x *NVIDIARequest) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NVIDIARequest.ProtoReflect.Descriptor instead.
func (*NVIDIARequest) Descriptor() ([]byte, []int) {
	return file_nvidia_proto_rawDescGZIP(), []int{0}
}

func (x *NVIDIARequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *NVIDIARequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type NVIDIAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *NVIDIAResponse) Reset() {
	*x = NVIDIAResponse{}
	mi := &file_nvidia_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NVIDIAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NVIDIAResponse) ProtoMessage() {}

func (x *NVIDIAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nvidia_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NVIDIAResponse.ProtoReflect.Descriptor instead.
func (*NVIDIAResponse) Descriptor() ([]byte, []int) {
	return file_nvidia_proto_rawDescGZIP(), []int{1}
}

func (x *NVIDIAResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_nvidia_proto protoreflect.FileDescriptor

var file_nvidia_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x22, 0x3d, 0x0a, 0x0d, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x28, 0x0a, 0x0e, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x45, 0x0a, 0x06, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x4e,
	0x56, 0x49, 0x44, 0x49, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e,
	0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x4e, 0x56, 0x49, 0x44, 0x49, 0x41, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6e, 0x76, 0x69, 0x64,
	0x69, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nvidia_proto_rawDescOnce sync.Once
	file_nvidia_proto_rawDescData = file_nvidia_proto_rawDesc
)

func file_nvidia_proto_rawDescGZIP() []byte {
	file_nvidia_proto_rawDescOnce.Do(func() {
		file_nvidia_proto_rawDescData = protoimpl.X.CompressGZIP(file_nvidia_proto_rawDescData)
	})
	return file_nvidia_proto_rawDescData
}

var file_nvidia_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nvidia_proto_goTypes = []any{
	(*NVIDIARequest)(nil),  // 0: nvidia.NVIDIARequest
	(*NVIDIAResponse)(nil), // 1: nvidia.NVIDIAResponse
}
var file_nvidia_proto_depIdxs = []int32{
	0, // 0: nvidia.NVIDIA.Generate:input_type -> nvidia.NVIDIARequest
	1, // 1: nvidia.NVIDIA.Generate:output_type -> nvidia.NVIDIAResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nvidia_proto_init() }
func file_nvidia_proto_init() {
	if File_nvidia_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nvidia_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nvidia_proto_goTypes,
		DependencyIndexes: file_nvidia_proto_depIdxs,
		MessageInfos:      file_nvidia_proto_msgTypes,
	}.Build()
	File_nvidia_proto = out.File
	file_nvidia_proto_rawDesc = nil
	file_nvidia_proto_goTypes = nil
	file_nvidia_proto_depIdxs = nil
}
