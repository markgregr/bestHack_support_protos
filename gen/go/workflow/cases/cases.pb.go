// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc3
// source: workflow/cases/cases.proto

package casesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskStatus int32

const (
	TaskStatus_OPEN        TaskStatus = 0
	TaskStatus_IN_PROGRESS TaskStatus = 1
	TaskStatus_CLOSED      TaskStatus = 2
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "OPEN",
		1: "IN_PROGRESS",
		2: "CLOSED",
	}
	TaskStatus_value = map[string]int32{
		"OPEN":        0,
		"IN_PROGRESS": 1,
		"CLOSED":      2,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_workflow_cases_cases_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_workflow_cases_cases_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{0}
}

type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cluster  *Cluster `protobuf:"bytes,2,opt,name=cluster,proto3,oneof" json:"cluster,omitempty"`
	Title    string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Solution string   `protobuf:"bytes,5,opt,name=solution,proto3" json:"solution,omitempty"`
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

func (x *Case) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *Case) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Case) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Frequency int64  `protobuf:"varint,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{1}
}

func (x *Cluster) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetFrequency() int64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      TaskStatus `protobuf:"varint,5,opt,name=status,proto3,enum=cases.TaskStatus" json:"status,omitempty"`
	Case        *Case      `protobuf:"bytes,6,opt,name=case,proto3,oneof" json:"case,omitempty"`
	CreatedAt   string     `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	FormedAt    *string    `protobuf:"bytes,9,opt,name=formed_at,json=formedAt,proto3,oneof" json:"formed_at,omitempty"`
	CompletedAt *string    `protobuf:"bytes,10,opt,name=completed_at,json=completedAt,proto3,oneof" json:"completed_at,omitempty"`
	User        *User      `protobuf:"bytes,11,opt,name=user,proto3,oneof" json:"user,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_OPEN
}

func (x *Task) GetCase() *Case {
	if x != nil {
		return x.Case
	}
	return nil
}

func (x *Task) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Task) GetFormedAt() string {
	if x != nil && x.FormedAt != nil {
		return *x.FormedAt
	}
	return ""
}

func (x *Task) GetCompletedAt() string {
	if x != nil && x.CompletedAt != nil {
		return *x.CompletedAt
	}
	return ""
}

func (x *Task) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId int64  `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Solution  string `protobuf:"bytes,4,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *CreateCaseRequest) Reset() {
	*x = CreateCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCaseRequest) ProtoMessage() {}

func (x *CreateCaseRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateCaseRequest.ProtoReflect.Descriptor instead.
func (*CreateCaseRequest) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{4}
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
		mi := &file_workflow_cases_cases_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaseRequest) ProtoMessage() {}

func (x *GetCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[5]
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
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{5}
}

func (x *GetCaseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCasesFromClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCasesFromClusterRequest) Reset() {
	*x = GetCasesFromClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCasesFromClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCasesFromClusterRequest) ProtoMessage() {}

func (x *GetCasesFromClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCasesFromClusterRequest.ProtoReflect.Descriptor instead.
func (*GetCasesFromClusterRequest) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{6}
}

func (x *GetCasesFromClusterRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCasesFromClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cases []*Case `protobuf:"bytes,1,rep,name=cases,proto3" json:"cases,omitempty"`
}

func (x *GetCasesFromClusterResponse) Reset() {
	*x = GetCasesFromClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCasesFromClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCasesFromClusterResponse) ProtoMessage() {}

func (x *GetCasesFromClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCasesFromClusterResponse.ProtoReflect.Descriptor instead.
func (*GetCasesFromClusterResponse) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{7}
}

func (x *GetCasesFromClusterResponse) GetCases() []*Case {
	if x != nil {
		return x.Cases
	}
	return nil
}

type ListClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *ListClustersResponse) Reset() {
	*x = ListClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersResponse) ProtoMessage() {}

func (x *ListClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersResponse.ProtoReflect.Descriptor instead.
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{8}
}

func (x *ListClustersResponse) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type UpdateCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Solution string `protobuf:"bytes,4,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *UpdateCaseRequest) Reset() {
	*x = UpdateCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCaseRequest) ProtoMessage() {}

func (x *UpdateCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCaseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCaseRequest) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateCaseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCaseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateCaseRequest) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

type DeleteCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCaseRequest) Reset() {
	*x = DeleteCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_cases_cases_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCaseRequest) ProtoMessage() {}

func (x *DeleteCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_cases_cases_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCaseRequest.ProtoReflect.Descriptor instead.
func (*DeleteCaseRequest) Descriptor() ([]byte, []int) {
	return file_workflow_cases_cases_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCaseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_workflow_cases_cases_proto protoreflect.FileDescriptor

var file_workflow_cases_cases_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0xdf, 0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x24, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x63, 0x61,
	0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x64, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x55, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x33, 0x0a, 0x0a, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45,
	0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x02,
	0x32, 0xda, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x33, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x12, 0x18,
	0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a,
	0x19, 0x67, 0x72, 0x65, 0x76, 0x74, 0x73, 0x6f, 0x76, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x3b, 0x63, 0x61, 0x73, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_workflow_cases_cases_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_workflow_cases_cases_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_workflow_cases_cases_proto_goTypes = []interface{}{
	(TaskStatus)(0),                     // 0: cases.TaskStatus
	(*Case)(nil),                        // 1: cases.Case
	(*Cluster)(nil),                     // 2: cases.Cluster
	(*Task)(nil),                        // 3: cases.Task
	(*User)(nil),                        // 4: cases.User
	(*CreateCaseRequest)(nil),           // 5: cases.CreateCaseRequest
	(*GetCaseRequest)(nil),              // 6: cases.GetCaseRequest
	(*GetCasesFromClusterRequest)(nil),  // 7: cases.GetCasesFromClusterRequest
	(*GetCasesFromClusterResponse)(nil), // 8: cases.GetCasesFromClusterResponse
	(*ListClustersResponse)(nil),        // 9: cases.ListClustersResponse
	(*UpdateCaseRequest)(nil),           // 10: cases.UpdateCaseRequest
	(*DeleteCaseRequest)(nil),           // 11: cases.DeleteCaseRequest
	(*emptypb.Empty)(nil),               // 12: google.protobuf.Empty
}
var file_workflow_cases_cases_proto_depIdxs = []int32{
	2,  // 0: cases.Case.cluster:type_name -> cases.Cluster
	0,  // 1: cases.Task.status:type_name -> cases.TaskStatus
	1,  // 2: cases.Task.case:type_name -> cases.Case
	4,  // 3: cases.Task.user:type_name -> cases.User
	1,  // 4: cases.GetCasesFromClusterResponse.cases:type_name -> cases.Case
	2,  // 5: cases.ListClustersResponse.clusters:type_name -> cases.Cluster
	5,  // 6: cases.CaseService.CreateCase:input_type -> cases.CreateCaseRequest
	10, // 7: cases.CaseService.UpdateCase:input_type -> cases.UpdateCaseRequest
	11, // 8: cases.CaseService.DeleteCase:input_type -> cases.DeleteCaseRequest
	12, // 9: cases.CaseService.ListClusters:input_type -> google.protobuf.Empty
	7,  // 10: cases.CaseService.GetCasesFromCluster:input_type -> cases.GetCasesFromClusterRequest
	1,  // 11: cases.CaseService.CreateCase:output_type -> cases.Case
	1,  // 12: cases.CaseService.UpdateCase:output_type -> cases.Case
	12, // 13: cases.CaseService.DeleteCase:output_type -> google.protobuf.Empty
	9,  // 14: cases.CaseService.ListClusters:output_type -> cases.ListClustersResponse
	8,  // 15: cases.CaseService.GetCasesFromCluster:output_type -> cases.GetCasesFromClusterResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
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
			switch v := v.(*Cluster); i {
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
			switch v := v.(*Task); i {
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
			switch v := v.(*User); i {
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
		file_workflow_cases_cases_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_workflow_cases_cases_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCasesFromClusterRequest); i {
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
		file_workflow_cases_cases_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCasesFromClusterResponse); i {
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
		file_workflow_cases_cases_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersResponse); i {
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
		file_workflow_cases_cases_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCaseRequest); i {
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
		file_workflow_cases_cases_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCaseRequest); i {
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
	file_workflow_cases_cases_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_workflow_cases_cases_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workflow_cases_cases_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_cases_cases_proto_goTypes,
		DependencyIndexes: file_workflow_cases_cases_proto_depIdxs,
		EnumInfos:         file_workflow_cases_cases_proto_enumTypes,
		MessageInfos:      file_workflow_cases_cases_proto_msgTypes,
	}.Build()
	File_workflow_cases_cases_proto = out.File
	file_workflow_cases_cases_proto_rawDesc = nil
	file_workflow_cases_cases_proto_goTypes = nil
	file_workflow_cases_cases_proto_depIdxs = nil
}
