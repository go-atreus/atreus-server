// Code generated by protoc-gen-go-server. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v0.0.1
// source: api/user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = grpc.Version

type UserImpl struct {
	cc UserClient
}

func NewUserImpl(conn *grpc.ClientConn) *UserImpl {
	return &UserImpl{cc: NewUserClient(conn)}
}

func (c *UserImpl) GetUserInfo(ctx context.Context, in *UserInfoReq) (*SysUser, error) {
	return c.cc.GetUserInfo(ctx, in)
}

func (c *UserImpl) SysUserCreate(ctx context.Context, in *UserCreateReq) (*SysUser, error) {
	return c.cc.SysUserCreate(ctx, in)
}

func (c *UserImpl) ListSysUser(ctx context.Context, in *emptypb.Empty) (*ListUser, error) {
	return c.cc.ListSysUser(ctx, in)
}

func (c *UserImpl) GetUserScope(ctx context.Context, in *SysUser) (*UserScopeResp, error) {
	return c.cc.GetUserScope(ctx, in)
}
