// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: interservice/cfgmgmt/request/inventory_nodes.proto

package request

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InventoryNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty" toml:"start,omitempty" mapstructure:"start,omitempty"`
	End        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty" toml:"end,omitempty" mapstructure:"end,omitempty"`
	Filter     []string               `protobuf:"bytes,3,rep,name=filter,proto3" json:"filter,omitempty" toml:"filter,omitempty" mapstructure:"filter,omitempty"`
	CursorDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=cursor_date,json=cursorDate,proto3" json:"cursor_date,omitempty" toml:"cursor_date,omitempty" mapstructure:"cursor_date,omitempty"`
	CursorId   string                 `protobuf:"bytes,5,opt,name=cursor_id,json=cursorId,proto3" json:"cursor_id,omitempty" toml:"cursor_id,omitempty" mapstructure:"cursor_id,omitempty"`
	PageSize   int32                  `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" toml:"page_size,omitempty" mapstructure:"page_size,omitempty"`
	Sorting    *Sorting               `protobuf:"bytes,7,opt,name=sorting,proto3" json:"sorting,omitempty" toml:"sorting,omitempty" mapstructure:"sorting,omitempty"`
}

func (x *InventoryNodes) Reset() {
	*x = InventoryNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_request_inventory_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryNodes) ProtoMessage() {}

func (x *InventoryNodes) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_request_inventory_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryNodes.ProtoReflect.Descriptor instead.
func (*InventoryNodes) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryNodes) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *InventoryNodes) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *InventoryNodes) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *InventoryNodes) GetCursorDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CursorDate
	}
	return nil
}

func (x *InventoryNodes) GetCursorId() string {
	if x != nil {
		return x.CursorId
	}
	return ""
}

func (x *InventoryNodes) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *InventoryNodes) GetSorting() *Sorting {
	if x != nil {
		return x.Sorting
	}
	return nil
}

var File_interservice_cfgmgmt_request_inventory_nodes_proto protoreflect.FileDescriptor

var file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDesc = []byte{
	0x0a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d,
	0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x0e, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x47, 0x0a, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescOnce sync.Once
	file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescData = file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDesc
)

func file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescGZIP() []byte {
	file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescOnce.Do(func() {
		file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescData)
	})
	return file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDescData
}

var file_interservice_cfgmgmt_request_inventory_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_interservice_cfgmgmt_request_inventory_nodes_proto_goTypes = []interface{}{
	(*InventoryNodes)(nil),        // 0: chef.automate.domain.cfgmgmt.request.InventoryNodes
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Sorting)(nil),               // 2: chef.automate.domain.cfgmgmt.request.Sorting
}
var file_interservice_cfgmgmt_request_inventory_nodes_proto_depIdxs = []int32{
	1, // 0: chef.automate.domain.cfgmgmt.request.InventoryNodes.start:type_name -> google.protobuf.Timestamp
	1, // 1: chef.automate.domain.cfgmgmt.request.InventoryNodes.end:type_name -> google.protobuf.Timestamp
	1, // 2: chef.automate.domain.cfgmgmt.request.InventoryNodes.cursor_date:type_name -> google.protobuf.Timestamp
	2, // 3: chef.automate.domain.cfgmgmt.request.InventoryNodes.sorting:type_name -> chef.automate.domain.cfgmgmt.request.Sorting
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_interservice_cfgmgmt_request_inventory_nodes_proto_init() }
func file_interservice_cfgmgmt_request_inventory_nodes_proto_init() {
	if File_interservice_cfgmgmt_request_inventory_nodes_proto != nil {
		return
	}
	file_interservice_cfgmgmt_request_parameters_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_interservice_cfgmgmt_request_inventory_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryNodes); i {
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
			RawDescriptor: file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_cfgmgmt_request_inventory_nodes_proto_goTypes,
		DependencyIndexes: file_interservice_cfgmgmt_request_inventory_nodes_proto_depIdxs,
		MessageInfos:      file_interservice_cfgmgmt_request_inventory_nodes_proto_msgTypes,
	}.Build()
	File_interservice_cfgmgmt_request_inventory_nodes_proto = out.File
	file_interservice_cfgmgmt_request_inventory_nodes_proto_rawDesc = nil
	file_interservice_cfgmgmt_request_inventory_nodes_proto_goTypes = nil
	file_interservice_cfgmgmt_request_inventory_nodes_proto_depIdxs = nil
}
