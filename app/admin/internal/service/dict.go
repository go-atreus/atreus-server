package service

import (
	"context"
	"github.com/go-atreus/atreus-server/app/admin/api/dict"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"strconv"
)

var _ dict.DictServer = (*DictService)(nil)

func NewDictService(logger log.Logger, data *data.Data) *DictService {
	return &DictService{
		DictDefaultServer: dict.DictDefaultServer{DB: data.ORM},
	}
}

type DictService struct {
	dict.DictDefaultServer
}

func (m *DictService) DictValidHash(context.Context, *emptypb.Empty) (*dict.ValidHashResp, error) {
	return &dict.ValidHashResp{Results: []string{strconv.Itoa(rand.Int())}}, nil
}

func (m *DictService) DictData(ctx context.Context, in *dict.DictDataReq) (*dict.DictDataResp, error) {
	switch in.DictCodes {
	case "role_type":
		return &dict.DictDataResp{Results: []*dict.DictData{
			{
				DictCode: "role_type",
				DictItems: []*dict.SysDictDataItem{
					{Value: "1", Name: "系统", Id: 1, Attributes: &dict.SysDictItemAttributes{TagColor: "orange"}},
					{Value: "2", Name: "自定义", Id: 2, Attributes: &dict.SysDictItemAttributes{TagColor: "green"}},
				}},
		}}, nil
	case "gender":
		return &dict.DictDataResp{Results: []*dict.DictData{
			{
				DictCode: "gender",
				DictItems: []*dict.SysDictDataItem{
					{Value: "1", Name: "男", Id: 1},
					{Value: "2", Name: "女", Id: 2},
					{Value: "3", Name: "未知", Id: 3},
				}},
		}}, nil
	case "user_status":
		return &dict.DictDataResp{Results: []*dict.DictData{
			{
				DictCode: "user_status",
				DictItems: []*dict.SysDictDataItem{
					{Value: "1", Name: "正常", Id: 1},
					{Value: "2", Name: "锁定", Id: 2},
				}},
		}}, nil
	case "dict_value_type":
		return &dict.DictDataResp{Results: []*dict.DictData{
			{
				DictCode: "dict_value_type",
				DictItems: []*dict.SysDictDataItem{
					{Value: "1", Name: "Number", Id: 1},
					{Value: "2", Name: "String", Id: 2},
					{Value: "3", Name: "Boolean", Id: 3},
				}},
		}}, nil
	}
	return &dict.DictDataResp{Results: []*dict.DictData{}}, nil
}
