// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: admin.proto

package __

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
	AdminService_AdminLogin_FullMethodName        = "/pbA.AdminService/AdminLogin"
	AdminService_CreateUserFn_FullMethodName      = "/pbA.AdminService/CreateUserFn"
	AdminService_FindUserByEmailFn_FullMethodName = "/pbA.AdminService/FindUserByEmailFn"
	AdminService_FindUserByIDFn_FullMethodName    = "/pbA.AdminService/FindUserByIDFn"
	AdminService_FindAllUserFn_FullMethodName     = "/pbA.AdminService/FindAllUserFn"
	AdminService_EditUserFn_FullMethodName        = "/pbA.AdminService/EditUserFn"
	AdminService_DeleteUserFn_FullMethodName      = "/pbA.AdminService/DeleteUserFn"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AdminLogin(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminRespone, error)
	CreateUserFn(ctx context.Context, in *User, opts ...grpc.CallOption) (*AdminRespone, error)
	FindUserByEmailFn(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	FindUserByIDFn(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
	FindAllUserFn(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*UserList, error)
	EditUserFn(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUserFn(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*AdminRespone, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AdminLogin(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminRespone)
	err := c.cc.Invoke(ctx, AdminService_AdminLogin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CreateUserFn(ctx context.Context, in *User, opts ...grpc.CallOption) (*AdminRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminRespone)
	err := c.cc.Invoke(ctx, AdminService_CreateUserFn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindUserByEmailFn(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, AdminService_FindUserByEmailFn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindUserByIDFn(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, AdminService_FindUserByIDFn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindAllUserFn(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*UserList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserList)
	err := c.cc.Invoke(ctx, AdminService_FindAllUserFn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) EditUserFn(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, AdminService_EditUserFn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteUserFn(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*AdminRespone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminRespone)
	err := c.cc.Invoke(ctx, AdminService_DeleteUserFn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	AdminLogin(context.Context, *AdminRequest) (*AdminRespone, error)
	CreateUserFn(context.Context, *User) (*AdminRespone, error)
	FindUserByEmailFn(context.Context, *UserRequest) (*User, error)
	FindUserByIDFn(context.Context, *UserID) (*User, error)
	FindAllUserFn(context.Context, *NoParam) (*UserList, error)
	EditUserFn(context.Context, *User) (*User, error)
	DeleteUserFn(context.Context, *UserID) (*AdminRespone, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AdminLogin(context.Context, *AdminRequest) (*AdminRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServiceServer) CreateUserFn(context.Context, *User) (*AdminRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserFn not implemented")
}
func (UnimplementedAdminServiceServer) FindUserByEmailFn(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByEmailFn not implemented")
}
func (UnimplementedAdminServiceServer) FindUserByIDFn(context.Context, *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByIDFn not implemented")
}
func (UnimplementedAdminServiceServer) FindAllUserFn(context.Context, *NoParam) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllUserFn not implemented")
}
func (UnimplementedAdminServiceServer) EditUserFn(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserFn not implemented")
}
func (UnimplementedAdminServiceServer) DeleteUserFn(context.Context, *UserID) (*AdminRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserFn not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminLogin(ctx, req.(*AdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CreateUserFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateUserFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CreateUserFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateUserFn(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindUserByEmailFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindUserByEmailFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindUserByEmailFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindUserByEmailFn(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindUserByIDFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindUserByIDFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindUserByIDFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindUserByIDFn(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindAllUserFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindAllUserFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindAllUserFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindAllUserFn(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_EditUserFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).EditUserFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_EditUserFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).EditUserFn(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteUserFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteUserFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteUserFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteUserFn(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbA.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLogin",
			Handler:    _AdminService_AdminLogin_Handler,
		},
		{
			MethodName: "CreateUserFn",
			Handler:    _AdminService_CreateUserFn_Handler,
		},
		{
			MethodName: "FindUserByEmailFn",
			Handler:    _AdminService_FindUserByEmailFn_Handler,
		},
		{
			MethodName: "FindUserByIDFn",
			Handler:    _AdminService_FindUserByIDFn_Handler,
		},
		{
			MethodName: "FindAllUserFn",
			Handler:    _AdminService_FindAllUserFn_Handler,
		},
		{
			MethodName: "EditUserFn",
			Handler:    _AdminService_EditUserFn_Handler,
		},
		{
			MethodName: "DeleteUserFn",
			Handler:    _AdminService_DeleteUserFn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
