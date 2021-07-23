// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ddbv1

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

// DDBAPIClient is the client API for DDBAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DDBAPIClient interface {
	DescribeTable(ctx context.Context, in *DescribeTableRequest, opts ...grpc.CallOption) (*DescribeTableResponse, error)
	UpdateTableCapacity(ctx context.Context, in *UpdateTableCapacityRequest, opts ...grpc.CallOption) (*UpdateTableCapacityResponse, error)
	UpdateGSICapacity(ctx context.Context, in *UpdateGSICapacityRequest, opts ...grpc.CallOption) (*UpdateGSICapacityResponse, error)
}

type dDBAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDDBAPIClient(cc grpc.ClientConnInterface) DDBAPIClient {
	return &dDBAPIClient{cc}
}

func (c *dDBAPIClient) DescribeTable(ctx context.Context, in *DescribeTableRequest, opts ...grpc.CallOption) (*DescribeTableResponse, error) {
	out := new(DescribeTableResponse)
	err := c.cc.Invoke(ctx, "/clutch.aws.dynamodb.v1.DDBAPI/DescribeTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDBAPIClient) UpdateTableCapacity(ctx context.Context, in *UpdateTableCapacityRequest, opts ...grpc.CallOption) (*UpdateTableCapacityResponse, error) {
	out := new(UpdateTableCapacityResponse)
	err := c.cc.Invoke(ctx, "/clutch.aws.dynamodb.v1.DDBAPI/UpdateTableCapacity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDBAPIClient) UpdateGSICapacity(ctx context.Context, in *UpdateGSICapacityRequest, opts ...grpc.CallOption) (*UpdateGSICapacityResponse, error) {
	out := new(UpdateGSICapacityResponse)
	err := c.cc.Invoke(ctx, "/clutch.aws.dynamodb.v1.DDBAPI/UpdateGSICapacity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DDBAPIServer is the server API for DDBAPI service.
// All implementations should embed UnimplementedDDBAPIServer
// for forward compatibility
type DDBAPIServer interface {
	DescribeTable(context.Context, *DescribeTableRequest) (*DescribeTableResponse, error)
	UpdateTableCapacity(context.Context, *UpdateTableCapacityRequest) (*UpdateTableCapacityResponse, error)
	UpdateGSICapacity(context.Context, *UpdateGSICapacityRequest) (*UpdateGSICapacityResponse, error)
}

// UnimplementedDDBAPIServer should be embedded to have forward compatible implementations.
type UnimplementedDDBAPIServer struct {
}

func (UnimplementedDDBAPIServer) DescribeTable(context.Context, *DescribeTableRequest) (*DescribeTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTable not implemented")
}
func (UnimplementedDDBAPIServer) UpdateTableCapacity(context.Context, *UpdateTableCapacityRequest) (*UpdateTableCapacityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTableCapacity not implemented")
}
func (UnimplementedDDBAPIServer) UpdateGSICapacity(context.Context, *UpdateGSICapacityRequest) (*UpdateGSICapacityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGSICapacity not implemented")
}

// UnsafeDDBAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DDBAPIServer will
// result in compilation errors.
type UnsafeDDBAPIServer interface {
	mustEmbedUnimplementedDDBAPIServer()
}

func RegisterDDBAPIServer(s grpc.ServiceRegistrar, srv DDBAPIServer) {
	s.RegisterService(&DDBAPI_ServiceDesc, srv)
}

func _DDBAPI_DescribeTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDBAPIServer).DescribeTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.aws.dynamodb.v1.DDBAPI/DescribeTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDBAPIServer).DescribeTable(ctx, req.(*DescribeTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDBAPI_UpdateTableCapacity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTableCapacityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDBAPIServer).UpdateTableCapacity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.aws.dynamodb.v1.DDBAPI/UpdateTableCapacity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDBAPIServer).UpdateTableCapacity(ctx, req.(*UpdateTableCapacityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDBAPI_UpdateGSICapacity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGSICapacityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDBAPIServer).UpdateGSICapacity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.aws.dynamodb.v1.DDBAPI/UpdateGSICapacity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDBAPIServer).UpdateGSICapacity(ctx, req.(*UpdateGSICapacityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DDBAPI_ServiceDesc is the grpc.ServiceDesc for DDBAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DDBAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.aws.dynamodb.v1.DDBAPI",
	HandlerType: (*DDBAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeTable",
			Handler:    _DDBAPI_DescribeTable_Handler,
		},
		{
			MethodName: "UpdateTableCapacity",
			Handler:    _DDBAPI_UpdateTableCapacity_Handler,
		},
		{
			MethodName: "UpdateGSICapacity",
			Handler:    _DDBAPI_UpdateGSICapacity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aws/dynamodb/v1/dynamodb.proto",
}