// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: internal/proto/resume/resume.proto

package resumePb

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

type STATUS int32

const (
	STATUS_STATUS_UNSPECIFIED STATUS = 0
	STATUS_STATUS_PENDING     STATUS = 1
	STATUS_STATUS_APPROVED    STATUS = 2
	STATUS_STATUS_REJECTED    STATUS = 3
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_PENDING",
		2: "STATUS_APPROVED",
		3: "STATUS_REJECTED",
	}
	STATUS_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_PENDING":     1,
		"STATUS_APPROVED":    2,
		"STATUS_REJECTED":    3,
	}
)

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}

func (x STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_resume_resume_proto_enumTypes[0].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_internal_proto_resume_resume_proto_enumTypes[0]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_resume_resume_proto_rawDescGZIP(), []int{0}
}

type Resume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Resume) Reset() {
	*x = Resume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_resume_resume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resume) ProtoMessage() {}

func (x *Resume) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_resume_resume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resume.ProtoReflect.Descriptor instead.
func (*Resume) Descriptor() ([]byte, []int) {
	return file_internal_proto_resume_resume_proto_rawDescGZIP(), []int{0}
}

func (x *Resume) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResumeUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resume *Resume `protobuf:"bytes,1,opt,name=resume,proto3" json:"resume,omitempty"`
}

func (x *ResumeUploadRequest) Reset() {
	*x = ResumeUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_resume_resume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeUploadRequest) ProtoMessage() {}

func (x *ResumeUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_resume_resume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeUploadRequest.ProtoReflect.Descriptor instead.
func (*ResumeUploadRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_resume_resume_proto_rawDescGZIP(), []int{1}
}

func (x *ResumeUploadRequest) GetResume() *Resume {
	if x != nil {
		return x.Resume
	}
	return nil
}

type ResumeUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status STATUS `protobuf:"varint,1,opt,name=status,proto3,enum=resume.STATUS" json:"status,omitempty"`
}

func (x *ResumeUploadResponse) Reset() {
	*x = ResumeUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_resume_resume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeUploadResponse) ProtoMessage() {}

func (x *ResumeUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_resume_resume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeUploadResponse.ProtoReflect.Descriptor instead.
func (*ResumeUploadResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_resume_resume_proto_rawDescGZIP(), []int{2}
}

func (x *ResumeUploadResponse) GetStatus() STATUS {
	if x != nil {
		return x.Status
	}
	return STATUS_STATUS_UNSPECIFIED
}

var File_internal_proto_resume_resume_proto protoreflect.FileDescriptor

var file_internal_proto_resume_resume_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x13, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x5e, 0x0a, 0x06, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x32, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_resume_resume_proto_rawDescOnce sync.Once
	file_internal_proto_resume_resume_proto_rawDescData = file_internal_proto_resume_resume_proto_rawDesc
)

func file_internal_proto_resume_resume_proto_rawDescGZIP() []byte {
	file_internal_proto_resume_resume_proto_rawDescOnce.Do(func() {
		file_internal_proto_resume_resume_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_resume_resume_proto_rawDescData)
	})
	return file_internal_proto_resume_resume_proto_rawDescData
}

var file_internal_proto_resume_resume_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_proto_resume_resume_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_proto_resume_resume_proto_goTypes = []interface{}{
	(STATUS)(0),                  // 0: resume.STATUS
	(*Resume)(nil),               // 1: resume.Resume
	(*ResumeUploadRequest)(nil),  // 2: resume.ResumeUploadRequest
	(*ResumeUploadResponse)(nil), // 3: resume.ResumeUploadResponse
}
var file_internal_proto_resume_resume_proto_depIdxs = []int32{
	1, // 0: resume.ResumeUploadRequest.resume:type_name -> resume.Resume
	0, // 1: resume.ResumeUploadResponse.status:type_name -> resume.STATUS
	2, // 2: resume.ResumeService.UploadResume:input_type -> resume.ResumeUploadRequest
	3, // 3: resume.ResumeService.UploadResume:output_type -> resume.ResumeUploadResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_resume_resume_proto_init() }
func file_internal_proto_resume_resume_proto_init() {
	if File_internal_proto_resume_resume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_resume_resume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resume); i {
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
		file_internal_proto_resume_resume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeUploadRequest); i {
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
		file_internal_proto_resume_resume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeUploadResponse); i {
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
			RawDescriptor: file_internal_proto_resume_resume_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_resume_resume_proto_goTypes,
		DependencyIndexes: file_internal_proto_resume_resume_proto_depIdxs,
		EnumInfos:         file_internal_proto_resume_resume_proto_enumTypes,
		MessageInfos:      file_internal_proto_resume_resume_proto_msgTypes,
	}.Build()
	File_internal_proto_resume_resume_proto = out.File
	file_internal_proto_resume_resume_proto_rawDesc = nil
	file_internal_proto_resume_resume_proto_goTypes = nil
	file_internal_proto_resume_resume_proto_depIdxs = nil
}