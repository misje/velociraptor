// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetArtifactsRequest struct {
	// Deprecated
	IncludeEventArtifacts  bool     `protobuf:"varint,1,opt,name=include_event_artifacts,json=includeEventArtifacts,proto3" json:"include_event_artifacts,omitempty"`
	IncludeServerArtifacts bool     `protobuf:"varint,2,opt,name=include_server_artifacts,json=includeServerArtifacts,proto3" json:"include_server_artifacts,omitempty"`
	SearchTerm             string   `protobuf:"bytes,3,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	NumberOfResults        uint64   `protobuf:"varint,4,opt,name=number_of_results,json=numberOfResults,proto3" json:"number_of_results,omitempty"`
	Type                   string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetArtifactsRequest) Reset()         { *m = GetArtifactsRequest{} }
func (m *GetArtifactsRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactsRequest) ProtoMessage()    {}
func (*GetArtifactsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_a25917c1f4a78cbd, []int{0}
}
func (m *GetArtifactsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactsRequest.Unmarshal(m, b)
}
func (m *GetArtifactsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactsRequest.Marshal(b, m, deterministic)
}
func (dst *GetArtifactsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactsRequest.Merge(dst, src)
}
func (m *GetArtifactsRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactsRequest.Size(m)
}
func (m *GetArtifactsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactsRequest proto.InternalMessageInfo

func (m *GetArtifactsRequest) GetIncludeEventArtifacts() bool {
	if m != nil {
		return m.IncludeEventArtifacts
	}
	return false
}

func (m *GetArtifactsRequest) GetIncludeServerArtifacts() bool {
	if m != nil {
		return m.IncludeServerArtifacts
	}
	return false
}

func (m *GetArtifactsRequest) GetSearchTerm() string {
	if m != nil {
		return m.SearchTerm
	}
	return ""
}

func (m *GetArtifactsRequest) GetNumberOfResults() uint64 {
	if m != nil {
		return m.NumberOfResults
	}
	return 0
}

func (m *GetArtifactsRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetArtifactRequest struct {
	// Deprecated.
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactRequest) Reset()         { *m = GetArtifactRequest{} }
func (m *GetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactRequest) ProtoMessage()    {}
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_a25917c1f4a78cbd, []int{1}
}
func (m *GetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactRequest.Unmarshal(m, b)
}
func (m *GetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *GetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactRequest.Merge(dst, src)
}
func (m *GetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactRequest.Size(m)
}
func (m *GetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactRequest proto.InternalMessageInfo

func (m *GetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

type GetArtifactResponse struct {
	Artifact             string   `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactResponse) Reset()         { *m = GetArtifactResponse{} }
func (m *GetArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactResponse) ProtoMessage()    {}
func (*GetArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_a25917c1f4a78cbd, []int{2}
}
func (m *GetArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactResponse.Unmarshal(m, b)
}
func (m *GetArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactResponse.Marshal(b, m, deterministic)
}
func (dst *GetArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactResponse.Merge(dst, src)
}
func (m *GetArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactResponse.Size(m)
}
func (m *GetArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactResponse proto.InternalMessageInfo

func (m *GetArtifactResponse) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type SetArtifactRequest struct {
	VfsPath              string   `protobuf:"bytes,1,opt,name=vfs_path,json=vfsPath,proto3" json:"vfs_path,omitempty"`
	Artifact             string   `protobuf:"bytes,2,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetArtifactRequest) Reset()         { *m = SetArtifactRequest{} }
func (m *SetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*SetArtifactRequest) ProtoMessage()    {}
func (*SetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_a25917c1f4a78cbd, []int{3}
}
func (m *SetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetArtifactRequest.Unmarshal(m, b)
}
func (m *SetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *SetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetArtifactRequest.Merge(dst, src)
}
func (m *SetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_SetArtifactRequest.Size(m)
}
func (m *SetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetArtifactRequest proto.InternalMessageInfo

func (m *SetArtifactRequest) GetVfsPath() string {
	if m != nil {
		return m.VfsPath
	}
	return ""
}

func (m *SetArtifactRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

type APIResponse struct {
	Error                bool     `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIResponse) Reset()         { *m = APIResponse{} }
func (m *APIResponse) String() string { return proto.CompactTextString(m) }
func (*APIResponse) ProtoMessage()    {}
func (*APIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_a25917c1f4a78cbd, []int{4}
}
func (m *APIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIResponse.Unmarshal(m, b)
}
func (m *APIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIResponse.Marshal(b, m, deterministic)
}
func (dst *APIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIResponse.Merge(dst, src)
}
func (m *APIResponse) XXX_Size() int {
	return xxx_messageInfo_APIResponse.Size(m)
}
func (m *APIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_APIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_APIResponse proto.InternalMessageInfo

func (m *APIResponse) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *APIResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type GetReportRequest struct {
	Artifact string `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Format   string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	// Common parameters
	ClientId string `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Parameters for MONITORING_DAILY
	DayName string `protobuf:"bytes,6,opt,name=day_name,json=dayName,proto3" json:"day_name,omitempty"`
	// Parameters for CLIENT
	FlowId               string           `protobuf:"bytes,7,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Parameters           []*proto1.VQLEnv `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetReportRequest) Reset()         { *m = GetReportRequest{} }
func (m *GetReportRequest) String() string { return proto.CompactTextString(m) }
func (*GetReportRequest) ProtoMessage()    {}
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_a25917c1f4a78cbd, []int{5}
}
func (m *GetReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportRequest.Unmarshal(m, b)
}
func (m *GetReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportRequest.Marshal(b, m, deterministic)
}
func (dst *GetReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportRequest.Merge(dst, src)
}
func (m *GetReportRequest) XXX_Size() int {
	return xxx_messageInfo_GetReportRequest.Size(m)
}
func (m *GetReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportRequest proto.InternalMessageInfo

func (m *GetReportRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

func (m *GetReportRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetReportRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *GetReportRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetReportRequest) GetDayName() string {
	if m != nil {
		return m.DayName
	}
	return ""
}

func (m *GetReportRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *GetReportRequest) GetParameters() []*proto1.VQLEnv {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// This presents the report in a form that can be rendered in the
// GUI. The data is presented in two parts - first "data" contains a
// json encoded object, then "template" is an angular template to be
// evaluated with the data.
type GetReportResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Template             string   `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Messages             []string `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReportResponse) Reset()         { *m = GetReportResponse{} }
func (m *GetReportResponse) String() string { return proto.CompactTextString(m) }
func (*GetReportResponse) ProtoMessage()    {}
func (*GetReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_a25917c1f4a78cbd, []int{6}
}
func (m *GetReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportResponse.Unmarshal(m, b)
}
func (m *GetReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportResponse.Marshal(b, m, deterministic)
}
func (dst *GetReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportResponse.Merge(dst, src)
}
func (m *GetReportResponse) XXX_Size() int {
	return xxx_messageInfo_GetReportResponse.Size(m)
}
func (m *GetReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportResponse proto.InternalMessageInfo

func (m *GetReportResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *GetReportResponse) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *GetReportResponse) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*GetArtifactsRequest)(nil), "proto.GetArtifactsRequest")
	proto.RegisterType((*GetArtifactRequest)(nil), "proto.GetArtifactRequest")
	proto.RegisterType((*GetArtifactResponse)(nil), "proto.GetArtifactResponse")
	proto.RegisterType((*SetArtifactRequest)(nil), "proto.SetArtifactRequest")
	proto.RegisterType((*APIResponse)(nil), "proto.APIResponse")
	proto.RegisterType((*GetReportRequest)(nil), "proto.GetReportRequest")
	proto.RegisterType((*GetReportResponse)(nil), "proto.GetReportResponse")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_artifacts_a25917c1f4a78cbd) }

var fileDescriptor_artifacts_a25917c1f4a78cbd = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdf, 0x6e, 0xfb, 0x34,
	0x14, 0x56, 0xd6, 0xae, 0x7f, 0x3c, 0xa6, 0x6d, 0x46, 0xb0, 0x50, 0x90, 0xb0, 0x82, 0x84, 0x02,
	0x42, 0x29, 0x7f, 0x24, 0x98, 0x26, 0x40, 0xb4, 0xda, 0x98, 0x2a, 0xf6, 0x0f, 0xaf, 0x42, 0x42,
	0xbb, 0x88, 0xbc, 0xe4, 0xa4, 0x89, 0x94, 0xd8, 0xa9, 0xed, 0xa6, 0xdb, 0x35, 0x8f, 0xc1, 0x83,
	0xf0, 0x02, 0x3c, 0x09, 0xdc, 0xf0, 0x10, 0x5c, 0xa0, 0xd8, 0x69, 0xd6, 0xdd, 0x4d, 0x5c, 0xfc,
	0xae, 0xea, 0x9c, 0x73, 0xbe, 0xcf, 0xe7, 0x7c, 0xe7, 0x73, 0xd1, 0x01, 0x93, 0x3a, 0x4b, 0x58,
	0xa4, 0x55, 0x50, 0x4a, 0xa1, 0x05, 0xde, 0x35, 0x3f, 0xa3, 0xd3, 0xf5, 0x7a, 0x1d, 0x54, 0x90,
	0x8b, 0x28, 0x8b, 0xe1, 0x31, 0x88, 0x44, 0x31, 0x5e, 0x88, 0x9c, 0xf1, 0xc5, 0xd8, 0x06, 0x25,
	0x2b, 0xb5, 0x90, 0x63, 0x53, 0x3c, 0x56, 0x50, 0x30, 0xae, 0xb3, 0xc8, 0x52, 0x8c, 0xbe, 0x7b,
	0x1d, 0x96, 0x45, 0x3a, 0x13, 0x5c, 0x35, 0x1c, 0xd5, 0x32, 0xb7, 0x70, 0xef, 0x1f, 0x07, 0xbd,
	0x7d, 0x01, 0x7a, 0xb2, 0x69, 0x8c, 0xc2, 0x72, 0x05, 0x4a, 0xe3, 0xaf, 0xd1, 0x71, 0xc6, 0xa3,
	0x7c, 0x15, 0x43, 0x08, 0x15, 0x70, 0x1d, 0xb6, 0xad, 0xbb, 0x0e, 0x71, 0xfc, 0x01, 0x7d, 0xa7,
	0x49, 0x9f, 0xd7, 0xd9, 0x16, 0x8e, 0x4f, 0x90, 0xbb, 0xc1, 0x29, 0x90, 0x15, 0xc8, 0x2d, 0xe0,
	0x8e, 0x01, 0xbe, 0xdb, 0xe4, 0xef, 0x4c, 0xfa, 0x19, 0xf9, 0x21, 0xda, 0x53, 0xc0, 0x64, 0x94,
	0x86, 0x1a, 0x64, 0xe1, 0x76, 0x88, 0xe3, 0x0f, 0x29, 0xb2, 0xa1, 0x39, 0xc8, 0x02, 0x7f, 0x8a,
	0x8e, 0xf8, 0xaa, 0x78, 0x00, 0x19, 0x8a, 0x24, 0x94, 0xa0, 0x56, 0xb9, 0x56, 0x6e, 0x97, 0x38,
	0x7e, 0x97, 0x1e, 0xd8, 0xc4, 0x4d, 0x42, 0x6d, 0x18, 0x63, 0xd4, 0xd5, 0x4f, 0x25, 0xb8, 0xbb,
	0x86, 0xc5, 0x9c, 0xbd, 0x25, 0xc2, 0x5b, 0x93, 0x6e, 0x06, 0xbd, 0x47, 0x83, 0x2a, 0x51, 0x61,
	0xc9, 0x74, 0x6a, 0x26, 0x1b, 0x4e, 0x7f, 0xf8, 0xeb, 0xdf, 0xbf, 0xff, 0x74, 0x4e, 0xf1, 0xc9,
	0x3c, 0x05, 0x52, 0x25, 0x8a, 0xd4, 0x39, 0x22, 0x21, 0x67, 0x3a, 0xab, 0x80, 0x68, 0x41, 0x74,
	0x0a, 0xa4, 0x1d, 0x8b, 0xc4, 0x90, 0x64, 0x3c, 0xab, 0xe5, 0x25, 0x4a, 0x0b, 0x09, 0x01, 0xed,
	0x57, 0x89, 0xba, 0x65, 0x3a, 0xf5, 0xee, 0x5f, 0x88, 0x4b, 0x41, 0x95, 0x82, 0x2b, 0xc0, 0x67,
	0x68, 0xb0, 0x81, 0x37, 0x77, 0xfa, 0xe6, 0x4e, 0x0f, 0x93, 0xf9, 0x16, 0x35, 0x89, 0x99, 0x66,
	0x9f, 0x11, 0x21, 0x09, 0xab, 0x2f, 0x61, 0xab, 0x5c, 0x07, 0xb4, 0x45, 0x7a, 0x7f, 0x38, 0x08,
	0xdf, 0xbd, 0xd9, 0x81, 0x5e, 0x74, 0xbe, 0xf3, 0xbf, 0x3b, 0x7f, 0x44, 0x7b, 0x93, 0xdb, 0xd9,
	0x96, 0x1c, 0xbb, 0x20, 0xa5, 0x90, 0xd6, 0x59, 0xd3, 0xc0, 0x30, 0xfa, 0xf8, 0xe3, 0x09, 0x27,
	0x26, 0x4e, 0x44, 0x14, 0xad, 0x24, 0xc4, 0x44, 0x81, 0xd6, 0x19, 0x5f, 0xbc, 0x68, 0x37, 0xa0,
	0x16, 0x8c, 0x3f, 0x42, 0xfb, 0xe6, 0x10, 0x16, 0xa0, 0x14, 0x5b, 0x80, 0xed, 0x8f, 0xbe, 0x65,
	0x82, 0x57, 0x36, 0xe6, 0xfd, 0xde, 0x41, 0x87, 0x17, 0xa0, 0x29, 0x94, 0x42, 0xb6, 0x8a, 0xbd,
	0x76, 0x1d, 0x89, 0x90, 0x64, 0x9d, 0x66, 0x51, 0x4a, 0xd6, 0x40, 0xa4, 0xa5, 0x68, 0x91, 0xf8,
	0xa7, 0xc6, 0x72, 0x56, 0x96, 0x6f, 0x0c, 0xc3, 0x17, 0x78, 0x5c, 0x33, 0xd8, 0x5a, 0x52, 0xa7,
	0x6b, 0x28, 0x07, 0x88, 0x89, 0x0f, 0xc1, 0x22, 0x20, 0x57, 0x37, 0xd7, 0xb3, 0xf9, 0x0d, 0x9d,
	0x5d, 0x5f, 0x84, 0x67, 0x93, 0xd9, 0xe5, 0xaf, 0x9f, 0x58, 0xaf, 0xe2, 0xcf, 0x51, 0x2f, 0x11,
	0xb2, 0x60, 0xda, 0xbe, 0x83, 0xa9, 0x6b, 0xe8, 0x30, 0x3e, 0xfc, 0xd1, 0x44, 0x89, 0x01, 0xa7,
	0xba, 0xc8, 0x69, 0x53, 0x87, 0xdf, 0x47, 0xc3, 0x28, 0xcf, 0xea, 0x97, 0x9a, 0xc5, 0x8d, 0xed,
	0x07, 0x36, 0x30, 0x8b, 0xf1, 0x7b, 0x68, 0x10, 0xb3, 0xa7, 0x90, 0xb3, 0x02, 0xdc, 0x9e, 0xc9,
	0xf5, 0x63, 0xf6, 0x74, 0xcd, 0x0a, 0xc0, 0xc7, 0xa8, 0x9f, 0xe4, 0x62, 0x5d, 0xa3, 0xfa, 0x26,
	0xd3, 0xab, 0x3f, 0x67, 0x31, 0x5e, 0x22, 0x54, 0x32, 0xc9, 0x0a, 0xd0, 0x20, 0xeb, 0x77, 0xd6,
	0xf1, 0xf7, 0xbe, 0xdc, 0xb7, 0xff, 0x1a, 0xc1, 0x2f, 0x3f, 0x5f, 0x9e, 0xf3, 0x6a, 0x3a, 0x35,
	0x5d, 0x7d, 0x8b, 0x4f, 0xad, 0x9e, 0xe4, 0xb9, 0x3e, 0xa8, 0x95, 0x53, 0x40, 0x62, 0x28, 0x81,
	0xc7, 0x44, 0x70, 0xb3, 0x33, 0x23, 0x80, 0x48, 0xcc, 0xd9, 0x6a, 0x12, 0xd0, 0xad, 0x4b, 0xbc,
	0xdf, 0x1c, 0x74, 0xb4, 0xb5, 0x9d, 0xc6, 0x1e, 0x18, 0x75, 0x6b, 0x3b, 0xd9, 0xd5, 0x50, 0x73,
	0xc6, 0x23, 0x34, 0xd0, 0x50, 0x94, 0x39, 0xd3, 0x9b, 0x3d, 0xb7, 0xdf, 0xf8, 0x7b, 0x34, 0x68,
	0x2c, 0xa0, 0xdc, 0x0e, 0xe9, 0xf8, 0xc3, 0xa9, 0x67, 0xfa, 0xfc, 0x00, 0x8f, 0xce, 0xad, 0x9d,
	0x24, 0x59, 0x33, 0xc9, 0x6b, 0x27, 0x6d, 0x0a, 0x03, 0xda, 0x62, 0x1e, 0x7a, 0x66, 0xc6, 0xaf,
	0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xc2, 0x4e, 0x98, 0xae, 0x05, 0x00, 0x00,
}
