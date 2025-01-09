// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: dashboard.proto

package pb

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
	OwnerService_CreateOwner_FullMethodName = "/dashboard.OwnerService/CreateOwner"
	OwnerService_IsBlocked_FullMethodName   = "/dashboard.OwnerService/IsBlocked"
)

// OwnerServiceClient is the client API for OwnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OwnerServiceClient interface {
	CreateOwner(ctx context.Context, in *CreateOwnerReq, opts ...grpc.CallOption) (*CreateOwnerReq, error)
	IsBlocked(ctx context.Context, in *IsBlockedReq, opts ...grpc.CallOption) (*IsBlockedRes, error)
}

type ownerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOwnerServiceClient(cc grpc.ClientConnInterface) OwnerServiceClient {
	return &ownerServiceClient{cc}
}

func (c *ownerServiceClient) CreateOwner(ctx context.Context, in *CreateOwnerReq, opts ...grpc.CallOption) (*CreateOwnerReq, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOwnerReq)
	err := c.cc.Invoke(ctx, OwnerService_CreateOwner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerServiceClient) IsBlocked(ctx context.Context, in *IsBlockedReq, opts ...grpc.CallOption) (*IsBlockedRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsBlockedRes)
	err := c.cc.Invoke(ctx, OwnerService_IsBlocked_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerServiceServer is the server API for OwnerService service.
// All implementations must embed UnimplementedOwnerServiceServer
// for forward compatibility.
type OwnerServiceServer interface {
	CreateOwner(context.Context, *CreateOwnerReq) (*CreateOwnerReq, error)
	IsBlocked(context.Context, *IsBlockedReq) (*IsBlockedRes, error)
	mustEmbedUnimplementedOwnerServiceServer()
}

// UnimplementedOwnerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOwnerServiceServer struct{}

func (UnimplementedOwnerServiceServer) CreateOwner(context.Context, *CreateOwnerReq) (*CreateOwnerReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOwner not implemented")
}
func (UnimplementedOwnerServiceServer) IsBlocked(context.Context, *IsBlockedReq) (*IsBlockedRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBlocked not implemented")
}
func (UnimplementedOwnerServiceServer) mustEmbedUnimplementedOwnerServiceServer() {}
func (UnimplementedOwnerServiceServer) testEmbeddedByValue()                      {}

// UnsafeOwnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OwnerServiceServer will
// result in compilation errors.
type UnsafeOwnerServiceServer interface {
	mustEmbedUnimplementedOwnerServiceServer()
}

func RegisterOwnerServiceServer(s grpc.ServiceRegistrar, srv OwnerServiceServer) {
	// If the following call pancis, it indicates UnimplementedOwnerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OwnerService_ServiceDesc, srv)
}

func _OwnerService_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOwnerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnerService_CreateOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).CreateOwner(ctx, req.(*CreateOwnerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerService_IsBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBlockedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).IsBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OwnerService_IsBlocked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).IsBlocked(ctx, req.(*IsBlockedReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OwnerService_ServiceDesc is the grpc.ServiceDesc for OwnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OwnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.OwnerService",
	HandlerType: (*OwnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOwner",
			Handler:    _OwnerService_CreateOwner_Handler,
		},
		{
			MethodName: "IsBlocked",
			Handler:    _OwnerService_IsBlocked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboard.proto",
}
