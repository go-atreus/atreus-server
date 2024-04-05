package service

import (
	context "context"
	"github.com/go-atreus/atreus-server/app/admin/api/menu"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

var _ menu.MenuServer = (*MenuServer)(nil)

type MenuServer struct {
	menu.MenuDefaultServer
}

func NewMenuServer(logger log.Logger, data *data.Data) *MenuServer {
	return &MenuServer{
		MenuDefaultServer: menu.MenuDefaultServer{DB: data.ORM},
	}

}

func (MenuServer) GetMenu(ctx context.Context, req *menu.GetMenuReq) (*menu.GetMenuResp, error) {
	return &menu.GetMenuResp{Menus: []*menu.Menu{{Name: "abc"}}}, nil
}
