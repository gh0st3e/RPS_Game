// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: result.proto

package proto

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

// ComputerClient is the client API for Computer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComputerClient interface {
	CallComputer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseComputer, error)
}

type computerClient struct {
	cc grpc.ClientConnInterface
}

func NewComputerClient(cc grpc.ClientConnInterface) ComputerClient {
	return &computerClient{cc}
}

func (c *computerClient) CallComputer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseComputer, error) {
	out := new(ResponseComputer)
	err := c.cc.Invoke(ctx, "/proto.Computer/CallComputer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComputerServer is the server API for Computer service.
// All implementations must embed UnimplementedComputerServer
// for forward compatibility
type ComputerServer interface {
	CallComputer(context.Context, *Empty) (*ResponseComputer, error)
	mustEmbedUnimplementedComputerServer()
}

// UnimplementedComputerServer must be embedded to have forward compatible implementations.
type UnimplementedComputerServer struct {
}

func (UnimplementedComputerServer) CallComputer(context.Context, *Empty) (*ResponseComputer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallComputer not implemented")
}
func (UnimplementedComputerServer) mustEmbedUnimplementedComputerServer() {}

// UnsafeComputerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComputerServer will
// result in compilation errors.
type UnsafeComputerServer interface {
	mustEmbedUnimplementedComputerServer()
}

func RegisterComputerServer(s grpc.ServiceRegistrar, srv ComputerServer) {
	s.RegisterService(&Computer_ServiceDesc, srv)
}

func _Computer_CallComputer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputerServer).CallComputer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Computer/CallComputer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputerServer).CallComputer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Computer_ServiceDesc is the grpc.ServiceDesc for Computer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Computer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Computer",
	HandlerType: (*ComputerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallComputer",
			Handler:    _Computer_CallComputer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "result.proto",
}

// ResultClient is the client API for Result service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResultClient interface {
	CallResult(ctx context.Context, in *RequestResult, opts ...grpc.CallOption) (*ResponseResult, error)
}

type resultClient struct {
	cc grpc.ClientConnInterface
}

func NewResultClient(cc grpc.ClientConnInterface) ResultClient {
	return &resultClient{cc}
}

func (c *resultClient) CallResult(ctx context.Context, in *RequestResult, opts ...grpc.CallOption) (*ResponseResult, error) {
	out := new(ResponseResult)
	err := c.cc.Invoke(ctx, "/proto.Result/CallResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultServer is the server API for Result service.
// All implementations must embed UnimplementedResultServer
// for forward compatibility
type ResultServer interface {
	CallResult(context.Context, *RequestResult) (*ResponseResult, error)
	mustEmbedUnimplementedResultServer()
}

// UnimplementedResultServer must be embedded to have forward compatible implementations.
type UnimplementedResultServer struct {
}

func (UnimplementedResultServer) CallResult(context.Context, *RequestResult) (*ResponseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallResult not implemented")
}
func (UnimplementedResultServer) mustEmbedUnimplementedResultServer() {}

// UnsafeResultServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResultServer will
// result in compilation errors.
type UnsafeResultServer interface {
	mustEmbedUnimplementedResultServer()
}

func RegisterResultServer(s grpc.ServiceRegistrar, srv ResultServer) {
	s.RegisterService(&Result_ServiceDesc, srv)
}

func _Result_CallResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultServer).CallResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Result/CallResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultServer).CallResult(ctx, req.(*RequestResult))
	}
	return interceptor(ctx, in, info, handler)
}

// Result_ServiceDesc is the grpc.ServiceDesc for Result service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Result_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Result",
	HandlerType: (*ResultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallResult",
			Handler:    _Result_CallResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "result.proto",
}
