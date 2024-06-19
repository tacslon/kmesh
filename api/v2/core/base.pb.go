// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.17.3
// source: api/core/base.proto

package core

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

// The status of the control-plane of the current data is identified to
// determine the corresponding lower control-plane operation.
type ApiStatus int32

const (
	ApiStatus_NONE      ApiStatus = 0
	ApiStatus_DELETE    ApiStatus = 1
	ApiStatus_UPDATE    ApiStatus = 2
	ApiStatus_UNCHANGED ApiStatus = 3
	ApiStatus_ALL       ApiStatus = 4
	// Waiting for associated endpoints before updating bpf map
	// Currently only apply to cluster resource.
	ApiStatus_WAITING ApiStatus = 5
)

// Enum value maps for ApiStatus.
var (
	ApiStatus_name = map[int32]string{
		0: "NONE",
		1: "DELETE",
		2: "UPDATE",
		3: "UNCHANGED",
		4: "ALL",
		5: "WAITING",
	}
	ApiStatus_value = map[string]int32{
		"NONE":      0,
		"DELETE":    1,
		"UPDATE":    2,
		"UNCHANGED": 3,
		"ALL":       4,
		"WAITING":   5,
	}
)

func (x ApiStatus) Enum() *ApiStatus {
	p := new(ApiStatus)
	*p = x
	return p
}

func (x ApiStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_core_base_proto_enumTypes[0].Descriptor()
}

func (ApiStatus) Type() protoreflect.EnumType {
	return &file_api_core_base_proto_enumTypes[0]
}

func (x ApiStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiStatus.Descriptor instead.
func (ApiStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_core_base_proto_rawDescGZIP(), []int{0}
}

type RoutingPriority int32

const (
	RoutingPriority_DEFAULT RoutingPriority = 0
	RoutingPriority_HIGH    RoutingPriority = 1
)

// Enum value maps for RoutingPriority.
var (
	RoutingPriority_name = map[int32]string{
		0: "DEFAULT",
		1: "HIGH",
	}
	RoutingPriority_value = map[string]int32{
		"DEFAULT": 0,
		"HIGH":    1,
	}
)

func (x RoutingPriority) Enum() *RoutingPriority {
	p := new(RoutingPriority)
	*p = x
	return p
}

func (x RoutingPriority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoutingPriority) Descriptor() protoreflect.EnumDescriptor {
	return file_api_core_base_proto_enumTypes[1].Descriptor()
}

func (RoutingPriority) Type() protoreflect.EnumType {
	return &file_api_core_base_proto_enumTypes[1]
}

func (x RoutingPriority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoutingPriority.Descriptor instead.
func (RoutingPriority) EnumDescriptor() ([]byte, []int) {
	return file_api_core_base_proto_rawDescGZIP(), []int{1}
}

var File_api_core_base_proto protoreflect.FileDescriptor

var file_api_core_base_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x2a, 0x52, 0x0a, 0x09, 0x41,
	0x70, 0x69, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x2a,
	0x28, 0x0a, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x42, 0x1f, 0x5a, 0x1d, 0x6b, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_core_base_proto_rawDescOnce sync.Once
	file_api_core_base_proto_rawDescData = file_api_core_base_proto_rawDesc
)

func file_api_core_base_proto_rawDescGZIP() []byte {
	file_api_core_base_proto_rawDescOnce.Do(func() {
		file_api_core_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_core_base_proto_rawDescData)
	})
	return file_api_core_base_proto_rawDescData
}

var file_api_core_base_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_core_base_proto_goTypes = []interface{}{
	(ApiStatus)(0),       // 0: core.ApiStatus
	(RoutingPriority)(0), // 1: core.RoutingPriority
}
var file_api_core_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_core_base_proto_init() }
func file_api_core_base_proto_init() {
	if File_api_core_base_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_core_base_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_core_base_proto_goTypes,
		DependencyIndexes: file_api_core_base_proto_depIdxs,
		EnumInfos:         file_api_core_base_proto_enumTypes,
	}.Build()
	File_api_core_base_proto = out.File
	file_api_core_base_proto_rawDesc = nil
	file_api_core_base_proto_goTypes = nil
	file_api_core_base_proto_depIdxs = nil
}
