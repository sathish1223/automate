// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: interservice/cfgmgmt/request/parameters.proto

package request

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

type Order int32

const (
	Order_ASC  Order = 0 // By default we order ascending
	Order_DESC Order = 1
)

// Enum value maps for Order.
var (
	Order_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	Order_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x Order) Enum() *Order {
	p := new(Order)
	*p = x
	return p
}

func (x Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order) Descriptor() protoreflect.EnumDescriptor {
	return file_interservice_cfgmgmt_request_parameters_proto_enumTypes[0].Descriptor()
}

func (Order) Type() protoreflect.EnumType {
	return &file_interservice_cfgmgmt_request_parameters_proto_enumTypes[0]
}

func (x Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order.Descriptor instead.
func (Order) EnumDescriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_parameters_proto_rawDescGZIP(), []int{0}
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" toml:"page,omitempty" mapstructure:"page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty" toml:"size,omitempty" mapstructure:"size,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_request_parameters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_parameters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_parameters_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Sorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty" toml:"field,omitempty" mapstructure:"field,omitempty"`
	Order Order  `protobuf:"varint,2,opt,name=order,proto3,enum=chef.automate.domain.cfgmgmt.request.Order" json:"order,omitempty" toml:"order,omitempty" mapstructure:"order,omitempty"`
}

func (x *Sorting) Reset() {
	*x = Sorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_request_parameters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sorting) ProtoMessage() {}

func (x *Sorting) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_parameters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sorting.ProtoReflect.Descriptor instead.
func (*Sorting) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_parameters_proto_rawDescGZIP(), []int{1}
}

func (x *Sorting) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Sorting) GetOrder() Order {
	if x != nil {
		return x.Order
	}
	return Order_ASC
}

type Suggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty" mapstructure:"type,omitempty"`
	Text   string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty" toml:"text,omitempty" mapstructure:"text,omitempty"`
	Filter []string `protobuf:"bytes,3,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
}

func (x *Suggestion) Reset() {
	*x = Suggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_request_parameters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggestion) ProtoMessage() {}

func (x *Suggestion) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_parameters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggestion.ProtoReflect.Descriptor instead.
func (*Suggestion) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_parameters_proto_rawDescGZIP(), []int{2}
}

func (x *Suggestion) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Suggestion) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Suggestion) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

var File_interservice_cfgmgmt_request_parameters_proto protoreflect.FileDescriptor

var file_interservice_cfgmgmt_request_parameters_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x24, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x62, 0x0a, 0x07, 0x53,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x41, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0x4c, 0x0a, 0x0a, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2a, 0x1a, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_cfgmgmt_request_parameters_proto_rawDescOnce sync.Once
	file_interservice_cfgmgmt_request_parameters_proto_rawDescData = file_interservice_cfgmgmt_request_parameters_proto_rawDesc
)

func file_interservice_cfgmgmt_request_parameters_proto_rawDescGZIP() []byte {
	file_interservice_cfgmgmt_request_parameters_proto_rawDescOnce.Do(func() {
		file_interservice_cfgmgmt_request_parameters_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_cfgmgmt_request_parameters_proto_rawDescData)
	})
	return file_interservice_cfgmgmt_request_parameters_proto_rawDescData
}

var file_interservice_cfgmgmt_request_parameters_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_interservice_cfgmgmt_request_parameters_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interservice_cfgmgmt_request_parameters_proto_goTypes = []interface{}{
	(Order)(0),         // 0: chef.automate.domain.cfgmgmt.request.Order
	(*Pagination)(nil), // 1: chef.automate.domain.cfgmgmt.request.Pagination
	(*Sorting)(nil),    // 2: chef.automate.domain.cfgmgmt.request.Sorting
	(*Suggestion)(nil), // 3: chef.automate.domain.cfgmgmt.request.Suggestion
}
var file_interservice_cfgmgmt_request_parameters_proto_depIdxs = []int32{
	0, // 0: chef.automate.domain.cfgmgmt.request.Sorting.order:type_name -> chef.automate.domain.cfgmgmt.request.Order
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interservice_cfgmgmt_request_parameters_proto_init() }
func file_interservice_cfgmgmt_request_parameters_proto_init() {
	if File_interservice_cfgmgmt_request_parameters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_cfgmgmt_request_parameters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_interservice_cfgmgmt_request_parameters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sorting); i {
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
		file_interservice_cfgmgmt_request_parameters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suggestion); i {
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
			RawDescriptor: file_interservice_cfgmgmt_request_parameters_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_cfgmgmt_request_parameters_proto_goTypes,
		DependencyIndexes: file_interservice_cfgmgmt_request_parameters_proto_depIdxs,
		EnumInfos:         file_interservice_cfgmgmt_request_parameters_proto_enumTypes,
		MessageInfos:      file_interservice_cfgmgmt_request_parameters_proto_msgTypes,
	}.Build()
	File_interservice_cfgmgmt_request_parameters_proto = out.File
	file_interservice_cfgmgmt_request_parameters_proto_rawDesc = nil
	file_interservice_cfgmgmt_request_parameters_proto_goTypes = nil
	file_interservice_cfgmgmt_request_parameters_proto_depIdxs = nil
}
