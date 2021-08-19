// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_quiz_api

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

// OcpQuizApiServiceClient is the client API for OcpQuizApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpQuizApiServiceClient interface {
	// Create new Quiz
	CreateQuiz(ctx context.Context, in *CreateQuizV1Request, opts ...grpc.CallOption) (*CreateQuizV1Response, error)
	// Describe Quiz
	DescribeQuiz(ctx context.Context, in *DescribeQuizV1Request, opts ...grpc.CallOption) (*DescribeQuizV1Response, error)
	// Get list of Quizes
	ListQuiz(ctx context.Context, in *ListQuizV1Request, opts ...grpc.CallOption) (*ListQuizV1Response, error)
	// Remove Quiz
	RemoveQuiz(ctx context.Context, in *RemoveQuizV1Request, opts ...grpc.CallOption) (*RemoveQuizV1Response, error)
}

type ocpQuizApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpQuizApiServiceClient(cc grpc.ClientConnInterface) OcpQuizApiServiceClient {
	return &ocpQuizApiServiceClient{cc}
}

func (c *ocpQuizApiServiceClient) CreateQuiz(ctx context.Context, in *CreateQuizV1Request, opts ...grpc.CallOption) (*CreateQuizV1Response, error) {
	out := new(CreateQuizV1Response)
	err := c.cc.Invoke(ctx, "/ocp.quiz.api.OcpQuizApiService/CreateQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpQuizApiServiceClient) DescribeQuiz(ctx context.Context, in *DescribeQuizV1Request, opts ...grpc.CallOption) (*DescribeQuizV1Response, error) {
	out := new(DescribeQuizV1Response)
	err := c.cc.Invoke(ctx, "/ocp.quiz.api.OcpQuizApiService/DescribeQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpQuizApiServiceClient) ListQuiz(ctx context.Context, in *ListQuizV1Request, opts ...grpc.CallOption) (*ListQuizV1Response, error) {
	out := new(ListQuizV1Response)
	err := c.cc.Invoke(ctx, "/ocp.quiz.api.OcpQuizApiService/ListQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpQuizApiServiceClient) RemoveQuiz(ctx context.Context, in *RemoveQuizV1Request, opts ...grpc.CallOption) (*RemoveQuizV1Response, error) {
	out := new(RemoveQuizV1Response)
	err := c.cc.Invoke(ctx, "/ocp.quiz.api.OcpQuizApiService/RemoveQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpQuizApiServiceServer is the server API for OcpQuizApiService service.
// All implementations must embed UnimplementedOcpQuizApiServiceServer
// for forward compatibility
type OcpQuizApiServiceServer interface {
	// Create new Quiz
	CreateQuiz(context.Context, *CreateQuizV1Request) (*CreateQuizV1Response, error)
	// Describe Quiz
	DescribeQuiz(context.Context, *DescribeQuizV1Request) (*DescribeQuizV1Response, error)
	// Get list of Quizes
	ListQuiz(context.Context, *ListQuizV1Request) (*ListQuizV1Response, error)
	// Remove Quiz
	RemoveQuiz(context.Context, *RemoveQuizV1Request) (*RemoveQuizV1Response, error)
	mustEmbedUnimplementedOcpQuizApiServiceServer()
}

// UnimplementedOcpQuizApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOcpQuizApiServiceServer struct {
}

func (UnimplementedOcpQuizApiServiceServer) CreateQuiz(context.Context, *CreateQuizV1Request) (*CreateQuizV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuiz not implemented")
}
func (UnimplementedOcpQuizApiServiceServer) DescribeQuiz(context.Context, *DescribeQuizV1Request) (*DescribeQuizV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeQuiz not implemented")
}
func (UnimplementedOcpQuizApiServiceServer) ListQuiz(context.Context, *ListQuizV1Request) (*ListQuizV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuiz not implemented")
}
func (UnimplementedOcpQuizApiServiceServer) RemoveQuiz(context.Context, *RemoveQuizV1Request) (*RemoveQuizV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveQuiz not implemented")
}
func (UnimplementedOcpQuizApiServiceServer) mustEmbedUnimplementedOcpQuizApiServiceServer() {}

// UnsafeOcpQuizApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpQuizApiServiceServer will
// result in compilation errors.
type UnsafeOcpQuizApiServiceServer interface {
	mustEmbedUnimplementedOcpQuizApiServiceServer()
}

func RegisterOcpQuizApiServiceServer(s grpc.ServiceRegistrar, srv OcpQuizApiServiceServer) {
	s.RegisterService(&OcpQuizApiService_ServiceDesc, srv)
}

func _OcpQuizApiService_CreateQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuizV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpQuizApiServiceServer).CreateQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.quiz.api.OcpQuizApiService/CreateQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpQuizApiServiceServer).CreateQuiz(ctx, req.(*CreateQuizV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpQuizApiService_DescribeQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeQuizV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpQuizApiServiceServer).DescribeQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.quiz.api.OcpQuizApiService/DescribeQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpQuizApiServiceServer).DescribeQuiz(ctx, req.(*DescribeQuizV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpQuizApiService_ListQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuizV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpQuizApiServiceServer).ListQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.quiz.api.OcpQuizApiService/ListQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpQuizApiServiceServer).ListQuiz(ctx, req.(*ListQuizV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpQuizApiService_RemoveQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveQuizV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpQuizApiServiceServer).RemoveQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.quiz.api.OcpQuizApiService/RemoveQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpQuizApiServiceServer).RemoveQuiz(ctx, req.(*RemoveQuizV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpQuizApiService_ServiceDesc is the grpc.ServiceDesc for OcpQuizApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpQuizApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.quiz.api.OcpQuizApiService",
	HandlerType: (*OcpQuizApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuiz",
			Handler:    _OcpQuizApiService_CreateQuiz_Handler,
		},
		{
			MethodName: "DescribeQuiz",
			Handler:    _OcpQuizApiService_DescribeQuiz_Handler,
		},
		{
			MethodName: "ListQuiz",
			Handler:    _OcpQuizApiService_ListQuiz_Handler,
		},
		{
			MethodName: "RemoveQuiz",
			Handler:    _OcpQuizApiService_RemoveQuiz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-quiz-api/ocp_quiz_api.proto",
}
