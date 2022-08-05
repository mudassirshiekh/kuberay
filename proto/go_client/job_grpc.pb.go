// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RayJobServiceClient is the client API for RayJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RayJobServiceClient interface {
	// Creates a new compute template.
	CreateRayJob(ctx context.Context, in *CreateRayJobRequest, opts ...grpc.CallOption) (*RayJob, error)
	// Finds a specific compute template by its name and namespace.
	GetRayJob(ctx context.Context, in *GetRayJobRequest, opts ...grpc.CallOption) (*RayJob, error)
	// Finds all compute templates in a given namespace. Supports pagination, and sorting on certain fields.
	ListRayJobs(ctx context.Context, in *ListRayJobsRequest, opts ...grpc.CallOption) (*ListRayJobsResponse, error)
	// Finds all compute templates in a given namespace. Supports pagination, and sorting on certain fields.
	ListAllRayJobs(ctx context.Context, in *ListAllRayJobsRequest, opts ...grpc.CallOption) (*ListAllRayJobsResponse, error)
	// Deletes a compute template by its name and namespace
	DeleteRayJob(ctx context.Context, in *DeleteRayJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type rayJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRayJobServiceClient(cc grpc.ClientConnInterface) RayJobServiceClient {
	return &rayJobServiceClient{cc}
}

func (c *rayJobServiceClient) CreateRayJob(ctx context.Context, in *CreateRayJobRequest, opts ...grpc.CallOption) (*RayJob, error) {
	out := new(RayJob)
	err := c.cc.Invoke(ctx, "/proto.RayJobService/CreateRayJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobServiceClient) GetRayJob(ctx context.Context, in *GetRayJobRequest, opts ...grpc.CallOption) (*RayJob, error) {
	out := new(RayJob)
	err := c.cc.Invoke(ctx, "/proto.RayJobService/GetRayJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobServiceClient) ListRayJobs(ctx context.Context, in *ListRayJobsRequest, opts ...grpc.CallOption) (*ListRayJobsResponse, error) {
	out := new(ListRayJobsResponse)
	err := c.cc.Invoke(ctx, "/proto.RayJobService/ListRayJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobServiceClient) ListAllRayJobs(ctx context.Context, in *ListAllRayJobsRequest, opts ...grpc.CallOption) (*ListAllRayJobsResponse, error) {
	out := new(ListAllRayJobsResponse)
	err := c.cc.Invoke(ctx, "/proto.RayJobService/ListAllRayJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobServiceClient) DeleteRayJob(ctx context.Context, in *DeleteRayJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.RayJobService/DeleteRayJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RayJobServiceServer is the server API for RayJobService service.
// All implementations must embed UnimplementedRayJobServiceServer
// for forward compatibility
type RayJobServiceServer interface {
	// Creates a new compute template.
	CreateRayJob(context.Context, *CreateRayJobRequest) (*RayJob, error)
	// Finds a specific compute template by its name and namespace.
	GetRayJob(context.Context, *GetRayJobRequest) (*RayJob, error)
	// Finds all compute templates in a given namespace. Supports pagination, and sorting on certain fields.
	ListRayJobs(context.Context, *ListRayJobsRequest) (*ListRayJobsResponse, error)
	// Finds all compute templates in a given namespace. Supports pagination, and sorting on certain fields.
	ListAllRayJobs(context.Context, *ListAllRayJobsRequest) (*ListAllRayJobsResponse, error)
	// Deletes a compute template by its name and namespace
	DeleteRayJob(context.Context, *DeleteRayJobRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRayJobServiceServer()
}

// UnimplementedRayJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRayJobServiceServer struct {
}

func (UnimplementedRayJobServiceServer) CreateRayJob(context.Context, *CreateRayJobRequest) (*RayJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRayJob not implemented")
}
func (UnimplementedRayJobServiceServer) GetRayJob(context.Context, *GetRayJobRequest) (*RayJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRayJob not implemented")
}
func (UnimplementedRayJobServiceServer) ListRayJobs(context.Context, *ListRayJobsRequest) (*ListRayJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRayJobs not implemented")
}
func (UnimplementedRayJobServiceServer) ListAllRayJobs(context.Context, *ListAllRayJobsRequest) (*ListAllRayJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllRayJobs not implemented")
}
func (UnimplementedRayJobServiceServer) DeleteRayJob(context.Context, *DeleteRayJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRayJob not implemented")
}
func (UnimplementedRayJobServiceServer) mustEmbedUnimplementedRayJobServiceServer() {}

// UnsafeRayJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RayJobServiceServer will
// result in compilation errors.
type UnsafeRayJobServiceServer interface {
	mustEmbedUnimplementedRayJobServiceServer()
}

func RegisterRayJobServiceServer(s grpc.ServiceRegistrar, srv RayJobServiceServer) {
	s.RegisterService(&RayJobService_ServiceDesc, srv)
}

func _RayJobService_CreateRayJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRayJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobServiceServer).CreateRayJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobService/CreateRayJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobServiceServer).CreateRayJob(ctx, req.(*CreateRayJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobService_GetRayJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRayJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobServiceServer).GetRayJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobService/GetRayJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobServiceServer).GetRayJob(ctx, req.(*GetRayJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobService_ListRayJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRayJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobServiceServer).ListRayJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobService/ListRayJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobServiceServer).ListRayJobs(ctx, req.(*ListRayJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobService_ListAllRayJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllRayJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobServiceServer).ListAllRayJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobService/ListAllRayJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobServiceServer).ListAllRayJobs(ctx, req.(*ListAllRayJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobService_DeleteRayJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRayJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobServiceServer).DeleteRayJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobService/DeleteRayJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobServiceServer).DeleteRayJob(ctx, req.(*DeleteRayJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RayJobService_ServiceDesc is the grpc.ServiceDesc for RayJobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RayJobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RayJobService",
	HandlerType: (*RayJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRayJob",
			Handler:    _RayJobService_CreateRayJob_Handler,
		},
		{
			MethodName: "GetRayJob",
			Handler:    _RayJobService_GetRayJob_Handler,
		},
		{
			MethodName: "ListRayJobs",
			Handler:    _RayJobService_ListRayJobs_Handler,
		},
		{
			MethodName: "ListAllRayJobs",
			Handler:    _RayJobService_ListAllRayJobs_Handler,
		},
		{
			MethodName: "DeleteRayJob",
			Handler:    _RayJobService_DeleteRayJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}