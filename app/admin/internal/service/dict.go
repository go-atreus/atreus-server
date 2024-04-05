package service

import (
	"context"
	"github.com/go-atreus/atreus-server/app/admin/api/dict"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
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
	return &dict.ValidHashResp{Results: []string{"hash"}}, nil
}

func (m *DictService) DictData(ctx context.Context, in *dict.DictDataReq) (*dict.DictDataResp, error) {
	switch in.DictCodes {
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
