// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0--rc3
// source: workflow/tasks/tasks.proto

package tasksv1

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
	return file_workflow_tasks_tasks_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_workflow_tasks_tasks_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{0}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Email       string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Status      TaskStatus `protobuf:"varint,5,opt,name=status,proto3,enum=tasks.TaskStatus" json:"status,omitempty"`
	Case        *Case      `protobuf:"bytes,6,opt,name=Case,proto3" json:"Case,omitempty"`
	Cluster     *Cluster   `protobuf:"bytes,7,opt,name=cluster,proto3" json:"cluster,omitempty"`
	CreatedAt   string     `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	FormedAt    string     `protobuf:"bytes,9,opt,name=formed_at,json=formedAt,proto3" json:"formed_at,omitempty"`
	CompletedAt string     `protobuf:"bytes,10,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	User        *User      `protobuf:"bytes,11,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[0]
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
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{0}
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

func (x *Task) GetEmail() string {
	if x != nil {
		return x.Email
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

func (x *Task) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
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
	if x != nil {
		return x.FormedAt
	}
	return ""
}

func (x *Task) GetCompletedAt() string {
	if x != nil {
		return x.CompletedAt
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
		mi := &file_workflow_tasks_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[1]
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
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{1}
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

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Frequency int64   `protobuf:"varint,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Cases     []*Case `protobuf:"bytes,4,rep,name=cases,proto3" json:"cases,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[2]
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
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{2}
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

func (x *Cluster) GetCases() []*Case {
	if x != nil {
		return x.Cases
	}
	return nil
}

type Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterId int64  `protobuf:"varint,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Solution  string `protobuf:"bytes,5,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *Case) Reset() {
	*x = Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Case) ProtoMessage() {}

func (x *Case) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[3]
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
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{3}
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

func (x *Case) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ClusterId   int64  `protobuf:"varint,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{5}
}

func (x *GetTaskRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type ListTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListTasksRequest) Reset() {
	*x = ListTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksRequest) ProtoMessage() {}

func (x *ListTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksRequest.ProtoReflect.Descriptor instead.
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{6}
}

func (x *ListTasksRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ListTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ListTasksResponse) Reset() {
	*x = ListTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksResponse) ProtoMessage() {}

func (x *ListTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksResponse.ProtoReflect.Descriptor instead.
func (*ListTasksResponse) Descriptor() ([]byte, []int) {
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{7}
}

func (x *ListTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type ChangeTaskStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *ChangeTaskStatusRequest) Reset() {
	*x = ChangeTaskStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_tasks_tasks_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeTaskStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeTaskStatusRequest) ProtoMessage() {}

func (x *ChangeTaskStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_tasks_tasks_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeTaskStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeTaskStatusRequest) Descriptor() ([]byte, []int) {
	return file_workflow_tasks_tasks_proto_rawDescGZIP(), []int{8}
}

func (x *ChangeTaskStatusRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

var File_workflow_tasks_tasks_proto protoreflect.FileDescriptor

var file_workflow_tasks_tasks_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x22, 0xda, 0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x2c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6e,
	0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x22, 0x67,
	0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x22, 0x32, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x2a, 0x33, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x02, 0x32, 0xf2, 0x01, 0x0a, 0x0b,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x3e, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x42, 0x1b, 0x5a, 0x19, 0x67, 0x72, 0x65, 0x76, 0x74, 0x73, 0x6f, 0x76, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_tasks_tasks_proto_rawDescOnce sync.Once
	file_workflow_tasks_tasks_proto_rawDescData = file_workflow_tasks_tasks_proto_rawDesc
)

func file_workflow_tasks_tasks_proto_rawDescGZIP() []byte {
	file_workflow_tasks_tasks_proto_rawDescOnce.Do(func() {
		file_workflow_tasks_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_tasks_tasks_proto_rawDescData)
	})
	return file_workflow_tasks_tasks_proto_rawDescData
}

var file_workflow_tasks_tasks_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_workflow_tasks_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_workflow_tasks_tasks_proto_goTypes = []interface{}{
	(TaskStatus)(0),                 // 0: tasks.TaskStatus
	(*Task)(nil),                    // 1: tasks.Task
	(*User)(nil),                    // 2: tasks.User
	(*Cluster)(nil),                 // 3: tasks.Cluster
	(*Case)(nil),                    // 4: tasks.Case
	(*CreateTaskRequest)(nil),       // 5: tasks.CreateTaskRequest
	(*GetTaskRequest)(nil),          // 6: tasks.GetTaskRequest
	(*ListTasksRequest)(nil),        // 7: tasks.ListTasksRequest
	(*ListTasksResponse)(nil),       // 8: tasks.ListTasksResponse
	(*ChangeTaskStatusRequest)(nil), // 9: tasks.ChangeTaskStatusRequest
}
var file_workflow_tasks_tasks_proto_depIdxs = []int32{
	0,  // 0: tasks.Task.status:type_name -> tasks.TaskStatus
	4,  // 1: tasks.Task.Case:type_name -> tasks.Case
	3,  // 2: tasks.Task.cluster:type_name -> tasks.Cluster
	2,  // 3: tasks.Task.user:type_name -> tasks.User
	4,  // 4: tasks.Cluster.cases:type_name -> tasks.Case
	1,  // 5: tasks.ListTasksResponse.tasks:type_name -> tasks.Task
	5,  // 6: tasks.TaskService.CreateTask:input_type -> tasks.CreateTaskRequest
	6,  // 7: tasks.TaskService.GetTask:input_type -> tasks.GetTaskRequest
	7,  // 8: tasks.TaskService.ListTasks:input_type -> tasks.ListTasksRequest
	9,  // 9: tasks.TaskService.ChangeTaskStatus:input_type -> tasks.ChangeTaskStatusRequest
	1,  // 10: tasks.TaskService.CreateTask:output_type -> tasks.Task
	1,  // 11: tasks.TaskService.GetTask:output_type -> tasks.Task
	8,  // 12: tasks.TaskService.ListTasks:output_type -> tasks.ListTasksResponse
	1,  // 13: tasks.TaskService.ChangeTaskStatus:output_type -> tasks.Task
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_workflow_tasks_tasks_proto_init() }
func file_workflow_tasks_tasks_proto_init() {
	if File_workflow_tasks_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflow_tasks_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_workflow_tasks_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_workflow_tasks_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_workflow_tasks_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_workflow_tasks_tasks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_workflow_tasks_tasks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_workflow_tasks_tasks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTasksRequest); i {
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
		file_workflow_tasks_tasks_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTasksResponse); i {
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
		file_workflow_tasks_tasks_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeTaskStatusRequest); i {
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
			RawDescriptor: file_workflow_tasks_tasks_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_tasks_tasks_proto_goTypes,
		DependencyIndexes: file_workflow_tasks_tasks_proto_depIdxs,
		EnumInfos:         file_workflow_tasks_tasks_proto_enumTypes,
		MessageInfos:      file_workflow_tasks_tasks_proto_msgTypes,
	}.Build()
	File_workflow_tasks_tasks_proto = out.File
	file_workflow_tasks_tasks_proto_rawDesc = nil
	file_workflow_tasks_tasks_proto_goTypes = nil
	file_workflow_tasks_tasks_proto_depIdxs = nil
}
