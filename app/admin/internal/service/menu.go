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

func (m *MenuServer) GetMenu(ctx context.Context, req *menu.GetMenuReq) (res *menu.GetMenuResp, err error) {
	sysMenu, err := m.MenuDefaultServer.ListSysMenu(ctx, req)
	if err != nil {
		return
	}
	treeMap := make(map[int32][]*menu.SysMenu)
	for _, v := range sysMenu.Results {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(menus[i], treeMap)
	}
	return &menu.GetMenuResp{Routes: menus}, nil
}

func getChildrenList(menu *menu.SysMenu, treeMap map[int32][]*menu.SysMenu) (err error) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(menu.Children[i], treeMap)
	}
	return err
}
