// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: datastruct.proto

package datastruct

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

type ProtoDataStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField string           `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	IntField    int32            `protobuf:"varint,2,opt,name=int_field,json=intField,proto3" json:"int_field,omitempty"`
	FloatField  float64          `protobuf:"fixed64,3,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	ArrayField  []string         `protobuf:"bytes,4,rep,name=array_field,json=arrayField,proto3" json:"array_field,omitempty"`
	MapField    map[string]int32 `protobuf:"bytes,5,rep,name=map_field,json=mapField,proto3" json:"map_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ProtoDataStruct) Reset() {
	*x = ProtoDataStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastruct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoDataStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoDataStruct) ProtoMessage() {}

func (x *ProtoDataStruct) ProtoReflect() protoreflect.Message {
	mi := &file_datastruct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoDataStruct.ProtoReflect.Descriptor instead.
func (*ProtoDataStruct) Descriptor() ([]byte, []int) {
	return file_datastruct_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoDataStruct) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *ProtoDataStruct) GetIntField() int32 {
	if x != nil {
		return x.IntField
	}
	return 0
}

func (x *ProtoDataStruct) GetFloatField() float64 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *ProtoDataStruct) GetArrayField() []string {
	if x != nil {
		return x.ArrayField
	}
	return nil
}

func (x *ProtoDataStruct) GetMapField() map[string]int32 {
	if x != nil {
		return x.MapField
	}
	return nil
}

var File_datastruct_proto protoreflect.FileDescriptor

var file_datastruct_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x92, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x40,
	0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datastruct_proto_rawDescOnce sync.Once
	file_datastruct_proto_rawDescData = file_datastruct_proto_rawDesc
)

func file_datastruct_proto_rawDescGZIP() []byte {
	file_datastruct_proto_rawDescOnce.Do(func() {
		file_datastruct_proto_rawDescData = protoimpl.X.CompressGZIP(file_datastruct_proto_rawDescData)
	})
	return file_datastruct_proto_rawDescData
}

var file_datastruct_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datastruct_proto_goTypes = []interface{}{
	(*ProtoDataStruct)(nil), // 0: main.ProtoDataStruct
	nil,                     // 1: main.ProtoDataStruct.MapFieldEntry
}
var file_datastruct_proto_depIdxs = []int32{
	1, // 0: main.ProtoDataStruct.map_field:type_name -> main.ProtoDataStruct.MapFieldEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datastruct_proto_init() }
func file_datastruct_proto_init() {
	if File_datastruct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datastruct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoDataStruct); i {
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
			RawDescriptor: file_datastruct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datastruct_proto_goTypes,
		DependencyIndexes: file_datastruct_proto_depIdxs,
		MessageInfos:      file_datastruct_proto_msgTypes,
	}.Build()
	File_datastruct_proto = out.File
	file_datastruct_proto_rawDesc = nil
	file_datastruct_proto_goTypes = nil
	file_datastruct_proto_depIdxs = nil
}
