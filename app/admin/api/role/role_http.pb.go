// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: api/role/role.proto

package role

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRoleCreateSysRole = "/atreus.role.Role/CreateSysRole"
const OperationRoleDeleteRole = "/atreus.role.Role/DeleteRole"
const OperationRoleGetRole = "/atreus.role.Role/GetRole"
const OperationRoleListRole = "/atreus.role.Role/ListRole"
const OperationRoleRolePermissions = "/atreus.role.Role/RolePermissions"
const OperationRoleRolePermissionsPut = "/atreus.role.Role/RolePermissionsPut"
const OperationRoleRoleSelect = "/atreus.role.Role/RoleSelect"
const OperationRoleUpdateRole = "/atreus.role.Role/UpdateRole"

type RoleHTTPServer interface {
	// CreateSysRole 创建角色
	CreateSysRole(context.Context, *SysRole) (*SysRole, error)
	// DeleteRole 删除角色
	DeleteRole(context.Context, *SysRole) (*emptypb.Empty, error)
	// GetRole 获取角色
	GetRole(context.Context, *SysRole) (*SysRole, error)
	// ListRole 获取角色列表
	ListRole(context.Context, *emptypb.Empty) (*ListRoleResp, error)
	RolePermissions(context.Context, *SysRole) (*RoleMenu, error)
	RolePermissionsPut(context.Context, *RolePermReq) (*RoleMenu, error)
	RoleSelect(context.Context, *emptypb.Empty) (*RoleSelectResp, error)
	// UpdateRole 更新角色
	UpdateRole(context.Context, *SysRole) (*SysRole, error)
}

func RegisterRoleHTTPServer(s *http.Server, srv RoleHTTPServer) {
	r := s.Route("/")
	r.POST("/system/role/create", _Role_CreateSysRole0_HTTP_Handler(srv))
	r.PUT("/system/role", _Role_UpdateRole0_HTTP_Handler(srv))
	r.DELETE("/system/role/delete", _Role_DeleteRole0_HTTP_Handler(srv))
	r.GET("/system/role/get", _Role_GetRole0_HTTP_Handler(srv))
	r.POST("/system/role/list", _Role_ListRole0_HTTP_Handler(srv))
	r.GET("/system/permission/code/{code}", _Role_RolePermissions0_HTTP_Handler(srv))
	r.PUT("/system/permission/code/{code}", _Role_RolePermissionsPut0_HTTP_Handler(srv))
	r.GET("/system/role/select", _Role_RoleSelect0_HTTP_Handler(srv))
}

func _Role_CreateSysRole0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysRole
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleCreateSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysRole(ctx, req.(*SysRole))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysRole)
		return ctx.Result(200, reply)
	}
}

func _Role_UpdateRole0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysRole
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleUpdateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateRole(ctx, req.(*SysRole))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysRole)
		return ctx.Result(200, reply)
	}
}

func _Role_DeleteRole0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysRole
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleDeleteRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRole(ctx, req.(*SysRole))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Role_GetRole0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysRole
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleGetRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRole(ctx, req.(*SysRole))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysRole)
		return ctx.Result(200, reply)
	}
}

func _Role_ListRole0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleListRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRole(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRoleResp)
		return ctx.Result(200, reply)
	}
}

func _Role_RolePermissions0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysRole
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleRolePermissions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RolePermissions(ctx, req.(*SysRole))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleMenu)
		return ctx.Result(200, reply)
	}
}

func _Role_RolePermissionsPut0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RolePermReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleRolePermissionsPut)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RolePermissionsPut(ctx, req.(*RolePermReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleMenu)
		return ctx.Result(200, reply)
	}
}

func _Role_RoleSelect0_HTTP_Handler(srv RoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleRoleSelect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleSelect(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleSelectResp)
		return ctx.Result(200, reply)
	}
}

type RoleHTTPClient interface {
	CreateSysRole(ctx context.Context, req *SysRole, opts ...http.CallOption) (rsp *SysRole, err error)
	DeleteRole(ctx context.Context, req *SysRole, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetRole(ctx context.Context, req *SysRole, opts ...http.CallOption) (rsp *SysRole, err error)
	ListRole(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListRoleResp, err error)
	RolePermissions(ctx context.Context, req *SysRole, opts ...http.CallOption) (rsp *RoleMenu, err error)
	RolePermissionsPut(ctx context.Context, req *RolePermReq, opts ...http.CallOption) (rsp *RoleMenu, err error)
	RoleSelect(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *RoleSelectResp, err error)
	UpdateRole(ctx context.Context, req *SysRole, opts ...http.CallOption) (rsp *SysRole, err error)
}

type RoleHTTPClientImpl struct {
	cc *http.Client
}

func NewRoleHTTPClient(client *http.Client) RoleHTTPClient {
	return &RoleHTTPClientImpl{client}
}

func (c *RoleHTTPClientImpl) CreateSysRole(ctx context.Context, in *SysRole, opts ...http.CallOption) (*SysRole, error) {
	var out SysRole
	pattern := "/system/role/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleCreateSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHTTPClientImpl) DeleteRole(ctx context.Context, in *SysRole, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/system/role/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoleDeleteRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHTTPClientImpl) GetRole(ctx context.Context, in *SysRole, opts ...http.CallOption) (*SysRole, error) {
	var out SysRole
	pattern := "/system/role/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoleGetRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHTTPClientImpl) ListRole(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListRoleResp, error) {
	var out ListRoleResp
	pattern := "/system/role/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleListRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHTTPClientImpl) RolePermissions(ctx context.Context, in *SysRole, opts ...http.CallOption) (*RoleMenu, error) {
	var out RoleMenu
	pattern := "/system/permission/code/{code}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoleRolePermissions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHTTPClientImpl) RolePermissionsPut(ctx context.Context, in *RolePermReq, opts ...http.CallOption) (*RoleMenu, error) {
	var out RoleMenu
	pattern := "/system/permission/code/{code}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleRolePermissionsPut))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHTTPClientImpl) RoleSelect(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*RoleSelectResp, error) {
	var out RoleSelectResp
	pattern := "/system/role/select"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRoleRoleSelect))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHTTPClientImpl) UpdateRole(ctx context.Context, in *SysRole, opts ...http.CallOption) (*SysRole, error) {
	var out SysRole
	pattern := "/system/role"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleUpdateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
