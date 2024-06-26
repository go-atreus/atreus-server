// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/organization/organization.proto

package organization

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
	Organization_CreateOrganization_FullMethodName = "/atreus.organization.Organization/CreateOrganization"
	Organization_QueryOrganization_FullMethodName  = "/atreus.organization.Organization/QueryOrganization"
	Organization_UpdateOrganization_FullMethodName = "/atreus.organization.Organization/UpdateOrganization"
	Organization_DeleteOrganization_FullMethodName = "/atreus.organization.Organization/DeleteOrganization"
	Organization_ListOrganization_FullMethodName   = "/atreus.organization.Organization/ListOrganization"
	Organization_OrganizationTree_FullMethodName   = "/atreus.organization.Organization/OrganizationTree"
)

// OrganizationClient is the client API for Organization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationClient interface {
	// 创建组织
	CreateOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error)
	// 查询组织
	QueryOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error)
	// 更新组织
	UpdateOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error)
	// 删除组织
	DeleteOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error)
	// 获取所有组织
	ListOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*ListSysOrganization, error)
	// 获取组织树
	OrganizationTree(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*ListSysOrganization, error)
}

type organizationClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationClient(cc grpc.ClientConnInterface) OrganizationClient {
	return &organizationClient{cc}
}

func (c *organizationClient) CreateOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error) {
	out := new(SysOrganization)
	err := c.cc.Invoke(ctx, Organization_CreateOrganization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) QueryOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error) {
	out := new(SysOrganization)
	err := c.cc.Invoke(ctx, Organization_QueryOrganization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) UpdateOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error) {
	out := new(SysOrganization)
	err := c.cc.Invoke(ctx, Organization_UpdateOrganization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) DeleteOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*SysOrganization, error) {
	out := new(SysOrganization)
	err := c.cc.Invoke(ctx, Organization_DeleteOrganization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) ListOrganization(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*ListSysOrganization, error) {
	out := new(ListSysOrganization)
	err := c.cc.Invoke(ctx, Organization_ListOrganization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) OrganizationTree(ctx context.Context, in *SysOrganization, opts ...grpc.CallOption) (*ListSysOrganization, error) {
	out := new(ListSysOrganization)
	err := c.cc.Invoke(ctx, Organization_OrganizationTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServer is the server API for Organization service.
// All implementations should embed UnimplementedOrganizationServer
// for forward compatibility
type OrganizationServer interface {
	// 创建组织
	CreateOrganization(context.Context, *SysOrganization) (*SysOrganization, error)
	// 查询组织
	QueryOrganization(context.Context, *SysOrganization) (*SysOrganization, error)
	// 更新组织
	UpdateOrganization(context.Context, *SysOrganization) (*SysOrganization, error)
	// 删除组织
	DeleteOrganization(context.Context, *SysOrganization) (*SysOrganization, error)
	// 获取所有组织
	ListOrganization(context.Context, *SysOrganization) (*ListSysOrganization, error)
	// 获取组织树
	OrganizationTree(context.Context, *SysOrganization) (*ListSysOrganization, error)
}

// UnimplementedOrganizationServer should be embedded to have forward compatible implementations.
type UnimplementedOrganizationServer struct {
}

func (UnimplementedOrganizationServer) CreateOrganization(context.Context, *SysOrganization) (*SysOrganization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOrganizationServer) QueryOrganization(context.Context, *SysOrganization) (*SysOrganization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrganization not implemented")
}
func (UnimplementedOrganizationServer) UpdateOrganization(context.Context, *SysOrganization) (*SysOrganization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationServer) DeleteOrganization(context.Context, *SysOrganization) (*SysOrganization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrganization not implemented")
}
func (UnimplementedOrganizationServer) ListOrganization(context.Context, *SysOrganization) (*ListSysOrganization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganization not implemented")
}
func (UnimplementedOrganizationServer) OrganizationTree(context.Context, *SysOrganization) (*ListSysOrganization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrganizationTree not implemented")
}

// UnsafeOrganizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationServer will
// result in compilation errors.
type UnsafeOrganizationServer interface {
	mustEmbedUnimplementedOrganizationServer()
}

func RegisterOrganizationServer(s grpc.ServiceRegistrar, srv OrganizationServer) {
	s.RegisterService(&Organization_ServiceDesc, srv)
}

func _Organization_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOrganization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_CreateOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).CreateOrganization(ctx, req.(*SysOrganization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_QueryOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOrganization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).QueryOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_QueryOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).QueryOrganization(ctx, req.(*SysOrganization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOrganization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_UpdateOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).UpdateOrganization(ctx, req.(*SysOrganization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_DeleteOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOrganization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).DeleteOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_DeleteOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).DeleteOrganization(ctx, req.(*SysOrganization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_ListOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOrganization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).ListOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_ListOrganization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).ListOrganization(ctx, req.(*SysOrganization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_OrganizationTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysOrganization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).OrganizationTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Organization_OrganizationTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).OrganizationTree(ctx, req.(*SysOrganization))
	}
	return interceptor(ctx, in, info, handler)
}

// Organization_ServiceDesc is the grpc.ServiceDesc for Organization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Organization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "atreus.organization.Organization",
	HandlerType: (*OrganizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrganization",
			Handler:    _Organization_CreateOrganization_Handler,
		},
		{
			MethodName: "QueryOrganization",
			Handler:    _Organization_QueryOrganization_Handler,
		},
		{
			MethodName: "UpdateOrganization",
			Handler:    _Organization_UpdateOrganization_Handler,
		},
		{
			MethodName: "DeleteOrganization",
			Handler:    _Organization_DeleteOrganization_Handler,
		},
		{
			MethodName: "ListOrganization",
			Handler:    _Organization_ListOrganization_Handler,
		},
		{
			MethodName: "OrganizationTree",
			Handler:    _Organization_OrganizationTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/organization/organization.proto",
}
