// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: internal/proto/student/student.proto

package studentPb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	// Any user can call this
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error)
	CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*CreateStudentResponse, error)
	// Only admin can call this to delete a student
	DeleteStudent(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*DeleteStudentResponse, error)
	// Only student of his id can call this
	UpdateStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error)
	IsProfileCompleted(ctx context.Context, in *IsProfileCompletedRequest, opts ...grpc.CallOption) (*IsProfileCompletedResponse, error)
	GetGPA(ctx context.Context, in *GPARequest, opts ...grpc.CallOption) (*GPAResponse, error)
	UpdateGPA(ctx context.Context, in *UpdateGPARequest, opts ...grpc.CallOption) (*UpdateGPAResponse, error)
	UpdateGPAWithImage(ctx context.Context, opts ...grpc.CallOption) (StudentService_UpdateGPAWithImageClient, error)
	AddGPA(ctx context.Context, opts ...grpc.CallOption) (StudentService_AddGPAClient, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error) {
	out := new(GetStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*CreateStudentResponse, error) {
	out := new(CreateStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/CreateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) DeleteStudent(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*DeleteStudentResponse, error) {
	out := new(DeleteStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/DeleteStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateStudent(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error) {
	out := new(UpdateStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/UpdateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) IsProfileCompleted(ctx context.Context, in *IsProfileCompletedRequest, opts ...grpc.CallOption) (*IsProfileCompletedResponse, error) {
	out := new(IsProfileCompletedResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/IsProfileCompleted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetGPA(ctx context.Context, in *GPARequest, opts ...grpc.CallOption) (*GPAResponse, error) {
	out := new(GPAResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/GetGPA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateGPA(ctx context.Context, in *UpdateGPARequest, opts ...grpc.CallOption) (*UpdateGPAResponse, error) {
	out := new(UpdateGPAResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/UpdateGPA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) UpdateGPAWithImage(ctx context.Context, opts ...grpc.CallOption) (StudentService_UpdateGPAWithImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &StudentService_ServiceDesc.Streams[0], "/student.StudentService/UpdateGPAWithImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentServiceUpdateGPAWithImageClient{stream}
	return x, nil
}

type StudentService_UpdateGPAWithImageClient interface {
	Send(*UpdateGPARequestWithImage) error
	CloseAndRecv() (*UpdateGPAResponseWithImage, error)
	grpc.ClientStream
}

type studentServiceUpdateGPAWithImageClient struct {
	grpc.ClientStream
}

func (x *studentServiceUpdateGPAWithImageClient) Send(m *UpdateGPARequestWithImage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *studentServiceUpdateGPAWithImageClient) CloseAndRecv() (*UpdateGPAResponseWithImage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateGPAResponseWithImage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *studentServiceClient) AddGPA(ctx context.Context, opts ...grpc.CallOption) (StudentService_AddGPAClient, error) {
	stream, err := c.cc.NewStream(ctx, &StudentService_ServiceDesc.Streams[1], "/student.StudentService/AddGPA", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentServiceAddGPAClient{stream}
	return x, nil
}

type StudentService_AddGPAClient interface {
	Send(*AddGPARequest) error
	CloseAndRecv() (*AddGPAResponse, error)
	grpc.ClientStream
}

type studentServiceAddGPAClient struct {
	grpc.ClientStream
}

func (x *studentServiceAddGPAClient) Send(m *AddGPARequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *studentServiceAddGPAClient) CloseAndRecv() (*AddGPAResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddGPAResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	// Any user can call this
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error)
	CreateStudent(context.Context, *CreateStudentRequest) (*CreateStudentResponse, error)
	// Only admin can call this to delete a student
	DeleteStudent(context.Context, *DeleteStudentRequest) (*DeleteStudentResponse, error)
	// Only student of his id can call this
	UpdateStudent(context.Context, *UpdateStudentRequest) (*UpdateStudentResponse, error)
	IsProfileCompleted(context.Context, *IsProfileCompletedRequest) (*IsProfileCompletedResponse, error)
	GetGPA(context.Context, *GPARequest) (*GPAResponse, error)
	UpdateGPA(context.Context, *UpdateGPARequest) (*UpdateGPAResponse, error)
	UpdateGPAWithImage(StudentService_UpdateGPAWithImageServer) error
	AddGPA(StudentService_AddGPAServer) error
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedStudentServiceServer) CreateStudent(context.Context, *CreateStudentRequest) (*CreateStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedStudentServiceServer) DeleteStudent(context.Context, *DeleteStudentRequest) (*DeleteStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedStudentServiceServer) UpdateStudent(context.Context, *UpdateStudentRequest) (*UpdateStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedStudentServiceServer) IsProfileCompleted(context.Context, *IsProfileCompletedRequest) (*IsProfileCompletedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsProfileCompleted not implemented")
}
func (UnimplementedStudentServiceServer) GetGPA(context.Context, *GPARequest) (*GPAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGPA not implemented")
}
func (UnimplementedStudentServiceServer) UpdateGPA(context.Context, *UpdateGPARequest) (*UpdateGPAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGPA not implemented")
}
func (UnimplementedStudentServiceServer) UpdateGPAWithImage(StudentService_UpdateGPAWithImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateGPAWithImage not implemented")
}
func (UnimplementedStudentServiceServer) AddGPA(StudentService_AddGPAServer) error {
	return status.Errorf(codes.Unimplemented, "method AddGPA not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/CreateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateStudent(ctx, req.(*CreateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/DeleteStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).DeleteStudent(ctx, req.(*DeleteStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/UpdateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateStudent(ctx, req.(*UpdateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_IsProfileCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsProfileCompletedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).IsProfileCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/IsProfileCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).IsProfileCompleted(ctx, req.(*IsProfileCompletedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetGPA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GPARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetGPA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/GetGPA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetGPA(ctx, req.(*GPARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateGPA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGPARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).UpdateGPA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/UpdateGPA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).UpdateGPA(ctx, req.(*UpdateGPARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_UpdateGPAWithImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StudentServiceServer).UpdateGPAWithImage(&studentServiceUpdateGPAWithImageServer{stream})
}

type StudentService_UpdateGPAWithImageServer interface {
	SendAndClose(*UpdateGPAResponseWithImage) error
	Recv() (*UpdateGPARequestWithImage, error)
	grpc.ServerStream
}

type studentServiceUpdateGPAWithImageServer struct {
	grpc.ServerStream
}

func (x *studentServiceUpdateGPAWithImageServer) SendAndClose(m *UpdateGPAResponseWithImage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *studentServiceUpdateGPAWithImageServer) Recv() (*UpdateGPARequestWithImage, error) {
	m := new(UpdateGPARequestWithImage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StudentService_AddGPA_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StudentServiceServer).AddGPA(&studentServiceAddGPAServer{stream})
}

type StudentService_AddGPAServer interface {
	SendAndClose(*AddGPAResponse) error
	Recv() (*AddGPARequest, error)
	grpc.ServerStream
}

type studentServiceAddGPAServer struct {
	grpc.ServerStream
}

func (x *studentServiceAddGPAServer) SendAndClose(m *AddGPAResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *studentServiceAddGPAServer) Recv() (*AddGPARequest, error) {
	m := new(AddGPARequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudent",
			Handler:    _StudentService_GetStudent_Handler,
		},
		{
			MethodName: "CreateStudent",
			Handler:    _StudentService_CreateStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _StudentService_DeleteStudent_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _StudentService_UpdateStudent_Handler,
		},
		{
			MethodName: "IsProfileCompleted",
			Handler:    _StudentService_IsProfileCompleted_Handler,
		},
		{
			MethodName: "GetGPA",
			Handler:    _StudentService_GetGPA_Handler,
		},
		{
			MethodName: "UpdateGPA",
			Handler:    _StudentService_UpdateGPA_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateGPAWithImage",
			Handler:       _StudentService_UpdateGPAWithImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddGPA",
			Handler:       _StudentService_AddGPA_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/student/student.proto",
}
