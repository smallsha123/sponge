// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/serverNameExample/v1/userExample.proto

package v1

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

// UserExampleServiceClient is the client API for UserExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserExampleServiceClient interface {
	Create(ctx context.Context, in *CreateUserExampleRequest, opts ...grpc.CallOption) (*CreateUserExampleReply, error)
	DeleteByID(ctx context.Context, in *DeleteUserExampleByIDRequest, opts ...grpc.CallOption) (*DeleteUserExampleByIDReply, error)
	UpdateByID(ctx context.Context, in *UpdateUserExampleByIDRequest, opts ...grpc.CallOption) (*UpdateUserExampleByIDReply, error)
	GetByID(ctx context.Context, in *GetUserExampleByIDRequest, opts ...grpc.CallOption) (*GetUserExampleByIDReply, error)
	ListByIDs(ctx context.Context, in *ListUserExampleByIDsRequest, opts ...grpc.CallOption) (*ListUserExampleByIDsReply, error)
	List(ctx context.Context, in *ListUserExampleRequest, opts ...grpc.CallOption) (*ListUserExampleReply, error)
}

type userExampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserExampleServiceClient(cc grpc.ClientConnInterface) UserExampleServiceClient {
	return &userExampleServiceClient{cc}
}

func (c *userExampleServiceClient) Create(ctx context.Context, in *CreateUserExampleRequest, opts ...grpc.CallOption) (*CreateUserExampleReply, error) {
	out := new(CreateUserExampleReply)
	err := c.cc.Invoke(ctx, "/api.serverNameExample.v1.userExampleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExampleServiceClient) DeleteByID(ctx context.Context, in *DeleteUserExampleByIDRequest, opts ...grpc.CallOption) (*DeleteUserExampleByIDReply, error) {
	out := new(DeleteUserExampleByIDReply)
	err := c.cc.Invoke(ctx, "/api.serverNameExample.v1.userExampleService/DeleteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExampleServiceClient) UpdateByID(ctx context.Context, in *UpdateUserExampleByIDRequest, opts ...grpc.CallOption) (*UpdateUserExampleByIDReply, error) {
	out := new(UpdateUserExampleByIDReply)
	err := c.cc.Invoke(ctx, "/api.serverNameExample.v1.userExampleService/UpdateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExampleServiceClient) GetByID(ctx context.Context, in *GetUserExampleByIDRequest, opts ...grpc.CallOption) (*GetUserExampleByIDReply, error) {
	out := new(GetUserExampleByIDReply)
	err := c.cc.Invoke(ctx, "/api.serverNameExample.v1.userExampleService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExampleServiceClient) ListByIDs(ctx context.Context, in *ListUserExampleByIDsRequest, opts ...grpc.CallOption) (*ListUserExampleByIDsReply, error) {
	out := new(ListUserExampleByIDsReply)
	err := c.cc.Invoke(ctx, "/api.serverNameExample.v1.userExampleService/ListByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExampleServiceClient) List(ctx context.Context, in *ListUserExampleRequest, opts ...grpc.CallOption) (*ListUserExampleReply, error) {
	out := new(ListUserExampleReply)
	err := c.cc.Invoke(ctx, "/api.serverNameExample.v1.userExampleService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserExampleServiceServer is the server API for UserExampleService service.
// All implementations must embed UnimplementedUserExampleServiceServer
// for forward compatibility
type UserExampleServiceServer interface {
	Create(context.Context, *CreateUserExampleRequest) (*CreateUserExampleReply, error)
	DeleteByID(context.Context, *DeleteUserExampleByIDRequest) (*DeleteUserExampleByIDReply, error)
	UpdateByID(context.Context, *UpdateUserExampleByIDRequest) (*UpdateUserExampleByIDReply, error)
	GetByID(context.Context, *GetUserExampleByIDRequest) (*GetUserExampleByIDReply, error)
	ListByIDs(context.Context, *ListUserExampleByIDsRequest) (*ListUserExampleByIDsReply, error)
	List(context.Context, *ListUserExampleRequest) (*ListUserExampleReply, error)
	mustEmbedUnimplementedUserExampleServiceServer()
}

// UnimplementedUserExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserExampleServiceServer struct {
}

func (UnimplementedUserExampleServiceServer) Create(context.Context, *CreateUserExampleRequest) (*CreateUserExampleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserExampleServiceServer) DeleteByID(context.Context, *DeleteUserExampleByIDRequest) (*DeleteUserExampleByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedUserExampleServiceServer) UpdateByID(context.Context, *UpdateUserExampleByIDRequest) (*UpdateUserExampleByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByID not implemented")
}
func (UnimplementedUserExampleServiceServer) GetByID(context.Context, *GetUserExampleByIDRequest) (*GetUserExampleByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserExampleServiceServer) ListByIDs(context.Context, *ListUserExampleByIDsRequest) (*ListUserExampleByIDsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByIDs not implemented")
}
func (UnimplementedUserExampleServiceServer) List(context.Context, *ListUserExampleRequest) (*ListUserExampleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserExampleServiceServer) mustEmbedUnimplementedUserExampleServiceServer() {}

// UnsafeUserExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserExampleServiceServer will
// result in compilation errors.
type UnsafeUserExampleServiceServer interface {
	mustEmbedUnimplementedUserExampleServiceServer()
}

func RegisterUserExampleServiceServer(s grpc.ServiceRegistrar, srv UserExampleServiceServer) {
	s.RegisterService(&UserExampleService_ServiceDesc, srv)
}

func _UserExampleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExampleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serverNameExample.v1.userExampleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExampleServiceServer).Create(ctx, req.(*CreateUserExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExampleService_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserExampleByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExampleServiceServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serverNameExample.v1.userExampleService/DeleteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExampleServiceServer).DeleteByID(ctx, req.(*DeleteUserExampleByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExampleService_UpdateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserExampleByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExampleServiceServer).UpdateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serverNameExample.v1.userExampleService/UpdateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExampleServiceServer).UpdateByID(ctx, req.(*UpdateUserExampleByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExampleService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserExampleByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExampleServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serverNameExample.v1.userExampleService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExampleServiceServer).GetByID(ctx, req.(*GetUserExampleByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExampleService_ListByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserExampleByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExampleServiceServer).ListByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serverNameExample.v1.userExampleService/ListByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExampleServiceServer).ListByIDs(ctx, req.(*ListUserExampleByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExampleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserExampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExampleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serverNameExample.v1.userExampleService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExampleServiceServer).List(ctx, req.(*ListUserExampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserExampleService_ServiceDesc is the grpc.ServiceDesc for UserExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.serverNameExample.v1.userExampleService",
	HandlerType: (*UserExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserExampleService_Create_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _UserExampleService_DeleteByID_Handler,
		},
		{
			MethodName: "UpdateByID",
			Handler:    _UserExampleService_UpdateByID_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _UserExampleService_GetByID_Handler,
		},
		{
			MethodName: "ListByIDs",
			Handler:    _UserExampleService_ListByIDs_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserExampleService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/serverNameExample/v1/userExample.proto",
}
