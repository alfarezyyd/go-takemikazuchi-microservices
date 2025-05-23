// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: user_address.proto

package user_address

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
	UserAddressService_UserAddressStore_FullMethodName    = "/UserAddressService/UserAddressStore"
	UserAddressService_FindUserAddressById_FullMethodName = "/UserAddressService/FindUserAddressById"
)

// UserAddressServiceClient is the client API for UserAddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAddressServiceClient interface {
	UserAddressStore(ctx context.Context, in *UserAddressCreateRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	FindUserAddressById(ctx context.Context, in *UserAddressSearchRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type userAddressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAddressServiceClient(cc grpc.ClientConnInterface) UserAddressServiceClient {
	return &userAddressServiceClient{cc}
}

func (c *userAddressServiceClient) UserAddressStore(ctx context.Context, in *UserAddressCreateRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, UserAddressService_UserAddressStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAddressServiceClient) FindUserAddressById(ctx context.Context, in *UserAddressSearchRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, UserAddressService_FindUserAddressById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAddressServiceServer is the server API for UserAddressService service.
// All implementations must embed UnimplementedUserAddressServiceServer
// for forward compatibility.
type UserAddressServiceServer interface {
	UserAddressStore(context.Context, *UserAddressCreateRequest) (*QueryResponse, error)
	FindUserAddressById(context.Context, *UserAddressSearchRequest) (*QueryResponse, error)
	mustEmbedUnimplementedUserAddressServiceServer()
}

// UnimplementedUserAddressServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserAddressServiceServer struct{}

func (UnimplementedUserAddressServiceServer) UserAddressStore(context.Context, *UserAddressCreateRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAddressStore not implemented")
}
func (UnimplementedUserAddressServiceServer) FindUserAddressById(context.Context, *UserAddressSearchRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserAddressById not implemented")
}
func (UnimplementedUserAddressServiceServer) mustEmbedUnimplementedUserAddressServiceServer() {}
func (UnimplementedUserAddressServiceServer) testEmbeddedByValue()                            {}

// UnsafeUserAddressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAddressServiceServer will
// result in compilation errors.
type UnsafeUserAddressServiceServer interface {
	mustEmbedUnimplementedUserAddressServiceServer()
}

func RegisterUserAddressServiceServer(s grpc.ServiceRegistrar, srv UserAddressServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserAddressServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserAddressService_ServiceDesc, srv)
}

func _UserAddressService_UserAddressStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddressCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAddressServiceServer).UserAddressStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAddressService_UserAddressStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAddressServiceServer).UserAddressStore(ctx, req.(*UserAddressCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAddressService_FindUserAddressById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddressSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAddressServiceServer).FindUserAddressById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAddressService_FindUserAddressById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAddressServiceServer).FindUserAddressById(ctx, req.(*UserAddressSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAddressService_ServiceDesc is the grpc.ServiceDesc for UserAddressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAddressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserAddressService",
	HandlerType: (*UserAddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAddressStore",
			Handler:    _UserAddressService_UserAddressStore_Handler,
		},
		{
			MethodName: "FindUserAddressById",
			Handler:    _UserAddressService_FindUserAddressById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_address.proto",
}
