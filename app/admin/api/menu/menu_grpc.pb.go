// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/menu/menu.proto

package menu

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

const (
	Menu_GetMenu_FullMethodName       = "/atreus.menu.Menu/GetMenu"
	Menu_ListSysMenu_FullMethodName   = "/atreus.menu.Menu/ListSysMenu"
	Menu_CreateSysMenu_FullMethodName = "/atreus.menu.Menu/CreateSysMenu"
)

// MenuClient is the client API for Menu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuClient interface {
	// 获取菜单树
	GetMenu(ctx context.Context, in *GetMenuReq, opts ...grpc.CallOption) (*GetMenuResp, error)
	// 分页获取基础menu列表
	ListSysMenu(ctx context.Context, in *GetMenuReq, opts ...grpc.CallOption) (*ListSysMenuResp, error)
	CreateSysMenu(ctx context.Context, in *SysMenu, opts ...grpc.CallOption) (*SysMenu, error)
}

type menuClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuClient(cc grpc.ClientConnInterface) MenuClient {
	return &menuClient{cc}
}

func (c *menuClient) GetMenu(ctx context.Context, in *GetMenuReq, opts ...grpc.CallOption) (*GetMenuResp, error) {
	out := new(GetMenuResp)
	err := c.cc.Invoke(ctx, Menu_GetMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuClient) ListSysMenu(ctx context.Context, in *GetMenuReq, opts ...grpc.CallOption) (*ListSysMenuResp, error) {
	out := new(ListSysMenuResp)
	err := c.cc.Invoke(ctx, Menu_ListSysMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuClient) CreateSysMenu(ctx context.Context, in *SysMenu, opts ...grpc.CallOption) (*SysMenu, error) {
	out := new(SysMenu)
	err := c.cc.Invoke(ctx, Menu_CreateSysMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServer is the server API for Menu service.
// All implementations should embed UnimplementedMenuServer
// for forward compatibility
type MenuServer interface {
	// 获取菜单树
	GetMenu(context.Context, *GetMenuReq) (*GetMenuResp, error)
	// 分页获取基础menu列表
	ListSysMenu(context.Context, *GetMenuReq) (*ListSysMenuResp, error)
	CreateSysMenu(context.Context, *SysMenu) (*SysMenu, error)
}

// UnimplementedMenuServer should be embedded to have forward compatible implementations.
type UnimplementedMenuServer struct {
}

func (UnimplementedMenuServer) GetMenu(context.Context, *GetMenuReq) (*GetMenuResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedMenuServer) ListSysMenu(context.Context, *GetMenuReq) (*ListSysMenuResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSysMenu not implemented")
}
func (UnimplementedMenuServer) CreateSysMenu(context.Context, *SysMenu) (*SysMenu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysMenu not implemented")
}

// UnsafeMenuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServer will
// result in compilation errors.
type UnsafeMenuServer interface {
	mustEmbedUnimplementedMenuServer()
}

func RegisterMenuServer(s grpc.ServiceRegistrar, srv MenuServer) {
	s.RegisterService(&Menu_ServiceDesc, srv)
}

func _Menu_GetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).GetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Menu_GetMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).GetMenu(ctx, req.(*GetMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Menu_ListSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).ListSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Menu_ListSysMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).ListSysMenu(ctx, req.(*GetMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Menu_CreateSysMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysMenu)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).CreateSysMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Menu_CreateSysMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).CreateSysMenu(ctx, req.(*SysMenu))
	}
	return interceptor(ctx, in, info, handler)
}

// Menu_ServiceDesc is the grpc.ServiceDesc for Menu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Menu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "atreus.menu.Menu",
	HandlerType: (*MenuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMenu",
			Handler:    _Menu_GetMenu_Handler,
		},
		{
			MethodName: "ListSysMenu",
			Handler:    _Menu_ListSysMenu_Handler,
		},
		{
			MethodName: "CreateSysMenu",
			Handler:    _Menu_CreateSysMenu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/menu/menu.proto",
}
