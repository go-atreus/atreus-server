package rpc

import (
	context "context"
	"github.com/go-atreus/protocol/admin/menu"
	"github.com/go-kratos/kratos/v2/log"
)

type MenuServer struct{}

func NewMenuServer(logger log.Logger) *MenuServer {
	return &MenuServer{}

}

func (MenuServer) GetMenu(ctx context.Context, req *menu.GetMenuReq) (*menu.GetMenuResp, error) {
	return &menu.GetMenuResp{Menus: []*menu.Menu{{Name: "abc"}}}, nil
}

func (MenuServer) GetMenuList(ctx context.Context, req *menu.GetMenuReq) (*menu.GetMenuResp, error) {
	//TODO implement me
	panic("implement me")
}
