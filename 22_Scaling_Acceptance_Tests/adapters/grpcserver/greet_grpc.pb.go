// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: greet.proto

package grpcserver

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReply, error)
	Curse(ctx context.Context, in *CurseRequest, opts ...grpc.CallOption) (*CurseReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReply, error) {
	out := new(GreetReply)
	err := c.cc.Invoke(ctx, "/grpcserver.Greeter/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Curse(ctx context.Context, in *CurseRequest, opts ...grpc.CallOption) (*CurseReply, error) {
	out := new(CurseReply)
	err := c.cc.Invoke(ctx, "/grpcserver.Greeter/Curse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	Greet(context.Context, *GreetRequest) (*GreetReply, error)
	Curse(context.Context, *CurseRequest) (*CurseReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Greet(context.Context, *GreetRequest) (*GreetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreeterServer) Curse(context.Context, *CurseRequest) (*CurseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Curse not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcserver.Greeter/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Curse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Curse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcserver.Greeter/Curse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Curse(ctx, req.(*CurseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcserver.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Greeter_Greet_Handler,
		},
		{
			MethodName: "Curse",
			Handler:    _Greeter_Curse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}
