// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: internal/proto/certificates/certificate.proto

package certificatePb

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

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateServiceClient interface {
	// Add a new certificate, then it will be uploaded to s3 and url will be stored in db
	AddCertificate(ctx context.Context, opts ...grpc.CallOption) (CertificateService_AddCertificateClient, error)
	// Get all certificate of a user
	GetAllCertificate(ctx context.Context, in *GetAllCertificateRequest, opts ...grpc.CallOption) (*GetAllCertificateResponse, error)
	GetCertificate(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*GetCertificateResponse, error)
	UpdateCertificate(ctx context.Context, in *UpdateCertificateRequest, opts ...grpc.CallOption) (*UpdateCertificateResponse, error)
	DeleteCertificate(ctx context.Context, in *DeleteCertificateRequest, opts ...grpc.CallOption) (*DeleteCertificateResponse, error)
	ChangeCertificateStatus(ctx context.Context, in *ChangeCertificateStatusRequest, opts ...grpc.CallOption) (*ChangeCertificateStatusResponse, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) AddCertificate(ctx context.Context, opts ...grpc.CallOption) (CertificateService_AddCertificateClient, error) {
	stream, err := c.cc.NewStream(ctx, &CertificateService_ServiceDesc.Streams[0], "/certificate.CertificateService/AddCertificate", opts...)
	if err != nil {
		return nil, err
	}
	x := &certificateServiceAddCertificateClient{stream}
	return x, nil
}

type CertificateService_AddCertificateClient interface {
	Send(*AddCertificateRequest) error
	CloseAndRecv() (*AddCertificateResponse, error)
	grpc.ClientStream
}

type certificateServiceAddCertificateClient struct {
	grpc.ClientStream
}

func (x *certificateServiceAddCertificateClient) Send(m *AddCertificateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *certificateServiceAddCertificateClient) CloseAndRecv() (*AddCertificateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddCertificateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *certificateServiceClient) GetAllCertificate(ctx context.Context, in *GetAllCertificateRequest, opts ...grpc.CallOption) (*GetAllCertificateResponse, error) {
	out := new(GetAllCertificateResponse)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/GetAllCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) GetCertificate(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*GetCertificateResponse, error) {
	out := new(GetCertificateResponse)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/GetCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) UpdateCertificate(ctx context.Context, in *UpdateCertificateRequest, opts ...grpc.CallOption) (*UpdateCertificateResponse, error) {
	out := new(UpdateCertificateResponse)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/UpdateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) DeleteCertificate(ctx context.Context, in *DeleteCertificateRequest, opts ...grpc.CallOption) (*DeleteCertificateResponse, error) {
	out := new(DeleteCertificateResponse)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/DeleteCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) ChangeCertificateStatus(ctx context.Context, in *ChangeCertificateStatusRequest, opts ...grpc.CallOption) (*ChangeCertificateStatusResponse, error) {
	out := new(ChangeCertificateStatusResponse)
	err := c.cc.Invoke(ctx, "/certificate.CertificateService/ChangeCertificateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations must embed UnimplementedCertificateServiceServer
// for forward compatibility
type CertificateServiceServer interface {
	// Add a new certificate, then it will be uploaded to s3 and url will be stored in db
	AddCertificate(CertificateService_AddCertificateServer) error
	// Get all certificate of a user
	GetAllCertificate(context.Context, *GetAllCertificateRequest) (*GetAllCertificateResponse, error)
	GetCertificate(context.Context, *GetCertificateRequest) (*GetCertificateResponse, error)
	UpdateCertificate(context.Context, *UpdateCertificateRequest) (*UpdateCertificateResponse, error)
	DeleteCertificate(context.Context, *DeleteCertificateRequest) (*DeleteCertificateResponse, error)
	ChangeCertificateStatus(context.Context, *ChangeCertificateStatusRequest) (*ChangeCertificateStatusResponse, error)
	mustEmbedUnimplementedCertificateServiceServer()
}

// UnimplementedCertificateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (UnimplementedCertificateServiceServer) AddCertificate(CertificateService_AddCertificateServer) error {
	return status.Errorf(codes.Unimplemented, "method AddCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) GetAllCertificate(context.Context, *GetAllCertificateRequest) (*GetAllCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) GetCertificate(context.Context, *GetCertificateRequest) (*GetCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) UpdateCertificate(context.Context, *UpdateCertificateRequest) (*UpdateCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) DeleteCertificate(context.Context, *DeleteCertificateRequest) (*DeleteCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) ChangeCertificateStatus(context.Context, *ChangeCertificateStatusRequest) (*ChangeCertificateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCertificateStatus not implemented")
}
func (UnimplementedCertificateServiceServer) mustEmbedUnimplementedCertificateServiceServer() {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_AddCertificate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CertificateServiceServer).AddCertificate(&certificateServiceAddCertificateServer{stream})
}

type CertificateService_AddCertificateServer interface {
	SendAndClose(*AddCertificateResponse) error
	Recv() (*AddCertificateRequest, error)
	grpc.ServerStream
}

type certificateServiceAddCertificateServer struct {
	grpc.ServerStream
}

func (x *certificateServiceAddCertificateServer) SendAndClose(m *AddCertificateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *certificateServiceAddCertificateServer) Recv() (*AddCertificateRequest, error) {
	m := new(AddCertificateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CertificateService_GetAllCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetAllCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/GetAllCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetAllCertificate(ctx, req.(*GetAllCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_GetCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/GetCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetCertificate(ctx, req.(*GetCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_UpdateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).UpdateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/UpdateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).UpdateCertificate(ctx, req.(*UpdateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_DeleteCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).DeleteCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/DeleteCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).DeleteCertificate(ctx, req.(*DeleteCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_ChangeCertificateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCertificateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).ChangeCertificateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.CertificateService/ChangeCertificateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).ChangeCertificateStatus(ctx, req.(*ChangeCertificateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCertificate",
			Handler:    _CertificateService_GetAllCertificate_Handler,
		},
		{
			MethodName: "GetCertificate",
			Handler:    _CertificateService_GetCertificate_Handler,
		},
		{
			MethodName: "UpdateCertificate",
			Handler:    _CertificateService_UpdateCertificate_Handler,
		},
		{
			MethodName: "DeleteCertificate",
			Handler:    _CertificateService_DeleteCertificate_Handler,
		},
		{
			MethodName: "ChangeCertificateStatus",
			Handler:    _CertificateService_ChangeCertificateStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddCertificate",
			Handler:       _CertificateService_AddCertificate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/certificates/certificate.proto",
}
