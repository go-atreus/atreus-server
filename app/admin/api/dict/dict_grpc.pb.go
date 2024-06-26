// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/dict/dict.proto

package dict

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
	Dict_DictData_FullMethodName      = "/atreus.dict.Dict/DictData"
	Dict_DictValidHash_FullMethodName = "/atreus.dict.Dict/DictValidHash"
	Dict_Create_FullMethodName        = "/atreus.dict.Dict/Create"
	Dict_Update_FullMethodName        = "/atreus.dict.Dict/Update"
	Dict_Delete_FullMethodName        = "/atreus.dict.Dict/Delete"
	Dict_Get_FullMethodName           = "/atreus.dict.Dict/Get"
	Dict_List_FullMethodName          = "/atreus.dict.Dict/List"
)

// DictClient is the client API for Dict service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictClient interface {
	DictData(ctx context.Context, in *DictDataReq, opts ...grpc.CallOption) (*DictDataResp, error)
	DictValidHash(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ValidHashResp, error)
	// 创建字典
	Create(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error)
	// 更新字典
	Update(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error)
	// 删除字典
	Delete(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error)
	// 获取字典
	Get(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error)
	// 获取字典列表
	List(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*ListSysDictResp, error)
}

type dictClient struct {
	cc grpc.ClientConnInterface
}

func NewDictClient(cc grpc.ClientConnInterface) DictClient {
	return &dictClient{cc}
}

func (c *dictClient) DictData(ctx context.Context, in *DictDataReq, opts ...grpc.CallOption) (*DictDataResp, error) {
	out := new(DictDataResp)
	err := c.cc.Invoke(ctx, Dict_DictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) DictValidHash(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ValidHashResp, error) {
	out := new(ValidHashResp)
	err := c.cc.Invoke(ctx, Dict_DictValidHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) Create(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error) {
	out := new(SysDict)
	err := c.cc.Invoke(ctx, Dict_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) Update(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error) {
	out := new(SysDict)
	err := c.cc.Invoke(ctx, Dict_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) Delete(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error) {
	out := new(SysDict)
	err := c.cc.Invoke(ctx, Dict_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) Get(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*SysDict, error) {
	out := new(SysDict)
	err := c.cc.Invoke(ctx, Dict_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictClient) List(ctx context.Context, in *SysDict, opts ...grpc.CallOption) (*ListSysDictResp, error) {
	out := new(ListSysDictResp)
	err := c.cc.Invoke(ctx, Dict_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServer is the server API for Dict service.
// All implementations should embed UnimplementedDictServer
// for forward compatibility
type DictServer interface {
	DictData(context.Context, *DictDataReq) (*DictDataResp, error)
	DictValidHash(context.Context, *emptypb.Empty) (*ValidHashResp, error)
	// 创建字典
	Create(context.Context, *SysDict) (*SysDict, error)
	// 更新字典
	Update(context.Context, *SysDict) (*SysDict, error)
	// 删除字典
	Delete(context.Context, *SysDict) (*SysDict, error)
	// 获取字典
	Get(context.Context, *SysDict) (*SysDict, error)
	// 获取字典列表
	List(context.Context, *SysDict) (*ListSysDictResp, error)
}

// UnimplementedDictServer should be embedded to have forward compatible implementations.
type UnimplementedDictServer struct {
}

func (UnimplementedDictServer) DictData(context.Context, *DictDataReq) (*DictDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DictData not implemented")
}
func (UnimplementedDictServer) DictValidHash(context.Context, *emptypb.Empty) (*ValidHashResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DictValidHash not implemented")
}
func (UnimplementedDictServer) Create(context.Context, *SysDict) (*SysDict, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDictServer) Update(context.Context, *SysDict) (*SysDict, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDictServer) Delete(context.Context, *SysDict) (*SysDict, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDictServer) Get(context.Context, *SysDict) (*SysDict, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDictServer) List(context.Context, *SysDict) (*ListSysDictResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeDictServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictServer will
// result in compilation errors.
type UnsafeDictServer interface {
	mustEmbedUnimplementedDictServer()
}

func RegisterDictServer(s grpc.ServiceRegistrar, srv DictServer) {
	s.RegisterService(&Dict_ServiceDesc, srv)
}

func _Dict_DictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DictDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).DictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_DictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).DictData(ctx, req.(*DictDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_DictValidHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).DictValidHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_DictValidHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).DictValidHash(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDict)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).Create(ctx, req.(*SysDict))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDict)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).Update(ctx, req.(*SysDict))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDict)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).Delete(ctx, req.(*SysDict))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDict)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).Get(ctx, req.(*SysDict))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dict_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDict)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Dict_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServer).List(ctx, req.(*SysDict))
	}
	return interceptor(ctx, in, info, handler)
}

// Dict_ServiceDesc is the grpc.ServiceDesc for Dict service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dict_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "atreus.dict.Dict",
	HandlerType: (*DictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DictData",
			Handler:    _Dict_DictData_Handler,
		},
		{
			MethodName: "DictValidHash",
			Handler:    _Dict_DictValidHash_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Dict_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Dict_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Dict_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Dict_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Dict_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/dict/dict.proto",
}
