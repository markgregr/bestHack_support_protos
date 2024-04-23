// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc3
// source: workflow/cases/cases.proto

package casesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterId   int64  `protobuf:"varint,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Solution    string `protobuf:"bytes,5,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *Case) Reset() {
	*x = Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Case) ProtoMessage() {}

func (x *Case) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Case.ProtoReflect.Descriptor instead.
func (*Case) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{0}
}

func (x *Case) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Case) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *Case) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Case) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Case) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

type CreateCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId   int64  `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Solution    string `protobuf:"bytes,4,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *CreateCaseRequest) Reset() {
	*x = CreateCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCaseRequest) ProtoMessage() {}

func (x *CreateCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCaseRequest.ProtoReflect.Descriptor instead.
func (*CreateCaseRequest) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCaseRequest) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *CreateCaseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateCaseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCaseRequest) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

type GetCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCaseRequest) Reset() {
	*x = GetCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaseRequest) ProtoMessage() {}

func (x *GetCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaseRequest.ProtoReflect.Descriptor instead.
func (*GetCaseRequest) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{2}
}

func (x *GetCaseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListCasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId int64 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *ListCasesRequest) Reset() {
	*x = ListCasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCasesRequest) ProtoMessage() {}

func (x *ListCasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCasesRequest.ProtoReflect.Descriptor instead.
func (*ListCasesRequest) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{3}
}

func (x *ListCasesRequest) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

type ListCasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cases []*Case `protobuf:"bytes,1,rep,name=cases,proto3" json:"cases,omitempty"`
}

func (x *ListCasesResponse) Reset() {
	*x = ListCasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCasesResponse) ProtoMessage() {}

func (x *ListCasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCasesResponse.ProtoReflect.Descriptor instead.
func (*ListCasesResponse) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{4}
}

func (x *ListCasesResponse) GetCases() []*Case {
	if x != nil {
		return x.Cases
	}
	return nil
}

var File_workflow_cases_cases_proto protoreflect.FileDescriptor

var file_workflow_cases_cases_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x89, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x01, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x32, 0xb1, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x12, 0x18, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x72, 0x65, 0x76, 0x74, 0x73,
	0x6f, 0x76, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_cases_cases_proto_rawDescOnce sync.Once
	file_workflow_cases_cases_proto_rawDescData = file_workflow_cases_cases_proto_rawDesc
)

func file_workflow_cases_cases_proto_rawDescGZIP() []byte {
	file_workflow_cases_cases_proto_rawDescOnce.Do(func() {
		file_workflow_cases_cases_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_cases_cases_proto_rawDescData)
	})
	return file_workflow_cases_cases_proto_rawDescData
}

var file_workflow_cases_cases_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_workflow_cases_cases_proto_goTypes = []interface{}{
	(*Case)(nil),              // 0: cases.Case
	(*CreateCaseRequest)(nil), // 1: cases.CreateCaseRequest
	(*GetCaseRequest)(nil),    // 2: cases.GetCaseRequest
	(*ListCasesRequest)(nil),  // 3: cases.ListCasesRequest
	(*ListCasesResponse)(nil), // 4: cases.ListCasesResponse
}
var file_workflow_cases_cases_proto_depIdxs = []int32{
	0, // 0: cases.ListCasesResponse.cases:type_name -> cases.Case
	1, // 1: cases.CaseService.CreateCase:input_type -> cases.CreateCaseRequest
	2, // 2: cases.CaseService.GetCase:input_type -> cases.GetCaseRequest
	3, // 3: cases.CaseService.ListCases:input_type -> cases.ListCasesRequest
	0, // 4: cases.CaseService.CreateCase:output_type -> cases.Case
	0, // 5: cases.CaseService.GetCase:output_type -> cases.Case
	4, // 6: cases.CaseService.ListCases:output_type -> cases.ListCasesResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_workflow_cases_cases_proto_init() }
func file_workflow_cases_cases_proto_init() {
	if File_workflow_cases_cases_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflow_cases_cases_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Case); i {
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
		file_workflow_cases_cases_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCaseRequest); i {
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
		file_workflow_cases_cases_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaseRequest); i {
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
		file_workflow_cases_cases_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCasesRequest); i {
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
		file_workflow_cases_cases_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCasesResponse); i {
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
			RawDescriptor: file_workflow_cases_cases_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_cases_cases_proto_goTypes,
		DependencyIndexes: file_workflow_cases_cases_proto_depIdxs,
		MessageInfos:      file_workflow_cases_cases_proto_msgTypes,
	}.Build()
	File_workflow_cases_cases_proto = out.File
	file_workflow_cases_cases_proto_rawDesc = nil
	file_workflow_cases_cases_proto_goTypes = nil
	file_workflow_cases_cases_proto_depIdxs = nil
}