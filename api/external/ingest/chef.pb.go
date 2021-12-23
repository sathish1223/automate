// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: external/ingest/chef.proto

package ingest

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	version "github.com/chef/automate/api/external/common/version"
	request "github.com/chef/automate/api/external/ingest/request"
	response "github.com/chef/automate/api/external/ingest/response"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_external_ingest_chef_proto protoreflect.FileDescriptor

var file_external_ingest_chef_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xed,
	0x0a, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x66, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12,
	0xdb, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x52,
	0x75, 0x6e, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x75, 0x6e, 0x1a, 0x39, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x72, 0x75, 0x6e, 0x3a, 0x01, 0x2a, 0x8a,
	0xb5, 0x18, 0x21, 0x0a, 0x1f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x3a, 0x7b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x3a,
	0x72, 0x75, 0x6e, 0x73, 0x8a, 0xb5, 0x18, 0x15, 0x12, 0x13, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xd8, 0x01,
	0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3c, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x10, 0x0a, 0x0e, 0x69,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x8a, 0xb5, 0x18,
	0x17, 0x12, 0x15, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xd8, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x3c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22, 0x25,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x69, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x8a, 0xb5, 0x18, 0x15, 0x12, 0x13,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x87, 0x02, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x73, 0x12, 0x3b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x44, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x22, 0x30, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x2d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x3a,
	0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x8a, 0xb5, 0x18, 0x15, 0x12, 0x13, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0xef, 0x01,
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73,
	0x73, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73,
	0x73, 0x1a, 0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x76,
	0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x70, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65,
	0x66, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18,
	0x25, 0x0a, 0x23, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a,
	0x7b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x3a, 0x6c, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x8a, 0xb5, 0x18, 0x15, 0x12, 0x13, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x3a, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0xcc, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x59, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x8a, 0xb5, 0x18, 0x18, 0x0a, 0x16, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x8a,
	0xb5, 0x18, 0x1b, 0x12, 0x19, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x67, 0x65, 0x74, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65,
	0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_external_ingest_chef_proto_goTypes = []interface{}{
	(*request.Run)(nil),                                // 0: chef.automate.api.ingest.request.Run
	(*request.Action)(nil),                             // 1: chef.automate.api.ingest.request.Action
	(*request.Delete)(nil),                             // 2: chef.automate.api.ingest.request.Delete
	(*request.MultipleNodeDeleteRequest)(nil),          // 3: chef.automate.api.ingest.request.MultipleNodeDeleteRequest
	(*request.Liveness)(nil),                           // 4: chef.automate.api.ingest.request.Liveness
	(*version.VersionInfoRequest)(nil),                 // 5: chef.automate.api.common.version.VersionInfoRequest
	(*response.ProcessChefRunResponse)(nil),            // 6: chef.automate.api.ingest.response.ProcessChefRunResponse
	(*response.ProcessChefActionResponse)(nil),         // 7: chef.automate.api.ingest.response.ProcessChefActionResponse
	(*response.ProcessNodeDeleteResponse)(nil),         // 8: chef.automate.api.ingest.response.ProcessNodeDeleteResponse
	(*response.ProcessMultipleNodeDeleteResponse)(nil), // 9: chef.automate.api.ingest.response.ProcessMultipleNodeDeleteResponse
	(*response.ProcessLivenessResponse)(nil),           // 10: chef.automate.api.ingest.response.ProcessLivenessResponse
	(*version.VersionInfo)(nil),                        // 11: chef.automate.api.common.version.VersionInfo
}
var file_external_ingest_chef_proto_depIdxs = []int32{
	0,  // 0: chef.automate.api.ingest.ChefIngester.ProcessChefRun:input_type -> chef.automate.api.ingest.request.Run
	1,  // 1: chef.automate.api.ingest.ChefIngester.ProcessChefAction:input_type -> chef.automate.api.ingest.request.Action
	2,  // 2: chef.automate.api.ingest.ChefIngester.ProcessNodeDelete:input_type -> chef.automate.api.ingest.request.Delete
	3,  // 3: chef.automate.api.ingest.ChefIngester.ProcessMultipleNodeDeletes:input_type -> chef.automate.api.ingest.request.MultipleNodeDeleteRequest
	4,  // 4: chef.automate.api.ingest.ChefIngester.ProcessLivenessPing:input_type -> chef.automate.api.ingest.request.Liveness
	5,  // 5: chef.automate.api.ingest.ChefIngester.GetVersion:input_type -> chef.automate.api.common.version.VersionInfoRequest
	6,  // 6: chef.automate.api.ingest.ChefIngester.ProcessChefRun:output_type -> chef.automate.api.ingest.response.ProcessChefRunResponse
	7,  // 7: chef.automate.api.ingest.ChefIngester.ProcessChefAction:output_type -> chef.automate.api.ingest.response.ProcessChefActionResponse
	8,  // 8: chef.automate.api.ingest.ChefIngester.ProcessNodeDelete:output_type -> chef.automate.api.ingest.response.ProcessNodeDeleteResponse
	9,  // 9: chef.automate.api.ingest.ChefIngester.ProcessMultipleNodeDeletes:output_type -> chef.automate.api.ingest.response.ProcessMultipleNodeDeleteResponse
	10, // 10: chef.automate.api.ingest.ChefIngester.ProcessLivenessPing:output_type -> chef.automate.api.ingest.response.ProcessLivenessResponse
	11, // 11: chef.automate.api.ingest.ChefIngester.GetVersion:output_type -> chef.automate.api.common.version.VersionInfo
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_external_ingest_chef_proto_init() }
func file_external_ingest_chef_proto_init() {
	if File_external_ingest_chef_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_ingest_chef_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_ingest_chef_proto_goTypes,
		DependencyIndexes: file_external_ingest_chef_proto_depIdxs,
	}.Build()
	File_external_ingest_chef_proto = out.File
	file_external_ingest_chef_proto_rawDesc = nil
	file_external_ingest_chef_proto_goTypes = nil
	file_external_ingest_chef_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChefIngesterClient is the client API for ChefIngester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChefIngesterClient interface {
	ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error)
	ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error)
	ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error)
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
}

type chefIngesterClient struct {
	cc grpc.ClientConnInterface
}

func NewChefIngesterClient(cc grpc.ClientConnInterface) ChefIngesterClient {
	return &chefIngesterClient{cc}
}

func (c *chefIngesterClient) ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error) {
	out := new(response.ProcessChefRunResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessChefRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error) {
	out := new(response.ProcessChefActionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessChefAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error) {
	out := new(response.ProcessNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessNodeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error) {
	out := new(response.ProcessMultipleNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessMultipleNodeDeletes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error) {
	out := new(response.ProcessLivenessResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/ProcessLivenessPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.api.ingest.ChefIngester/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChefIngesterServer is the server API for ChefIngester service.
type ChefIngesterServer interface {
	ProcessChefRun(context.Context, *request.Run) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(context.Context, *request.Action) (*response.ProcessChefActionResponse, error)
	ProcessNodeDelete(context.Context, *request.Delete) (*response.ProcessNodeDeleteResponse, error)
	ProcessMultipleNodeDeletes(context.Context, *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessLivenessPing(context.Context, *request.Liveness) (*response.ProcessLivenessResponse, error)
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
}

// UnimplementedChefIngesterServer can be embedded to have forward compatible implementations.
type UnimplementedChefIngesterServer struct {
}

func (*UnimplementedChefIngesterServer) ProcessChefRun(context.Context, *request.Run) (*response.ProcessChefRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefRun not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessChefAction(context.Context, *request.Action) (*response.ProcessChefActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefAction not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessNodeDelete(context.Context, *request.Delete) (*response.ProcessNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessNodeDelete not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessMultipleNodeDeletes(context.Context, *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMultipleNodeDeletes not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessLivenessPing(context.Context, *request.Liveness) (*response.ProcessLivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessLivenessPing not implemented")
}
func (*UnimplementedChefIngesterServer) GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterChefIngesterServer(s *grpc.Server, srv ChefIngesterServer) {
	s.RegisterService(&_ChefIngester_serviceDesc, srv)
}

func _ChefIngester_ProcessChefRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Run)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessChefRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, req.(*request.Run))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessChefAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessChefAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, req.(*request.Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessNodeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Delete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessNodeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, req.(*request.Delete))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessMultipleNodeDeletes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.MultipleNodeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessMultipleNodeDeletes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, req.(*request.MultipleNodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessLivenessPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Liveness)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/ProcessLivenessPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, req.(*request.Liveness))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.ingest.ChefIngester/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChefIngester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.ingest.ChefIngester",
	HandlerType: (*ChefIngesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessChefRun",
			Handler:    _ChefIngester_ProcessChefRun_Handler,
		},
		{
			MethodName: "ProcessChefAction",
			Handler:    _ChefIngester_ProcessChefAction_Handler,
		},
		{
			MethodName: "ProcessNodeDelete",
			Handler:    _ChefIngester_ProcessNodeDelete_Handler,
		},
		{
			MethodName: "ProcessMultipleNodeDeletes",
			Handler:    _ChefIngester_ProcessMultipleNodeDeletes_Handler,
		},
		{
			MethodName: "ProcessLivenessPing",
			Handler:    _ChefIngester_ProcessLivenessPing_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ChefIngester_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/ingest/chef.proto",
}
