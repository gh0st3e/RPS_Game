// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: result.proto

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

type RequestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserValue     string `protobuf:"bytes,1,opt,name=userValue,proto3" json:"userValue,omitempty"`
	ComputerValue string `protobuf:"bytes,2,opt,name=computerValue,proto3" json:"computerValue,omitempty"`
}

func (x *RequestResult) Reset() {
	*x = RequestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResult) ProtoMessage() {}

func (x *RequestResult) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResult.ProtoReflect.Descriptor instead.
func (*RequestResult) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

func (x *RequestResult) GetUserValue() string {
	if x != nil {
		return x.UserValue
	}
	return ""
}

func (x *RequestResult) GetComputerValue() string {
	if x != nil {
		return x.ComputerValue
	}
	return ""
}

type ResponseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Winner string `protobuf:"bytes,1,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *ResponseResult) Reset() {
	*x = ResponseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseResult) ProtoMessage() {}

func (x *ResponseResult) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseResult.ProtoReflect.Descriptor instead.
func (*ResponseResult) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseResult) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

var File_result_proto protoreflect.FileDescriptor

var file_result_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x32, 0x45, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b,
	0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_result_proto_rawDescOnce sync.Once
	file_result_proto_rawDescData = file_result_proto_rawDesc
)

func file_result_proto_rawDescGZIP() []byte {
	file_result_proto_rawDescOnce.Do(func() {
		file_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_proto_rawDescData)
	})
	return file_result_proto_rawDescData
}

var file_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_result_proto_goTypes = []interface{}{
	(*RequestResult)(nil),  // 0: proto.RequestResult
	(*ResponseResult)(nil), // 1: proto.ResponseResult
}
var file_result_proto_depIdxs = []int32{
	0, // 0: proto.Result.CallResult:input_type -> proto.RequestResult
	1, // 1: proto.Result.CallResult:output_type -> proto.ResponseResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_result_proto_init() }
func file_result_proto_init() {
	if File_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResult); i {
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
		file_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseResult); i {
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
			RawDescriptor: file_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_result_proto_goTypes,
		DependencyIndexes: file_result_proto_depIdxs,
		MessageInfos:      file_result_proto_msgTypes,
	}.Build()
	File_result_proto = out.File
	file_result_proto_rawDesc = nil
	file_result_proto_goTypes = nil
	file_result_proto_depIdxs = nil
}
