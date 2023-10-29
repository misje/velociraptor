// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

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

type ScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The queue we want to receive jobs on
	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	// First request must be "register" then for each completed job "completion"
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The request ID this is a completion for.
	Id       uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Priority int64  `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	// A json encoded response
	Response string `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	OrgId    string `protobuf:"bytes,7,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// An error message or "" for no error.
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *ScheduleRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ScheduleRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ScheduleRequest) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ScheduleRequest) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *ScheduleRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *ScheduleRequest) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// This represents a job request from the server to the minion,
type ScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id must be matched with the response.
	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Queue string `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	// A json serialized request suitable for the named queue.
	Job   string `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
	OrgId string `protobuf:"bytes,7,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
}

func (x *ScheduleResponse) Reset() {
	*x = ScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleResponse) ProtoMessage() {}

func (x *ScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleResponse.ProtoReflect.Descriptor instead.
func (*ScheduleResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ScheduleResponse) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *ScheduleResponse) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *ScheduleResponse) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

var File_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x61, 0x0a, 0x10, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x42, 0x31,
	0x5a, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63,
	0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_proto_rawDescData = file_scheduler_proto_rawDesc
)

func file_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_proto_rawDescData)
	})
	return file_scheduler_proto_rawDescData
}

var file_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scheduler_proto_goTypes = []interface{}{
	(*ScheduleRequest)(nil),  // 0: proto.ScheduleRequest
	(*ScheduleResponse)(nil), // 1: proto.ScheduleResponse
}
var file_scheduler_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scheduler_proto_init() }
func file_scheduler_proto_init() {
	if File_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequest); i {
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
		file_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleResponse); i {
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
			RawDescriptor: file_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_proto = out.File
	file_scheduler_proto_rawDesc = nil
	file_scheduler_proto_goTypes = nil
	file_scheduler_proto_depIdxs = nil
}