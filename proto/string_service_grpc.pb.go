// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.3
// source: string_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	StringService_ToUpperCase_FullMethodName     = "/StringService/ToUpperCase"
	StringService_ToLowerCase_FullMethodName     = "/StringService/ToLowerCase"
	StringService_ReverseString_FullMethodName   = "/StringService/ReverseString"
	StringService_GetStringLength_FullMethodName = "/StringService/GetStringLength"
)

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StringServiceClient interface {
	ToUpperCase(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringReply, error)
	ToLowerCase(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringReply, error)
	ReverseString(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringReply, error)
	GetStringLength(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*LengthReply, error)
}

type stringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStringServiceClient(cc grpc.ClientConnInterface) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) ToUpperCase(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringReply)
	err := c.cc.Invoke(ctx, StringService_ToUpperCase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) ToLowerCase(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringReply)
	err := c.cc.Invoke(ctx, StringService_ToLowerCase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) ReverseString(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringReply)
	err := c.cc.Invoke(ctx, StringService_ReverseString_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) GetStringLength(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*LengthReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LengthReply)
	err := c.cc.Invoke(ctx, StringService_GetStringLength_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServiceServer is the server API for StringService service.
// All implementations must embed UnimplementedStringServiceServer
// for forward compatibility
type StringServiceServer interface {
	ToUpperCase(context.Context, *StringRequest) (*StringReply, error)
	ToLowerCase(context.Context, *StringRequest) (*StringReply, error)
	ReverseString(context.Context, *StringRequest) (*StringReply, error)
	GetStringLength(context.Context, *StringRequest) (*LengthReply, error)
	mustEmbedUnimplementedStringServiceServer()
}

// UnimplementedStringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStringServiceServer struct {
}

func (UnimplementedStringServiceServer) ToUpperCase(context.Context, *StringRequest) (*StringReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToUpperCase not implemented")
}
func (UnimplementedStringServiceServer) ToLowerCase(context.Context, *StringRequest) (*StringReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToLowerCase not implemented")
}
func (UnimplementedStringServiceServer) ReverseString(context.Context, *StringRequest) (*StringReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseString not implemented")
}
func (UnimplementedStringServiceServer) GetStringLength(context.Context, *StringRequest) (*LengthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStringLength not implemented")
}
func (UnimplementedStringServiceServer) mustEmbedUnimplementedStringServiceServer() {}

// UnsafeStringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StringServiceServer will
// result in compilation errors.
type UnsafeStringServiceServer interface {
	mustEmbedUnimplementedStringServiceServer()
}

func RegisterStringServiceServer(s grpc.ServiceRegistrar, srv StringServiceServer) {
	s.RegisterService(&StringService_ServiceDesc, srv)
}

func _StringService_ToUpperCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).ToUpperCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_ToUpperCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).ToUpperCase(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_ToLowerCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).ToLowerCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_ToLowerCase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).ToLowerCase(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_ReverseString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).ReverseString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_ReverseString_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).ReverseString(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_GetStringLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).GetStringLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_GetStringLength_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).GetStringLength(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StringService_ServiceDesc is the grpc.ServiceDesc for StringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ToUpperCase",
			Handler:    _StringService_ToUpperCase_Handler,
		},
		{
			MethodName: "ToLowerCase",
			Handler:    _StringService_ToLowerCase_Handler,
		},
		{
			MethodName: "ReverseString",
			Handler:    _StringService_ReverseString_Handler,
		},
		{
			MethodName: "GetStringLength",
			Handler:    _StringService_GetStringLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "string_service.proto",
}