// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: internal/proto/student/student.proto

package studentPb

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

type Student_PhoneType int32

const (
	Student_MOBILE Student_PhoneType = 0
	Student_HOME   Student_PhoneType = 1
)

// Enum value maps for Student_PhoneType.
var (
	Student_PhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
	}
	Student_PhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
	}
)

func (x Student_PhoneType) Enum() *Student_PhoneType {
	p := new(Student_PhoneType)
	*p = x
	return p
}

func (x Student_PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Student_PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_student_student_proto_enumTypes[0].Descriptor()
}

func (Student_PhoneType) Type() protoreflect.EnumType {
	return &file_internal_proto_student_student_proto_enumTypes[0]
}

func (x Student_PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Student_PhoneType.Descriptor instead.
func (Student_PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{0, 0}
}

// Assuming there are only 4 types of AcademicStatus
// Currently assuming that only one cycle of admission is possible
//
// There won't be a case where 2 students have the same academic status
// But different semesters
type StudentAcademicStatus int32

const (
	Student_FRESHMAN  StudentAcademicStatus = 0
	Student_SOPHOMORE StudentAcademicStatus = 1
	Student_JUNIOR    StudentAcademicStatus = 2
	Student_SENIOR    StudentAcademicStatus = 3
)

// Enum value maps for StudentAcademicStatus.
var (
	StudentAcademicStatus_name = map[int32]string{
		0: "FRESHMAN",
		1: "SOPHOMORE",
		2: "JUNIOR",
		3: "SENIOR",
	}
	StudentAcademicStatus_value = map[string]int32{
		"FRESHMAN":  0,
		"SOPHOMORE": 1,
		"JUNIOR":    2,
		"SENIOR":    3,
	}
)

func (x StudentAcademicStatus) Enum() *StudentAcademicStatus {
	p := new(StudentAcademicStatus)
	*p = x
	return p
}

func (x StudentAcademicStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StudentAcademicStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_student_student_proto_enumTypes[1].Descriptor()
}

func (StudentAcademicStatus) Type() protoreflect.EnumType {
	return &file_internal_proto_student_student_proto_enumTypes[1]
}

func (x StudentAcademicStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StudentAcademicStatus.Descriptor instead.
func (StudentAcademicStatus) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{0, 1}
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName    string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	MiddleName   string                 `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	Email        string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phones       []*Student_PhoneNumber `protobuf:"bytes,6,rep,name=phones,proto3" json:"phones,omitempty"`
	Status       *StudentAcademicStatus `protobuf:"varint,7,opt,name=status,proto3,enum=student.StudentAcademicStatus,oneof" json:"status,omitempty"`
	FieldOfStudy *string                `protobuf:"bytes,8,opt,name=field_of_study,json=fieldOfStudy,proto3,oneof" json:"field_of_study,omitempty"` // e.g. Computer Science
	IsBanned     *bool                  `protobuf:"varint,9,opt,name=is_banned,json=isBanned,proto3,oneof" json:"is_banned,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Student) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Student) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Student) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Student) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Student) GetPhones() []*Student_PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *Student) GetStatus() StudentAcademicStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Student_FRESHMAN
}

func (x *Student) GetFieldOfStudy() string {
	if x != nil && x.FieldOfStudy != nil {
		return *x.FieldOfStudy
	}
	return ""
}

func (x *Student) GetIsBanned() bool {
	if x != nil && x.IsBanned != nil {
		return *x.IsBanned
	}
	return false
}

type UpdateStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Don't update the id
	Student *Student `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
}

func (x *UpdateStudentRequest) Reset() {
	*x = UpdateStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentRequest) ProtoMessage() {}

func (x *UpdateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentRequest.ProtoReflect.Descriptor instead.
func (*UpdateStudentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStudentRequest) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type UpdateStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateStudentResponse) Reset() {
	*x = UpdateStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentResponse) ProtoMessage() {}

func (x *UpdateStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentResponse.ProtoReflect.Descriptor instead.
func (*UpdateStudentResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateStudentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateStudentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Student *Student `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
}

func (x *CreateStudentRequest) Reset() {
	*x = CreateStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentRequest) ProtoMessage() {}

func (x *CreateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStudentRequest) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type CreateStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateStudentResponse) Reset() {
	*x = CreateStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentResponse) ProtoMessage() {}

func (x *CreateStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentResponse.ProtoReflect.Descriptor instead.
func (*CreateStudentResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{4}
}

func (x *CreateStudentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateStudentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStudentRequest) Reset() {
	*x = DeleteStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentRequest) ProtoMessage() {}

func (x *DeleteStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentRequest.ProtoReflect.Descriptor instead.
func (*DeleteStudentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteStudentResponse) Reset() {
	*x = DeleteStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentResponse) ProtoMessage() {}

func (x *DeleteStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentResponse.ProtoReflect.Descriptor instead.
func (*DeleteStudentResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteStudentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteStudentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Any user can execute this
type GetStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStudentRequest) Reset() {
	*x = GetStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentRequest) ProtoMessage() {}

func (x *GetStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentRequest.ProtoReflect.Descriptor instead.
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{7}
}

func (x *GetStudentRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Student []*Student `protobuf:"bytes,1,rep,name=student,proto3" json:"student,omitempty"`
}

func (x *GetStudentResponse) Reset() {
	*x = GetStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentResponse) ProtoMessage() {}

func (x *GetStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentResponse.ProtoReflect.Descriptor instead.
func (*GetStudentResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{8}
}

func (x *GetStudentResponse) GetStudent() []*Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type Student_PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string            `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   Student_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=student.Student_PhoneType" json:"type,omitempty"`
}

func (x *Student_PhoneNumber) Reset() {
	*x = Student_PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_student_student_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student_PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student_PhoneNumber) ProtoMessage() {}

func (x *Student_PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_student_student_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student_PhoneNumber.ProtoReflect.Descriptor instead.
func (*Student_PhoneNumber) Descriptor() ([]byte, []int) {
	return file_internal_proto_student_student_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Student_PhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Student_PhoneNumber) GetType() Student_PhoneType {
	if x != nil {
		return x.Type
	}
	return Student_MOBILE
}

var File_internal_proto_student_student_proto protoreflect.FileDescriptor

var file_internal_proto_student_student_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22,
	0xba, 0x04, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x34,
	0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x66, 0x53, 0x74, 0x75, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x02, 0x52, 0x08, 0x69, 0x73, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x88, 0x01, 0x01, 0x1a,
	0x55, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x01, 0x22, 0x45, 0x0a, 0x0e, 0x61, 0x63, 0x61,
	0x64, 0x65, 0x6d, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x46,
	0x52, 0x45, 0x53, 0x48, 0x4d, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4f, 0x50,
	0x48, 0x4f, 0x4d, 0x4f, 0x52, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x55, 0x4e, 0x49,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x4e, 0x49, 0x4f, 0x52, 0x10, 0x03,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x22, 0x42, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x22, 0x4b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x42, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x22, 0x4b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x32, 0xc7, 0x02, 0x0a, 0x0e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_proto_student_student_proto_rawDescOnce sync.Once
	file_internal_proto_student_student_proto_rawDescData = file_internal_proto_student_student_proto_rawDesc
)

func file_internal_proto_student_student_proto_rawDescGZIP() []byte {
	file_internal_proto_student_student_proto_rawDescOnce.Do(func() {
		file_internal_proto_student_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_student_student_proto_rawDescData)
	})
	return file_internal_proto_student_student_proto_rawDescData
}

var file_internal_proto_student_student_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_proto_student_student_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_proto_student_student_proto_goTypes = []interface{}{
	(Student_PhoneType)(0),        // 0: student.Student.PhoneType
	(StudentAcademicStatus)(0),    // 1: student.Student.academicStatus
	(*Student)(nil),               // 2: student.Student
	(*UpdateStudentRequest)(nil),  // 3: student.UpdateStudentRequest
	(*UpdateStudentResponse)(nil), // 4: student.UpdateStudentResponse
	(*CreateStudentRequest)(nil),  // 5: student.CreateStudentRequest
	(*CreateStudentResponse)(nil), // 6: student.CreateStudentResponse
	(*DeleteStudentRequest)(nil),  // 7: student.DeleteStudentRequest
	(*DeleteStudentResponse)(nil), // 8: student.DeleteStudentResponse
	(*GetStudentRequest)(nil),     // 9: student.GetStudentRequest
	(*GetStudentResponse)(nil),    // 10: student.GetStudentResponse
	(*Student_PhoneNumber)(nil),   // 11: student.Student.PhoneNumber
}
var file_internal_proto_student_student_proto_depIdxs = []int32{
	11, // 0: student.Student.phones:type_name -> student.Student.PhoneNumber
	1,  // 1: student.Student.status:type_name -> student.Student.academicStatus
	2,  // 2: student.UpdateStudentRequest.student:type_name -> student.Student
	2,  // 3: student.CreateStudentRequest.student:type_name -> student.Student
	2,  // 4: student.GetStudentResponse.student:type_name -> student.Student
	0,  // 5: student.Student.PhoneNumber.type:type_name -> student.Student.PhoneType
	9,  // 6: student.StudentService.GetStudent:input_type -> student.GetStudentRequest
	5,  // 7: student.StudentService.CreateStudent:input_type -> student.CreateStudentRequest
	7,  // 8: student.StudentService.DeleteStudent:input_type -> student.DeleteStudentRequest
	3,  // 9: student.StudentService.UpdateStudent:input_type -> student.UpdateStudentRequest
	10, // 10: student.StudentService.GetStudent:output_type -> student.GetStudentResponse
	6,  // 11: student.StudentService.CreateStudent:output_type -> student.CreateStudentResponse
	8,  // 12: student.StudentService.DeleteStudent:output_type -> student.DeleteStudentResponse
	4,  // 13: student.StudentService.UpdateStudent:output_type -> student.UpdateStudentResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_internal_proto_student_student_proto_init() }
func file_internal_proto_student_student_proto_init() {
	if File_internal_proto_student_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_student_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_internal_proto_student_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentRequest); i {
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
		file_internal_proto_student_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStudentResponse); i {
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
		file_internal_proto_student_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentRequest); i {
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
		file_internal_proto_student_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStudentResponse); i {
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
		file_internal_proto_student_student_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStudentRequest); i {
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
		file_internal_proto_student_student_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStudentResponse); i {
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
		file_internal_proto_student_student_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentRequest); i {
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
		file_internal_proto_student_student_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentResponse); i {
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
		file_internal_proto_student_student_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student_PhoneNumber); i {
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
	file_internal_proto_student_student_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_student_student_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_student_student_proto_goTypes,
		DependencyIndexes: file_internal_proto_student_student_proto_depIdxs,
		EnumInfos:         file_internal_proto_student_student_proto_enumTypes,
		MessageInfos:      file_internal_proto_student_student_proto_msgTypes,
	}.Build()
	File_internal_proto_student_student_proto = out.File
	file_internal_proto_student_student_proto_rawDesc = nil
	file_internal_proto_student_student_proto_goTypes = nil
	file_internal_proto_student_student_proto_depIdxs = nil
}
