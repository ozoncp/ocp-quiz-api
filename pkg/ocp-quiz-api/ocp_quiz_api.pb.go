// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ocp-quiz-api/ocp_quiz_api.proto

package ocp_quiz_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Quiz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ClassroomId uint64 `protobuf:"varint,2,opt,name=ClassroomId,proto3" json:"ClassroomId,omitempty"`
	UserId      uint64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Link        string `protobuf:"bytes,4,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *Quiz) Reset() {
	*x = Quiz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quiz) ProtoMessage() {}

func (x *Quiz) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quiz.ProtoReflect.Descriptor instead.
func (*Quiz) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{0}
}

func (x *Quiz) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Quiz) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *Quiz) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Quiz) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

// Request of Create new Quiz
type CreateQuizV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassroomId uint64 `protobuf:"varint,1,opt,name=ClassroomId,proto3" json:"ClassroomId,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *CreateQuizV1Request) Reset() {
	*x = CreateQuizV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuizV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuizV1Request) ProtoMessage() {}

func (x *CreateQuizV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuizV1Request.ProtoReflect.Descriptor instead.
func (*CreateQuizV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuizV1Request) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *CreateQuizV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateQuizV1Request) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

// Response of Create new Quiz
type CreateQuizV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizId uint64 `protobuf:"varint,1,opt,name=QuizId,proto3" json:"QuizId,omitempty"`
}

func (x *CreateQuizV1Response) Reset() {
	*x = CreateQuizV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuizV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuizV1Response) ProtoMessage() {}

func (x *CreateQuizV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuizV1Response.ProtoReflect.Descriptor instead.
func (*CreateQuizV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQuizV1Response) GetQuizId() uint64 {
	if x != nil {
		return x.QuizId
	}
	return 0
}

// Request of Describe Quiz
type DescribeQuizV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizId uint64 `protobuf:"varint,1,opt,name=QuizId,proto3" json:"QuizId,omitempty"`
}

func (x *DescribeQuizV1Request) Reset() {
	*x = DescribeQuizV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeQuizV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeQuizV1Request) ProtoMessage() {}

func (x *DescribeQuizV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeQuizV1Request.ProtoReflect.Descriptor instead.
func (*DescribeQuizV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeQuizV1Request) GetQuizId() uint64 {
	if x != nil {
		return x.QuizId
	}
	return 0
}

// Response of Describe Quiz
type DescribeQuizV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quiz *Quiz `protobuf:"bytes,1,opt,name=Quiz,proto3" json:"Quiz,omitempty"`
}

func (x *DescribeQuizV1Response) Reset() {
	*x = DescribeQuizV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeQuizV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeQuizV1Response) ProtoMessage() {}

func (x *DescribeQuizV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeQuizV1Response.ProtoReflect.Descriptor instead.
func (*DescribeQuizV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeQuizV1Response) GetQuiz() *Quiz {
	if x != nil {
		return x.Quiz
	}
	return nil
}

// Request on paginate for List of Quizes
// Limit - quizes on page
// Offset - page number
type ListQuizV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *ListQuizV1Request) Reset() {
	*x = ListQuizV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuizV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuizV1Request) ProtoMessage() {}

func (x *ListQuizV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuizV1Request.ProtoReflect.Descriptor instead.
func (*ListQuizV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListQuizV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListQuizV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Response on paginate for List of Quizes
// Quizes - founded quizes
// CurrentPage - current page
// IsLast - is last page?
type ListQuizV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quizes      []*Quiz `protobuf:"bytes,1,rep,name=Quizes,proto3" json:"Quizes,omitempty"`
	CurrentPage uint64  `protobuf:"varint,2,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	IsLast      bool    `protobuf:"varint,3,opt,name=IsLast,proto3" json:"IsLast,omitempty"`
}

func (x *ListQuizV1Response) Reset() {
	*x = ListQuizV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuizV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuizV1Response) ProtoMessage() {}

func (x *ListQuizV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuizV1Response.ProtoReflect.Descriptor instead.
func (*ListQuizV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListQuizV1Response) GetQuizes() []*Quiz {
	if x != nil {
		return x.Quizes
	}
	return nil
}

func (x *ListQuizV1Response) GetCurrentPage() uint64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *ListQuizV1Response) GetIsLast() bool {
	if x != nil {
		return x.IsLast
	}
	return false
}

// Request on remove quiz by Id
type RemoveQuizV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuizId uint64 `protobuf:"varint,1,opt,name=QuizId,proto3" json:"QuizId,omitempty"`
}

func (x *RemoveQuizV1Request) Reset() {
	*x = RemoveQuizV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveQuizV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveQuizV1Request) ProtoMessage() {}

func (x *RemoveQuizV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveQuizV1Request.ProtoReflect.Descriptor instead.
func (*RemoveQuizV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveQuizV1Request) GetQuizId() uint64 {
	if x != nil {
		return x.QuizId
	}
	return 0
}

// Response on remove quiz
// Found - something deleted?
type RemoveQuizV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=Found,proto3" json:"Found,omitempty"`
}

func (x *RemoveQuizV1Response) Reset() {
	*x = RemoveQuizV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveQuizV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveQuizV1Response) ProtoMessage() {}

func (x *RemoveQuizV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveQuizV1Response.ProtoReflect.Descriptor instead.
func (*RemoveQuizV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveQuizV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

var File_api_ocp_quiz_api_ocp_quiz_api_proto protoreflect.FileDescriptor

var file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x71, 0x75, 0x69, 0x7a, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x5f, 0x71, 0x75, 0x69, 0x7a, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x7e, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x0b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x0b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x51, 0x75, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x51, 0x75, 0x69, 0x7a, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x51, 0x75, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x51, 0x75,
	0x69, 0x7a, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x51, 0x75, 0x69, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x69, 0x7a,
	0x52, 0x04, 0x51, 0x75, 0x69, 0x7a, 0x22, 0x53, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75,
	0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x01, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x01, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x7a, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x51, 0x75, 0x69, 0x7a, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x06, 0x51, 0x75, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x49, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x06, 0x51, 0x75, 0x69, 0x7a, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x06, 0x51, 0x75, 0x69, 0x7a, 0x49, 0x64, 0x22,
	0x2c, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0xd8, 0x03,
	0x0a, 0x11, 0x4f, 0x63, 0x70, 0x51, 0x75, 0x69, 0x7a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69,
	0x7a, 0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a,
	0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x23, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2f, 0x7b, 0x51, 0x75, 0x69, 0x7a,
	0x49, 0x64, 0x7d, 0x12, 0x70, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x12,
	0x1f, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x71, 0x75, 0x69, 0x7a, 0x2f, 0x7b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x7d, 0x2f, 0x7b, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x7d, 0x12, 0x71, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x51,
	0x75, 0x69, 0x7a, 0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x51, 0x75, 0x69, 0x7a, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x71, 0x75, 0x69,
	0x7a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x51, 0x75, 0x69, 0x7a,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2f, 0x7b, 0x51, 0x75,
	0x69, 0x7a, 0x49, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63,
	0x70, 0x2d, 0x71, 0x75, 0x69, 0x7a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f,
	0x63, 0x70, 0x2d, 0x71, 0x75, 0x69, 0x7a, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f,
	0x71, 0x75, 0x69, 0x7a, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescOnce sync.Once
	file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescData = file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDesc
)

func file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescGZIP() []byte {
	file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescData)
	})
	return file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDescData
}

var file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ocp_quiz_api_ocp_quiz_api_proto_goTypes = []interface{}{
	(*Quiz)(nil),                   // 0: ocp.quiz.api.Quiz
	(*CreateQuizV1Request)(nil),    // 1: ocp.quiz.api.CreateQuizV1Request
	(*CreateQuizV1Response)(nil),   // 2: ocp.quiz.api.CreateQuizV1Response
	(*DescribeQuizV1Request)(nil),  // 3: ocp.quiz.api.DescribeQuizV1Request
	(*DescribeQuizV1Response)(nil), // 4: ocp.quiz.api.DescribeQuizV1Response
	(*ListQuizV1Request)(nil),      // 5: ocp.quiz.api.ListQuizV1Request
	(*ListQuizV1Response)(nil),     // 6: ocp.quiz.api.ListQuizV1Response
	(*RemoveQuizV1Request)(nil),    // 7: ocp.quiz.api.RemoveQuizV1Request
	(*RemoveQuizV1Response)(nil),   // 8: ocp.quiz.api.RemoveQuizV1Response
}
var file_api_ocp_quiz_api_ocp_quiz_api_proto_depIdxs = []int32{
	0, // 0: ocp.quiz.api.DescribeQuizV1Response.Quiz:type_name -> ocp.quiz.api.Quiz
	0, // 1: ocp.quiz.api.ListQuizV1Response.Quizes:type_name -> ocp.quiz.api.Quiz
	1, // 2: ocp.quiz.api.OcpQuizApiService.CreateQuiz:input_type -> ocp.quiz.api.CreateQuizV1Request
	3, // 3: ocp.quiz.api.OcpQuizApiService.DescribeQuiz:input_type -> ocp.quiz.api.DescribeQuizV1Request
	5, // 4: ocp.quiz.api.OcpQuizApiService.ListQuiz:input_type -> ocp.quiz.api.ListQuizV1Request
	7, // 5: ocp.quiz.api.OcpQuizApiService.RemoveQuiz:input_type -> ocp.quiz.api.RemoveQuizV1Request
	2, // 6: ocp.quiz.api.OcpQuizApiService.CreateQuiz:output_type -> ocp.quiz.api.CreateQuizV1Response
	4, // 7: ocp.quiz.api.OcpQuizApiService.DescribeQuiz:output_type -> ocp.quiz.api.DescribeQuizV1Response
	6, // 8: ocp.quiz.api.OcpQuizApiService.ListQuiz:output_type -> ocp.quiz.api.ListQuizV1Response
	8, // 9: ocp.quiz.api.OcpQuizApiService.RemoveQuiz:output_type -> ocp.quiz.api.RemoveQuizV1Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ocp_quiz_api_ocp_quiz_api_proto_init() }
func file_api_ocp_quiz_api_ocp_quiz_api_proto_init() {
	if File_api_ocp_quiz_api_ocp_quiz_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quiz); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuizV1Request); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuizV1Response); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeQuizV1Request); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeQuizV1Response); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuizV1Request); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuizV1Response); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveQuizV1Request); i {
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
		file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveQuizV1Response); i {
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
			RawDescriptor: file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_quiz_api_ocp_quiz_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_quiz_api_ocp_quiz_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_quiz_api_ocp_quiz_api_proto_msgTypes,
	}.Build()
	File_api_ocp_quiz_api_ocp_quiz_api_proto = out.File
	file_api_ocp_quiz_api_ocp_quiz_api_proto_rawDesc = nil
	file_api_ocp_quiz_api_ocp_quiz_api_proto_goTypes = nil
	file_api_ocp_quiz_api_ocp_quiz_api_proto_depIdxs = nil
}
