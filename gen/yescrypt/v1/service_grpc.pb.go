// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: yescrypt/v1/service.proto

package yescryptv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	YescryptService_Hash_FullMethodName   = "/yescrypt.v1.YescryptService/Hash"
	YescryptService_Verify_FullMethodName = "/yescrypt.v1.YescryptService/Verify"
)

// YescryptServiceClient is the client API for YescryptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YescryptServiceClient interface {
	Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type yescryptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYescryptServiceClient(cc grpc.ClientConnInterface) YescryptServiceClient {
	return &yescryptServiceClient{cc}
}

func (c *yescryptServiceClient) Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, YescryptService_Hash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yescryptServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, YescryptService_Verify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YescryptServiceServer is the server API for YescryptService service.
// All implementations must embed UnimplementedYescryptServiceServer
// for forward compatibility.
type YescryptServiceServer interface {
	Hash(context.Context, *HashRequest) (*HashResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	mustEmbedUnimplementedYescryptServiceServer()
}

// UnimplementedYescryptServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedYescryptServiceServer struct{}

func (UnimplementedYescryptServiceServer) Hash(context.Context, *HashRequest) (*HashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash not implemented")
}
func (UnimplementedYescryptServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedYescryptServiceServer) mustEmbedUnimplementedYescryptServiceServer() {}
func (UnimplementedYescryptServiceServer) testEmbeddedByValue()                         {}

// UnsafeYescryptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YescryptServiceServer will
// result in compilation errors.
type UnsafeYescryptServiceServer interface {
	mustEmbedUnimplementedYescryptServiceServer()
}

func RegisterYescryptServiceServer(s grpc.ServiceRegistrar, srv YescryptServiceServer) {
	// If the following call pancis, it indicates UnimplementedYescryptServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&YescryptService_ServiceDesc, srv)
}

func _YescryptService_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YescryptServiceServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YescryptService_Hash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YescryptServiceServer).Hash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YescryptService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YescryptServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YescryptService_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YescryptServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YescryptService_ServiceDesc is the grpc.ServiceDesc for YescryptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YescryptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yescrypt.v1.YescryptService",
	HandlerType: (*YescryptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hash",
			Handler:    _YescryptService_Hash_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _YescryptService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yescrypt/v1/service.proto",
}
