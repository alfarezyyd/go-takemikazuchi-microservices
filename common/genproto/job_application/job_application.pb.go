// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: job_application.proto

package job_application

import (
	user "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FindJobApplicationByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApplicantId   uint64                 `protobuf:"varint,1,opt,name=applicantId,proto3" json:"applicantId,omitempty"`
	JobId         uint64                 `protobuf:"varint,2,opt,name=jobId,proto3" json:"jobId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindJobApplicationByIdRequest) Reset() {
	*x = FindJobApplicationByIdRequest{}
	mi := &file_job_application_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindJobApplicationByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindJobApplicationByIdRequest) ProtoMessage() {}

func (x *FindJobApplicationByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindJobApplicationByIdRequest.ProtoReflect.Descriptor instead.
func (*FindJobApplicationByIdRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{0}
}

func (x *FindJobApplicationByIdRequest) GetApplicantId() uint64 {
	if x != nil {
		return x.ApplicantId
	}
	return 0
}

func (x *FindJobApplicationByIdRequest) GetJobId() uint64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

type JobApplicationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	JobId         uint64                 `protobuf:"varint,2,opt,name=JobId,proto3" json:"JobId,omitempty"`
	ApplicantId   uint64                 `protobuf:"varint,3,opt,name=ApplicantId,proto3" json:"ApplicantId,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobApplicationResponse) Reset() {
	*x = JobApplicationResponse{}
	mi := &file_job_application_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobApplicationResponse) ProtoMessage() {}

func (x *JobApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobApplicationResponse.ProtoReflect.Descriptor instead.
func (*JobApplicationResponse) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{1}
}

func (x *JobApplicationResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *JobApplicationResponse) GetJobId() uint64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *JobApplicationResponse) GetApplicantId() uint64 {
	if x != nil {
		return x.ApplicantId
	}
	return 0
}

func (x *JobApplicationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobApplicationResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *JobApplicationResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type FindAllApplicationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserJwtClaim  *user.UserJwtClaim     `protobuf:"bytes,1,opt,name=UserJwtClaim,proto3" json:"UserJwtClaim,omitempty"`
	JobId         string                 `protobuf:"bytes,2,opt,name=JobId,proto3" json:"JobId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindAllApplicationRequest) Reset() {
	*x = FindAllApplicationRequest{}
	mi := &file_job_application_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindAllApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllApplicationRequest) ProtoMessage() {}

func (x *FindAllApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllApplicationRequest.ProtoReflect.Descriptor instead.
func (*FindAllApplicationRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllApplicationRequest) GetUserJwtClaim() *user.UserJwtClaim {
	if x != nil {
		return x.UserJwtClaim
	}
	return nil
}

func (x *FindAllApplicationRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type SelectApplicationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserJwtClaim  *user.UserJwtClaim     `protobuf:"bytes,1,opt,name=UserJwtClaim,proto3" json:"UserJwtClaim,omitempty"`
	UserId        uint64                 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	JobId         uint64                 `protobuf:"varint,3,opt,name=JobId,proto3" json:"JobId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SelectApplicationRequest) Reset() {
	*x = SelectApplicationRequest{}
	mi := &file_job_application_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SelectApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectApplicationRequest) ProtoMessage() {}

func (x *SelectApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectApplicationRequest.ProtoReflect.Descriptor instead.
func (*SelectApplicationRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{3}
}

func (x *SelectApplicationRequest) GetUserJwtClaim() *user.UserJwtClaim {
	if x != nil {
		return x.UserJwtClaim
	}
	return nil
}

func (x *SelectApplicationRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SelectApplicationRequest) GetJobId() uint64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

type JobApplicationResponses struct {
	state                   protoimpl.MessageState    `protogen:"open.v1"`
	JobApplicationResponses []*JobApplicationResponse `protobuf:"bytes,1,rep,name=JobApplicationResponses,proto3" json:"JobApplicationResponses,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *JobApplicationResponses) Reset() {
	*x = JobApplicationResponses{}
	mi := &file_job_application_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobApplicationResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobApplicationResponses) ProtoMessage() {}

func (x *JobApplicationResponses) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobApplicationResponses.ProtoReflect.Descriptor instead.
func (*JobApplicationResponses) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{4}
}

func (x *JobApplicationResponses) GetJobApplicationResponses() []*JobApplicationResponse {
	if x != nil {
		return x.JobApplicationResponses
	}
	return nil
}

type ApplyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserJwtClaim  *user.UserJwtClaim     `protobuf:"bytes,1,opt,name=UserJwtClaim,proto3" json:"UserJwtClaim,omitempty"`
	JobId         uint64                 `protobuf:"varint,2,opt,name=JobId,proto3" json:"JobId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplyRequest) Reset() {
	*x = ApplyRequest{}
	mi := &file_job_application_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRequest) ProtoMessage() {}

func (x *ApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_application_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRequest.ProtoReflect.Descriptor instead.
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return file_job_application_proto_rawDescGZIP(), []int{5}
}

func (x *ApplyRequest) GetUserJwtClaim() *user.UserJwtClaim {
	if x != nil {
		return x.UserJwtClaim
	}
	return nil
}

func (x *ApplyRequest) GetJobId() uint64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

var File_job_application_proto protoreflect.FileDescriptor

var file_job_application_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x16, 0x4a, 0x6f,
	0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x64, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x18, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4a,
	0x77, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4a, 0x6f,
	0x62, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x17, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x51,
	0x0a, 0x17, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x17, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x73, 0x22, 0x57, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77,
	0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x32, 0xa6, 0x02, 0x0a, 0x15, 0x4a,
	0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1e, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4a, 0x6f,
	0x62, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a,
	0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x0d, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x66, 0x61, 0x72, 0x65, 0x7a, 0x79, 0x79, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x74, 0x61, 0x6b, 0x65, 0x6d, 0x69, 0x6b, 0x61, 0x7a, 0x75, 0x63, 0x68, 0x69, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x5f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_job_application_proto_rawDescOnce sync.Once
	file_job_application_proto_rawDescData []byte
)

func file_job_application_proto_rawDescGZIP() []byte {
	file_job_application_proto_rawDescOnce.Do(func() {
		file_job_application_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_job_application_proto_rawDesc), len(file_job_application_proto_rawDesc)))
	})
	return file_job_application_proto_rawDescData
}

var file_job_application_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_job_application_proto_goTypes = []any{
	(*FindJobApplicationByIdRequest)(nil), // 0: FindJobApplicationByIdRequest
	(*JobApplicationResponse)(nil),        // 1: JobApplicationResponse
	(*FindAllApplicationRequest)(nil),     // 2: FindAllApplicationRequest
	(*SelectApplicationRequest)(nil),      // 3: SelectApplicationRequest
	(*JobApplicationResponses)(nil),       // 4: JobApplicationResponses
	(*ApplyRequest)(nil),                  // 5: ApplyRequest
	(*user.UserJwtClaim)(nil),             // 6: UserJwtClaim
	(*emptypb.Empty)(nil),                 // 7: google.protobuf.Empty
}
var file_job_application_proto_depIdxs = []int32{
	6, // 0: FindAllApplicationRequest.UserJwtClaim:type_name -> UserJwtClaim
	6, // 1: SelectApplicationRequest.UserJwtClaim:type_name -> UserJwtClaim
	1, // 2: JobApplicationResponses.JobApplicationResponses:type_name -> JobApplicationResponse
	6, // 3: ApplyRequest.UserJwtClaim:type_name -> UserJwtClaim
	0, // 4: JobApplicationService.FindById:input_type -> FindJobApplicationByIdRequest
	2, // 5: JobApplicationService.FindAllApplication:input_type -> FindAllApplicationRequest
	3, // 6: JobApplicationService.SelectApplication:input_type -> SelectApplicationRequest
	5, // 7: JobApplicationService.HandleApply:input_type -> ApplyRequest
	1, // 8: JobApplicationService.FindById:output_type -> JobApplicationResponse
	4, // 9: JobApplicationService.FindAllApplication:output_type -> JobApplicationResponses
	7, // 10: JobApplicationService.SelectApplication:output_type -> google.protobuf.Empty
	7, // 11: JobApplicationService.HandleApply:output_type -> google.protobuf.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_job_application_proto_init() }
func file_job_application_proto_init() {
	if File_job_application_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_job_application_proto_rawDesc), len(file_job_application_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_application_proto_goTypes,
		DependencyIndexes: file_job_application_proto_depIdxs,
		MessageInfos:      file_job_application_proto_msgTypes,
	}.Build()
	File_job_application_proto = out.File
	file_job_application_proto_goTypes = nil
	file_job_application_proto_depIdxs = nil
}
