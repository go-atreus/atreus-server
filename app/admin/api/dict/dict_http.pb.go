// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: api/dict/dict.proto

package dict

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

const OperationDictCreate = "/atreus.dict.Dict/Create"
const OperationDictDelete = "/atreus.dict.Dict/Delete"
const OperationDictDictData = "/atreus.dict.Dict/DictData"
const OperationDictDictValidHash = "/atreus.dict.Dict/DictValidHash"
const OperationDictGet = "/atreus.dict.Dict/Get"
const OperationDictList = "/atreus.dict.Dict/List"
const OperationDictUpdate = "/atreus.dict.Dict/Update"

type DictHTTPServer interface {
	// Create 创建字典
	Create(context.Context, *SysDict) (*SysDict, error)
	// Delete 删除字典
	Delete(context.Context, *SysDict) (*SysDict, error)
	DictData(context.Context, *DictDataReq) (*DictDataResp, error)
	DictValidHash(context.Context, *emptypb.Empty) (*ValidHashResp, error)
	// Get 获取字典
	Get(context.Context, *SysDict) (*SysDict, error)
	// List 获取字典列表
	List(context.Context, *SysDict) (*ListSysDictResp, error)
	// Update 更新字典
	Update(context.Context, *SysDict) (*SysDict, error)
}

func RegisterDictHTTPServer(s *http.Server, srv DictHTTPServer) {
	r := s.Route("/")
	r.GET("/system/dict/data", _Dict_DictData0_HTTP_Handler(srv))
	r.POST("/system/dict/invalid-hash", _Dict_DictValidHash0_HTTP_Handler(srv))
	r.POST("/system/dict/create", _Dict_Create0_HTTP_Handler(srv))
	r.PUT("/dict", _Dict_Update0_HTTP_Handler(srv))
	r.DELETE("/dict", _Dict_Delete0_HTTP_Handler(srv))
	r.GET("/dict", _Dict_Get0_HTTP_Handler(srv))
	r.POST("/system/dict/list", _Dict_List0_HTTP_Handler(srv))
}

func _Dict_DictData0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DictDataReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DictData(ctx, req.(*DictDataReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DictDataResp)
		return ctx.Result(200, reply)
	}
}

func _Dict_DictValidHash0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDictValidHash)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DictValidHash(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ValidHashResp)
		return ctx.Result(200, reply)
	}
}

func _Dict_Create0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysDict
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*SysDict))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysDict)
		return ctx.Result(200, reply)
	}
}

func _Dict_Update0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysDict
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*SysDict))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysDict)
		return ctx.Result(200, reply)
	}
}

func _Dict_Delete0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysDict
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*SysDict))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysDict)
		return ctx.Result(200, reply)
	}
}

func _Dict_Get0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysDict
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*SysDict))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SysDict)
		return ctx.Result(200, reply)
	}
}

func _Dict_List0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysDict
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*SysDict))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSysDictResp)
		return ctx.Result(200, reply)
	}
}

type DictHTTPClient interface {
	Create(ctx context.Context, req *SysDict, opts ...http.CallOption) (rsp *SysDict, err error)
	Delete(ctx context.Context, req *SysDict, opts ...http.CallOption) (rsp *SysDict, err error)
	DictData(ctx context.Context, req *DictDataReq, opts ...http.CallOption) (rsp *DictDataResp, err error)
	DictValidHash(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ValidHashResp, err error)
	Get(ctx context.Context, req *SysDict, opts ...http.CallOption) (rsp *SysDict, err error)
	List(ctx context.Context, req *SysDict, opts ...http.CallOption) (rsp *ListSysDictResp, err error)
	Update(ctx context.Context, req *SysDict, opts ...http.CallOption) (rsp *SysDict, err error)
}

type DictHTTPClientImpl struct {
	cc *http.Client
}

func NewDictHTTPClient(client *http.Client) DictHTTPClient {
	return &DictHTTPClientImpl{client}
}

func (c *DictHTTPClientImpl) Create(ctx context.Context, in *SysDict, opts ...http.CallOption) (*SysDict, error) {
	var out SysDict
	pattern := "/system/dict/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) Delete(ctx context.Context, in *SysDict, opts ...http.CallOption) (*SysDict, error) {
	var out SysDict
	pattern := "/dict"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) DictData(ctx context.Context, in *DictDataReq, opts ...http.CallOption) (*DictDataResp, error) {
	var out DictDataResp
	pattern := "/system/dict/data"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) DictValidHash(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ValidHashResp, error) {
	var out ValidHashResp
	pattern := "/system/dict/invalid-hash"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictDictValidHash))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) Get(ctx context.Context, in *SysDict, opts ...http.CallOption) (*SysDict, error) {
	var out SysDict
	pattern := "/dict"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) List(ctx context.Context, in *SysDict, opts ...http.CallOption) (*ListSysDictResp, error) {
	var out ListSysDictResp
	pattern := "/system/dict/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) Update(ctx context.Context, in *SysDict, opts ...http.CallOption) (*SysDict, error) {
	var out SysDict
	pattern := "/dict"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
