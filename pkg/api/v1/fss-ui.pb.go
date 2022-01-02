// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: fss-ui.proto

package v1

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

type ListExtinguisherSpecsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListExtinguisherSpecsRequest) Reset() {
	*x = ListExtinguisherSpecsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fss_ui_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExtinguisherSpecsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExtinguisherSpecsRequest) ProtoMessage() {}

func (x *ListExtinguisherSpecsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fss_ui_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExtinguisherSpecsRequest.ProtoReflect.Descriptor instead.
func (*ListExtinguisherSpecsRequest) Descriptor() ([]byte, []int) {
	return file_fss_ui_proto_rawDescGZIP(), []int{0}
}

type ListExtinguisherSpecsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo        *Repository          `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path        string               `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Description string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Arguments   []*DesiredAnnotation `protobuf:"bytes,5,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *ListExtinguisherSpecsResponse) Reset() {
	*x = ListExtinguisherSpecsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fss_ui_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExtinguisherSpecsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExtinguisherSpecsResponse) ProtoMessage() {}

func (x *ListExtinguisherSpecsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fss_ui_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExtinguisherSpecsResponse.ProtoReflect.Descriptor instead.
func (*ListExtinguisherSpecsResponse) Descriptor() ([]byte, []int) {
	return file_fss_ui_proto_rawDescGZIP(), []int{1}
}

func (x *ListExtinguisherSpecsResponse) GetRepo() *Repository {
	if x != nil {
		return x.Repo
	}
	return nil
}

func (x *ListExtinguisherSpecsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListExtinguisherSpecsResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListExtinguisherSpecsResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListExtinguisherSpecsResponse) GetArguments() []*DesiredAnnotation {
	if x != nil {
		return x.Arguments
	}
	return nil
}

// DesiredAnnotation describes an annotation a Extinguisher should have
type DesiredAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required    bool   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DesiredAnnotation) Reset() {
	*x = DesiredAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fss_ui_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DesiredAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DesiredAnnotation) ProtoMessage() {}

func (x *DesiredAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_fss_ui_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DesiredAnnotation.ProtoReflect.Descriptor instead.
func (*DesiredAnnotation) Descriptor() ([]byte, []int) {
	return file_fss_ui_proto_rawDescGZIP(), []int{2}
}

func (x *DesiredAnnotation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DesiredAnnotation) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *DesiredAnnotation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type IsReadOnlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IsReadOnlyRequest) Reset() {
	*x = IsReadOnlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fss_ui_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsReadOnlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsReadOnlyRequest) ProtoMessage() {}

func (x *IsReadOnlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fss_ui_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsReadOnlyRequest.ProtoReflect.Descriptor instead.
func (*IsReadOnlyRequest) Descriptor() ([]byte, []int) {
	return file_fss_ui_proto_rawDescGZIP(), []int{3}
}

type IsReadOnlyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Readonly bool `protobuf:"varint,1,opt,name=readonly,proto3" json:"readonly,omitempty"`
}

func (x *IsReadOnlyResponse) Reset() {
	*x = IsReadOnlyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fss_ui_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsReadOnlyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsReadOnlyResponse) ProtoMessage() {}

func (x *IsReadOnlyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fss_ui_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsReadOnlyResponse.ProtoReflect.Descriptor instead.
func (*IsReadOnlyResponse) Descriptor() ([]byte, []int) {
	return file_fss_ui_proto_rawDescGZIP(), []int{4}
}

func (x *IsReadOnlyResponse) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

var File_fss_ui_proto protoreflect.FileDescriptor

var file_fss_ui_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x73, 0x73, 0x2d, 0x75, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x09, 0x66, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a,
	0x1c, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x69, 0x6e, 0x67, 0x75, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc2, 0x01,
	0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x69, 0x6e, 0x67, 0x75, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x72,
	0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x65, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x49, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30,
	0x0a, 0x12, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79,
	0x32, 0xa8, 0x01, 0x0a, 0x05, 0x46, 0x73, 0x73, 0x55, 0x49, 0x12, 0x60, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x78, 0x74, 0x69, 0x6e, 0x67, 0x75, 0x69, 0x73, 0x68, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x63, 0x73, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74,
	0x69, 0x6e, 0x67, 0x75, 0x69, 0x73, 0x68, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x78, 0x74, 0x69, 0x6e, 0x67, 0x75, 0x69, 0x73, 0x68, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x0a,
	0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x68, 0x6f, 0x6a, 0x70, 0x75,
	0x72, 0x2f, 0x66, 0x73, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fss_ui_proto_rawDescOnce sync.Once
	file_fss_ui_proto_rawDescData = file_fss_ui_proto_rawDesc
)

func file_fss_ui_proto_rawDescGZIP() []byte {
	file_fss_ui_proto_rawDescOnce.Do(func() {
		file_fss_ui_proto_rawDescData = protoimpl.X.CompressGZIP(file_fss_ui_proto_rawDescData)
	})
	return file_fss_ui_proto_rawDescData
}

var file_fss_ui_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_fss_ui_proto_goTypes = []interface{}{
	(*ListExtinguisherSpecsRequest)(nil),  // 0: v1.ListExtinguisherSpecsRequest
	(*ListExtinguisherSpecsResponse)(nil), // 1: v1.ListExtinguisherSpecsResponse
	(*DesiredAnnotation)(nil),             // 2: v1.DesiredAnnotation
	(*IsReadOnlyRequest)(nil),             // 3: v1.IsReadOnlyRequest
	(*IsReadOnlyResponse)(nil),            // 4: v1.IsReadOnlyResponse
	(*Repository)(nil),                    // 5: v1.Repository
}
var file_fss_ui_proto_depIdxs = []int32{
	5, // 0: v1.ListExtinguisherSpecsResponse.repo:type_name -> v1.Repository
	2, // 1: v1.ListExtinguisherSpecsResponse.arguments:type_name -> v1.DesiredAnnotation
	0, // 2: v1.FssUI.ListExtinguisherSpecs:input_type -> v1.ListExtinguisherSpecsRequest
	3, // 3: v1.FssUI.IsReadOnly:input_type -> v1.IsReadOnlyRequest
	1, // 4: v1.FssUI.ListExtinguisherSpecs:output_type -> v1.ListExtinguisherSpecsResponse
	4, // 5: v1.FssUI.IsReadOnly:output_type -> v1.IsReadOnlyResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fss_ui_proto_init() }
func file_fss_ui_proto_init() {
	if File_fss_ui_proto != nil {
		return
	}
	file_fss_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fss_ui_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExtinguisherSpecsRequest); i {
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
		file_fss_ui_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExtinguisherSpecsResponse); i {
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
		file_fss_ui_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DesiredAnnotation); i {
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
		file_fss_ui_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsReadOnlyRequest); i {
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
		file_fss_ui_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsReadOnlyResponse); i {
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
			RawDescriptor: file_fss_ui_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fss_ui_proto_goTypes,
		DependencyIndexes: file_fss_ui_proto_depIdxs,
		MessageInfos:      file_fss_ui_proto_msgTypes,
	}.Build()
	File_fss_ui_proto = out.File
	file_fss_ui_proto_rawDesc = nil
	file_fss_ui_proto_goTypes = nil
	file_fss_ui_proto_depIdxs = nil
}
