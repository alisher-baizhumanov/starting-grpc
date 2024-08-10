// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: auth.proto

package auth_v1

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

// AuthServiceV1Client is the client API for AuthServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceV1Client interface {
	Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*LoginOut, error)
	GetRefreshToken(ctx context.Context, in *GetRefreshTokenIn, opts ...grpc.CallOption) (*GetRefreshTokenOut, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenIn, opts ...grpc.CallOption) (*GetAccessTokenOut, error)
	Check(ctx context.Context, in *CheckIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceV1Client(cc grpc.ClientConnInterface) AuthServiceV1Client {
	return &authServiceV1Client{cc}
}

func (c *authServiceV1Client) Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*LoginOut, error) {
	out := new(LoginOut)
	err := c.cc.Invoke(ctx, "/user_v1.AuthServiceV1/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceV1Client) GetRefreshToken(ctx context.Context, in *GetRefreshTokenIn, opts ...grpc.CallOption) (*GetRefreshTokenOut, error) {
	out := new(GetRefreshTokenOut)
	err := c.cc.Invoke(ctx, "/user_v1.AuthServiceV1/GetRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceV1Client) GetAccessToken(ctx context.Context, in *GetAccessTokenIn, opts ...grpc.CallOption) (*GetAccessTokenOut, error) {
	out := new(GetAccessTokenOut)
	err := c.cc.Invoke(ctx, "/user_v1.AuthServiceV1/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceV1Client) Check(ctx context.Context, in *CheckIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.AuthServiceV1/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceV1Server is the server API for AuthServiceV1 service.
// All implementations must embed UnimplementedAuthServiceV1Server
// for forward compatibility
type AuthServiceV1Server interface {
	Login(context.Context, *LoginIn) (*LoginOut, error)
	GetRefreshToken(context.Context, *GetRefreshTokenIn) (*GetRefreshTokenOut, error)
	GetAccessToken(context.Context, *GetAccessTokenIn) (*GetAccessTokenOut, error)
	Check(context.Context, *CheckIn) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthServiceV1Server()
}

// UnimplementedAuthServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceV1Server struct {
}

func (UnimplementedAuthServiceV1Server) Login(context.Context, *LoginIn) (*LoginOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceV1Server) GetRefreshToken(context.Context, *GetRefreshTokenIn) (*GetRefreshTokenOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshToken not implemented")
}
func (UnimplementedAuthServiceV1Server) GetAccessToken(context.Context, *GetAccessTokenIn) (*GetAccessTokenOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedAuthServiceV1Server) Check(context.Context, *CheckIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedAuthServiceV1Server) mustEmbedUnimplementedAuthServiceV1Server() {}

// UnsafeAuthServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceV1Server will
// result in compilation errors.
type UnsafeAuthServiceV1Server interface {
	mustEmbedUnimplementedAuthServiceV1Server()
}

func RegisterAuthServiceV1Server(s grpc.ServiceRegistrar, srv AuthServiceV1Server) {
	s.RegisterService(&AuthServiceV1_ServiceDesc, srv)
}

func _AuthServiceV1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.AuthServiceV1/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).Login(ctx, req.(*LoginIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceV1_GetRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefreshTokenIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).GetRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.AuthServiceV1/GetRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).GetRefreshToken(ctx, req.(*GetRefreshTokenIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceV1_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.AuthServiceV1/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).GetAccessToken(ctx, req.(*GetAccessTokenIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceV1_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.AuthServiceV1/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).Check(ctx, req.(*CheckIn))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthServiceV1_ServiceDesc is the grpc.ServiceDesc for AuthServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.AuthServiceV1",
	HandlerType: (*AuthServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthServiceV1_Login_Handler,
		},
		{
			MethodName: "GetRefreshToken",
			Handler:    _AuthServiceV1_GetRefreshToken_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _AuthServiceV1_GetAccessToken_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _AuthServiceV1_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
