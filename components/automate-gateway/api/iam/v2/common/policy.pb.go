// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/common/policy.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Type int32

const (
	Type_CHEF_MANAGED Type = 0
	Type_CUSTOM       Type = 1
)

var Type_name = map[int32]string{
	0: "CHEF_MANAGED",
	1: "CUSTOM",
}

var Type_value = map[string]int32{
	"CHEF_MANAGED": 0,
	"CUSTOM":       1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{0}
}

type Flag int32

const (
	Flag_VERSION_2_0 Flag = 0
	Flag_VERSION_2_1 Flag = 1
)

var Flag_name = map[int32]string{
	0: "VERSION_2_0",
	1: "VERSION_2_1",
}

var Flag_value = map[string]int32{
	"VERSION_2_0": 0,
	"VERSION_2_1": 1,
}

func (x Flag) String() string {
	return proto.EnumName(Flag_name, int32(x))
}

func (Flag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{1}
}

type Statement_Effect int32

const (
	Statement_ALLOW Statement_Effect = 0
	Statement_DENY  Statement_Effect = 1
)

var Statement_Effect_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var Statement_Effect_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x Statement_Effect) String() string {
	return proto.EnumName(Statement_Effect_name, int32(x))
}

func (Statement_Effect) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{1, 0}
}

type Version_VersionNumber int32

const (
	Version_V0 Version_VersionNumber = 0
	Version_V1 Version_VersionNumber = 1
	Version_V2 Version_VersionNumber = 2
)

var Version_VersionNumber_name = map[int32]string{
	0: "V0",
	1: "V1",
	2: "V2",
}

var Version_VersionNumber_value = map[string]int32{
	"V0": 0,
	"V1": 1,
	"V2": 2,
}

func (x Version_VersionNumber) String() string {
	return proto.EnumName(Version_VersionNumber_name, int32(x))
}

func (Version_VersionNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{4, 0}
}

type Policy struct {
	// Name for the policy.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// This doc-comment is ignored for an enum.
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// Members affected by this policy. May be empty.
	Members []string `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	// Statements for the policy. Will contain one or more.
	Statements []*Statement `protobuf:"bytes,5,rep,name=statements,proto3" json:"statements,omitempty"`
	// List of projects this policy belongs to. May be empty.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{0}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Policy) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Policy) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Policy) GetStatements() []*Statement {
	if m != nil {
		return m.Statements
	}
	return nil
}

func (m *Policy) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Statement struct {
	// This doc-comment is ignored for an enum.
	Effect Statement_Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=chef.automate.api.iam.v2.Statement_Effect" json:"effect,omitempty"`
	// Actions defined inline. May be empty.
	// Best practices recommend that you use custom roles rather than inline actions where practical.
	Actions []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	// The role defines a set of actions that the statement is scoped to.
	Role string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// DEPRECATED: Resources defined inline. Use projects instead.
	Resources []string `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
	// The project list defines the set of resources that the statement is scoped to. May be empty.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{1}
}

func (m *Statement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statement.Unmarshal(m, b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return xxx_messageInfo_Statement.Size(m)
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetEffect() Statement_Effect {
	if m != nil {
		return m.Effect
	}
	return Statement_ALLOW
}

func (m *Statement) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Statement) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Statement) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Statement) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Role struct {
	// Name for the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Whether this policy is user created (`CUSTOM`) or chef managed (`CHEF_MANAGED`).
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// List of actions this role scopes to. Will contain one or more.
	Actions []string `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	// List of projects this role belongs to. May be empty.
	Projects             []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{2}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Role) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Role) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type Project struct {
	// Name for the project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Whether this policy is user created (`CUSTOM`) or chef managed (`CHEF_MANAGED`).
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// The current status of the rules for this project.
	Status               ProjectRulesStatus `protobuf:"varint,4,opt,name=status,proto3,enum=chef.automate.api.iam.v2.ProjectRulesStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{3}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_CHEF_MANAGED
}

func (m *Project) GetStatus() ProjectRulesStatus {
	if m != nil {
		return m.Status
	}
	return ProjectRulesStatus_PROJECT_RULES_STATUS_UNSET
}

type Version struct {
	Major                Version_VersionNumber `protobuf:"varint,1,opt,name=major,proto3,enum=chef.automate.api.iam.v2.Version_VersionNumber" json:"major,omitempty"`
	Minor                Version_VersionNumber `protobuf:"varint,2,opt,name=minor,proto3,enum=chef.automate.api.iam.v2.Version_VersionNumber" json:"minor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f942825c854ad5, []int{4}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() Version_VersionNumber {
	if m != nil {
		return m.Major
	}
	return Version_V0
}

func (m *Version) GetMinor() Version_VersionNumber {
	if m != nil {
		return m.Minor
	}
	return Version_V0
}

func init() {
	proto.RegisterEnum("chef.automate.api.iam.v2.Type", Type_name, Type_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.Flag", Flag_name, Flag_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.Statement_Effect", Statement_Effect_name, Statement_Effect_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.Version_VersionNumber", Version_VersionNumber_name, Version_VersionNumber_value)
	proto.RegisterType((*Policy)(nil), "chef.automate.api.iam.v2.Policy")
	proto.RegisterType((*Statement)(nil), "chef.automate.api.iam.v2.Statement")
	proto.RegisterType((*Role)(nil), "chef.automate.api.iam.v2.Role")
	proto.RegisterType((*Project)(nil), "chef.automate.api.iam.v2.Project")
	proto.RegisterType((*Version)(nil), "chef.automate.api.iam.v2.Version")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/common/policy.proto", fileDescriptor_96f942825c854ad5)
}

var fileDescriptor_96f942825c854ad5 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0x9e, 0xdb, 0x34, 0x5d, 0xce, 0xfe, 0xbf, 0x44, 0xbe, 0x8a, 0x26, 0x40, 0x55, 0x41, 0xa2,
	0xaa, 0x20, 0x61, 0xe1, 0x0a, 0xee, 0xba, 0x36, 0x65, 0x48, 0x5b, 0x3b, 0xa5, 0xa3, 0x08, 0x6e,
	0x2a, 0x37, 0x73, 0x3b, 0x4f, 0x75, 0x1c, 0x25, 0xce, 0x50, 0xdf, 0x83, 0x77, 0xe0, 0x35, 0x78,
	0x0a, 0x6e, 0x78, 0x19, 0x64, 0x27, 0x2d, 0x2b, 0x52, 0x11, 0x43, 0xda, 0x95, 0x7d, 0x8e, 0xfc,
	0x7d, 0xe7, 0xfb, 0x8e, 0x8f, 0x0e, 0xbc, 0x89, 0x04, 0x4f, 0x44, 0x4c, 0x63, 0x99, 0x79, 0x24,
	0x97, 0x82, 0x13, 0x49, 0x5f, 0x2c, 0x88, 0xa4, 0x9f, 0xc9, 0xca, 0x23, 0x09, 0xf3, 0x18, 0xe1,
	0xde, 0x8d, 0xef, 0x45, 0x82, 0x73, 0x11, 0x7b, 0x89, 0x58, 0xb2, 0x68, 0xe5, 0x26, 0xa9, 0x90,
	0x02, 0x3b, 0xd1, 0x15, 0x9d, 0xbb, 0x6b, 0x94, 0x4b, 0x12, 0xe6, 0x32, 0xc2, 0xdd, 0x1b, 0xff,
	0xf0, 0xf5, 0xdd, 0x58, 0xd3, 0x7c, 0x49, 0xb3, 0x82, 0xb4, 0xf5, 0x03, 0x81, 0x79, 0xae, 0xab,
	0x60, 0x0c, 0x46, 0x4c, 0x38, 0x75, 0x50, 0x13, 0xb5, 0xad, 0x50, 0xdf, 0x71, 0x03, 0x2a, 0xec,
	0xd2, 0xa9, 0xe8, 0x4c, 0x85, 0x5d, 0x62, 0x1f, 0x0c, 0xb9, 0x4a, 0xa8, 0x53, 0x6d, 0xa2, 0x76,
	0xc3, 0x7f, 0xec, 0xee, 0x92, 0xe4, 0x5e, 0xac, 0x12, 0x1a, 0xea, 0xb7, 0xd8, 0x81, 0x3a, 0xa7,
	0x7c, 0x46, 0xd3, 0xcc, 0x31, 0x9a, 0xd5, 0xb6, 0x15, 0xae, 0x43, 0xdc, 0x03, 0xc8, 0x24, 0x91,
	0x94, 0x2b, 0xe5, 0x4e, 0xad, 0x59, 0x6d, 0x1f, 0xf8, 0x4f, 0x76, 0x73, 0x8e, 0xd7, 0x6f, 0xc3,
	0x5b, 0x30, 0x7c, 0x08, 0xfb, 0x49, 0x2a, 0xae, 0x69, 0x24, 0x33, 0xc7, 0xd4, 0xfc, 0x9b, 0xb8,
	0xf5, 0x1d, 0x81, 0xb5, 0x41, 0xe1, 0x63, 0x30, 0xe9, 0x7c, 0x4e, 0x23, 0xa9, 0x2d, 0x36, 0xfc,
	0xce, 0x5f, 0x94, 0x72, 0x03, 0x8d, 0x08, 0x4b, 0xa4, 0x32, 0x43, 0x22, 0xc9, 0x44, 0x9c, 0x39,
	0xd5, 0xc2, 0x4c, 0x19, 0xaa, 0xf6, 0xa5, 0x62, 0x49, 0x1d, 0xa3, 0x68, 0x9f, 0xba, 0xe3, 0x87,
	0x60, 0xa5, 0x34, 0x13, 0x79, 0x1a, 0xd1, 0xc2, 0x9f, 0x15, 0xfe, 0x4a, 0xfc, 0x51, 0xf9, 0x23,
	0x30, 0x8b, 0xca, 0xd8, 0x82, 0x5a, 0xf7, 0xf4, 0x74, 0xf4, 0xc1, 0xde, 0xc3, 0xfb, 0x60, 0xf4,
	0x83, 0xe1, 0x47, 0x1b, 0xb5, 0xbe, 0x20, 0x30, 0x42, 0x55, 0xe1, 0x1e, 0x3f, 0x6d, 0xed, 0xd3,
	0xd8, 0xf6, 0x79, 0x5b, 0x75, 0xed, 0x37, 0xd5, 0x5f, 0x11, 0xd4, 0xcf, 0x8b, 0xe0, 0xde, 0x94,
	0xf5, 0xc1, 0x54, 0xbf, 0x9f, 0x67, 0xba, 0xd3, 0x0d, 0xff, 0xf9, 0x6e, 0x54, 0x29, 0x25, 0x54,
	0xf3, 0x3e, 0xd6, 0x98, 0xb0, 0xc4, 0xb6, 0xbe, 0x21, 0xa8, 0x4f, 0x68, 0x9a, 0x31, 0x11, 0xe3,
	0x00, 0x6a, 0x9c, 0x5c, 0x8b, 0xb4, 0x1c, 0x0b, 0x6f, 0x37, 0x61, 0x89, 0x58, 0x9f, 0xc3, 0x5c,
	0xcd, 0x71, 0x58, 0xa0, 0x35, 0x0d, 0x8b, 0x45, 0xaa, 0xfd, 0xfd, 0x13, 0x8d, 0x42, 0xb7, 0x9e,
	0xc1, 0xff, 0x5b, 0x79, 0x6c, 0x42, 0x65, 0xf2, 0xd2, 0xde, 0xd3, 0xe7, 0x91, 0x8d, 0xf4, 0xe9,
	0xdb, 0x95, 0xce, 0x53, 0x30, 0x54, 0x5b, 0xb0, 0x0d, 0xff, 0xf5, 0x4e, 0x82, 0xc1, 0xf4, 0xac,
	0x3b, 0xec, 0xbe, 0x0d, 0xfa, 0xf6, 0x1e, 0x06, 0x30, 0x7b, 0xef, 0xc7, 0x17, 0xa3, 0x33, 0x1b,
	0x75, 0xda, 0x60, 0x0c, 0x96, 0x64, 0x81, 0x1f, 0xc0, 0xc1, 0x24, 0x08, 0xc7, 0xef, 0x46, 0xc3,
	0xa9, 0x3f, 0x55, 0x74, 0x5b, 0x89, 0x23, 0x1b, 0x1d, 0x9f, 0x7c, 0x1a, 0x2c, 0x98, 0xbc, 0xca,
	0x67, 0x6e, 0x24, 0xb8, 0xa7, 0xc4, 0x6f, 0x96, 0x89, 0x77, 0xa7, 0x05, 0x33, 0x33, 0xf5, 0x6e,
	0x79, 0xf5, 0x33, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x0a, 0x42, 0xa5, 0xee, 0x04, 0x00, 0x00,
}
