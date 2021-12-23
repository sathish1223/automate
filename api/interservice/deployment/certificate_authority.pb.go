// -*- mode: protobuf; indent-tabs-mode: t; c-basic-offset: 8; tab-width: 8 -*-

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: interservice/deployment/certificate_authority.proto

package deployment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RootCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RootCertRequest) Reset() {
	*x = RootCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootCertRequest) ProtoMessage() {}

func (x *RootCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootCertRequest.ProtoReflect.Descriptor instead.
func (*RootCertRequest) Descriptor() ([]byte, []int) {
	return file_interservice_deployment_certificate_authority_proto_rawDescGZIP(), []int{0}
}

type RootCertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert string `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty" toml:"cert,omitempty" mapstructure:"cert,omitempty"`
}

func (x *RootCertResponse) Reset() {
	*x = RootCertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootCertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootCertResponse) ProtoMessage() {}

func (x *RootCertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootCertResponse.ProtoReflect.Descriptor instead.
func (*RootCertResponse) Descriptor() ([]byte, []int) {
	return file_interservice_deployment_certificate_authority_proto_rawDescGZIP(), []int{1}
}

func (x *RootCertResponse) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

type RegenerateRootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegenerateRootRequest) Reset() {
	*x = RegenerateRootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegenerateRootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegenerateRootRequest) ProtoMessage() {}

func (x *RegenerateRootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegenerateRootRequest.ProtoReflect.Descriptor instead.
func (*RegenerateRootRequest) Descriptor() ([]byte, []int) {
	return file_interservice_deployment_certificate_authority_proto_rawDescGZIP(), []int{2}
}

type RegenerateRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegenerateRootResponse) Reset() {
	*x = RegenerateRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegenerateRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegenerateRootResponse) ProtoMessage() {}

func (x *RegenerateRootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_deployment_certificate_authority_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegenerateRootResponse.ProtoReflect.Descriptor instead.
func (*RegenerateRootResponse) Descriptor() ([]byte, []int) {
	return file_interservice_deployment_certificate_authority_proto_rawDescGZIP(), []int{3}
}

var File_interservice_deployment_certificate_authority_proto protoreflect.FileDescriptor

var file_interservice_deployment_certificate_authority_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x6f, 0x6f,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72,
	0x74, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x95, 0x02, 0x0a, 0x1b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x30, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x36, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_deployment_certificate_authority_proto_rawDescOnce sync.Once
	file_interservice_deployment_certificate_authority_proto_rawDescData = file_interservice_deployment_certificate_authority_proto_rawDesc
)

func file_interservice_deployment_certificate_authority_proto_rawDescGZIP() []byte {
	file_interservice_deployment_certificate_authority_proto_rawDescOnce.Do(func() {
		file_interservice_deployment_certificate_authority_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_deployment_certificate_authority_proto_rawDescData)
	})
	return file_interservice_deployment_certificate_authority_proto_rawDescData
}

var file_interservice_deployment_certificate_authority_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_interservice_deployment_certificate_authority_proto_goTypes = []interface{}{
	(*RootCertRequest)(nil),        // 0: chef.automate.domain.deployment.RootCertRequest
	(*RootCertResponse)(nil),       // 1: chef.automate.domain.deployment.RootCertResponse
	(*RegenerateRootRequest)(nil),  // 2: chef.automate.domain.deployment.RegenerateRootRequest
	(*RegenerateRootResponse)(nil), // 3: chef.automate.domain.deployment.RegenerateRootResponse
}
var file_interservice_deployment_certificate_authority_proto_depIdxs = []int32{
	0, // 0: chef.automate.domain.deployment.CertificateAuthorityService.GetRootCert:input_type -> chef.automate.domain.deployment.RootCertRequest
	2, // 1: chef.automate.domain.deployment.CertificateAuthorityService.RegenerateRoot:input_type -> chef.automate.domain.deployment.RegenerateRootRequest
	1, // 2: chef.automate.domain.deployment.CertificateAuthorityService.GetRootCert:output_type -> chef.automate.domain.deployment.RootCertResponse
	3, // 3: chef.automate.domain.deployment.CertificateAuthorityService.RegenerateRoot:output_type -> chef.automate.domain.deployment.RegenerateRootResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_deployment_certificate_authority_proto_init() }
func file_interservice_deployment_certificate_authority_proto_init() {
	if File_interservice_deployment_certificate_authority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_deployment_certificate_authority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootCertRequest); i {
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
		file_interservice_deployment_certificate_authority_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RootCertResponse); i {
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
		file_interservice_deployment_certificate_authority_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegenerateRootRequest); i {
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
		file_interservice_deployment_certificate_authority_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegenerateRootResponse); i {
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
			RawDescriptor: file_interservice_deployment_certificate_authority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_deployment_certificate_authority_proto_goTypes,
		DependencyIndexes: file_interservice_deployment_certificate_authority_proto_depIdxs,
		MessageInfos:      file_interservice_deployment_certificate_authority_proto_msgTypes,
	}.Build()
	File_interservice_deployment_certificate_authority_proto = out.File
	file_interservice_deployment_certificate_authority_proto_rawDesc = nil
	file_interservice_deployment_certificate_authority_proto_goTypes = nil
	file_interservice_deployment_certificate_authority_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CertificateAuthorityServiceClient is the client API for CertificateAuthorityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAuthorityServiceClient interface {
	GetRootCert(ctx context.Context, in *RootCertRequest, opts ...grpc.CallOption) (*RootCertResponse, error)
	RegenerateRoot(ctx context.Context, in *RegenerateRootRequest, opts ...grpc.CallOption) (*RegenerateRootResponse, error)
}

type certificateAuthorityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateAuthorityServiceClient(cc grpc.ClientConnInterface) CertificateAuthorityServiceClient {
	return &certificateAuthorityServiceClient{cc}
}

func (c *certificateAuthorityServiceClient) GetRootCert(ctx context.Context, in *RootCertRequest, opts ...grpc.CallOption) (*RootCertResponse, error) {
	out := new(RootCertResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.deployment.CertificateAuthorityService/GetRootCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityServiceClient) RegenerateRoot(ctx context.Context, in *RegenerateRootRequest, opts ...grpc.CallOption) (*RegenerateRootResponse, error) {
	out := new(RegenerateRootResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.deployment.CertificateAuthorityService/RegenerateRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAuthorityServiceServer is the server API for CertificateAuthorityService service.
type CertificateAuthorityServiceServer interface {
	GetRootCert(context.Context, *RootCertRequest) (*RootCertResponse, error)
	RegenerateRoot(context.Context, *RegenerateRootRequest) (*RegenerateRootResponse, error)
}

// UnimplementedCertificateAuthorityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAuthorityServiceServer struct {
}

func (*UnimplementedCertificateAuthorityServiceServer) GetRootCert(context.Context, *RootCertRequest) (*RootCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootCert not implemented")
}
func (*UnimplementedCertificateAuthorityServiceServer) RegenerateRoot(context.Context, *RegenerateRootRequest) (*RegenerateRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateRoot not implemented")
}

func RegisterCertificateAuthorityServiceServer(s *grpc.Server, srv CertificateAuthorityServiceServer) {
	s.RegisterService(&_CertificateAuthorityService_serviceDesc, srv)
}

func _CertificateAuthorityService_GetRootCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).GetRootCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.deployment.CertificateAuthorityService/GetRootCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).GetRootCert(ctx, req.(*RootCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthorityService_RegenerateRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServiceServer).RegenerateRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.deployment.CertificateAuthorityService/RegenerateRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServiceServer).RegenerateRoot(ctx, req.(*RegenerateRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthorityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.deployment.CertificateAuthorityService",
	HandlerType: (*CertificateAuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRootCert",
			Handler:    _CertificateAuthorityService_GetRootCert_Handler,
		},
		{
			MethodName: "RegenerateRoot",
			Handler:    _CertificateAuthorityService_RegenerateRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/deployment/certificate_authority.proto",
}
