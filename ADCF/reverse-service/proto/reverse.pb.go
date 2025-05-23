// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: proto/reverse.proto

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

type ReverseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *ReverseRequest) Reset() {
	*x = ReverseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reverse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseRequest) ProtoMessage() {}

func (x *ReverseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reverse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseRequest.ProtoReflect.Descriptor instead.
func (*ReverseRequest) Descriptor() ([]byte, []int) {
	return file_proto_reverse_proto_rawDescGZIP(), []int{0}
}

func (x *ReverseRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type ReverseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ReverseResponse) Reset() {
	*x = ReverseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reverse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseResponse) ProtoMessage() {}

func (x *ReverseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reverse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseResponse.ProtoReflect.Descriptor instead.
func (*ReverseResponse) Descriptor() ([]byte, []int) {
	return file_proto_reverse_proto_rawDescGZIP(), []int{1}
}

func (x *ReverseResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_proto_reverse_proto protoreflect.FileDescriptor

var file_proto_reverse_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x22, 0x26,
	0x0a, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x32, 0x50, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x17,
	0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reverse_proto_rawDescOnce sync.Once
	file_proto_reverse_proto_rawDescData = file_proto_reverse_proto_rawDesc
)

func file_proto_reverse_proto_rawDescGZIP() []byte {
	file_proto_reverse_proto_rawDescOnce.Do(func() {
		file_proto_reverse_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reverse_proto_rawDescData)
	})
	return file_proto_reverse_proto_rawDescData
}

var file_proto_reverse_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_reverse_proto_goTypes = []interface{}{
	(*ReverseRequest)(nil),  // 0: reverse.ReverseRequest
	(*ReverseResponse)(nil), // 1: reverse.ReverseResponse
}
var file_proto_reverse_proto_depIdxs = []int32{
	0, // 0: reverse.ReverseService.Reverse:input_type -> reverse.ReverseRequest
	1, // 1: reverse.ReverseService.Reverse:output_type -> reverse.ReverseResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_reverse_proto_init() }
func file_proto_reverse_proto_init() {
	if File_proto_reverse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_reverse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseRequest); i {
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
		file_proto_reverse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseResponse); i {
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
			RawDescriptor: file_proto_reverse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reverse_proto_goTypes,
		DependencyIndexes: file_proto_reverse_proto_depIdxs,
		MessageInfos:      file_proto_reverse_proto_msgTypes,
	}.Build()
	File_proto_reverse_proto = out.File
	file_proto_reverse_proto_rawDesc = nil
	file_proto_reverse_proto_goTypes = nil
	file_proto_reverse_proto_depIdxs = nil
}
