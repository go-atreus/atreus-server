// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/role/role.proto

package role

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

const (
	Role_CreateSysRole_FullMethodName   = "/atreus.role.Role/CreateSysRole"
	Role_UpdateRole_FullMethodName      = "/atreus.role.Role/UpdateRole"
	Role_DeleteRole_FullMethodName      = "/atreus.role.Role/DeleteRole"
	Role_GetRole_FullMethodName         = "/atreus.role.Role/GetRole"
	Role_ListRole_FullMethodName        = "/atreus.role.Role/ListRole"
	Role_RolePermissions_FullMethodName = "/atreus.role.Role/RolePermissions"
)

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleClient interface {
	// 创建角色
	CreateSysRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*SysRole, error)
	// 更新角色
	UpdateRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*SysRole, error)
	// 删除角色
	DeleteRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取角色
	GetRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*SysRole, error)
	// 获取角色列表
	ListRole(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRoleResp, error)
	RolePermissions(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*RoleMenu, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) CreateSysRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*SysRole, error) {
	out := new(SysRole)
	err := c.cc.Invoke(ctx, Role_CreateSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) UpdateRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*SysRole, error) {
	out := new(SysRole)
	err := c.cc.Invoke(ctx, Role_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) DeleteRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Role_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetRole(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*SysRole, error) {
	out := new(SysRole)
	err := c.cc.Invoke(ctx, Role_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) ListRole(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRoleResp, error) {
	out := new(ListRoleResp)
	err := c.cc.Invoke(ctx, Role_ListRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) RolePermissions(ctx context.Context, in *SysRole, opts ...grpc.CallOption) (*RoleMenu, error) {
	out := new(RoleMenu)
	err := c.cc.Invoke(ctx, Role_RolePermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
// All implementations should embed UnimplementedRoleServer
// for forward compatibility
type RoleServer interface {
	// 创建角色
	CreateSysRole(context.Context, *SysRole) (*SysRole, error)
	// 更新角色
	UpdateRole(context.Context, *SysRole) (*SysRole, error)
	// 删除角色
	DeleteRole(context.Context, *SysRole) (*emptypb.Empty, error)
	// 获取角色
	GetRole(context.Context, *SysRole) (*SysRole, error)
	// 获取角色列表
	ListRole(context.Context, *emptypb.Empty) (*ListRoleResp, error)
	RolePermissions(context.Context, *SysRole) (*RoleMenu, error)
}

// UnimplementedRoleServer should be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct {
}

func (UnimplementedRoleServer) CreateSysRole(context.Context, *SysRole) (*SysRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysRole not implemented")
}
func (UnimplementedRoleServer) UpdateRole(context.Context, *SysRole) (*SysRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServer) DeleteRole(context.Context, *SysRole) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRoleServer) GetRole(context.Context, *SysRole) (*SysRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRoleServer) ListRole(context.Context, *emptypb.Empty) (*ListRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedRoleServer) RolePermissions(context.Context, *SysRole) (*RoleMenu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RolePermissions not implemented")
}

// UnsafeRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServer will
// result in compilation errors.
type UnsafeRoleServer interface {
	mustEmbedUnimplementedRoleServer()
}

func RegisterRoleServer(s grpc.ServiceRegistrar, srv RoleServer) {
	s.RegisterService(&Role_ServiceDesc, srv)
}

func _Role_CreateSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).CreateSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_CreateSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).CreateSysRole(ctx, req.(*SysRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).UpdateRole(ctx, req.(*SysRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).DeleteRole(ctx, req.(*SysRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetRole(ctx, req.(*SysRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_ListRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).ListRole(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_RolePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).RolePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_RolePermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).RolePermissions(ctx, req.(*SysRole))
	}
	return interceptor(ctx, in, info, handler)
}

// Role_ServiceDesc is the grpc.ServiceDesc for Role service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Role_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "atreus.role.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSysRole",
			Handler:    _Role_CreateSysRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Role_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _Role_DeleteRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _Role_GetRole_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _Role_ListRole_Handler,
		},
		{
			MethodName: "RolePermissions",
			Handler:    _Role_RolePermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/role/role.proto",
}
