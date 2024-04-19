package service

import (
	"context"
	"github.com/go-atreus/atreus-server/app/admin/api/role"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

var _ role.RoleServer = (*RoleService)(nil)

func NewRoleService(logger log.Logger, data *data.Data) *RoleService {
	return &RoleService{
		RoleDefaultServer: role.RoleDefaultServer{DB: data.ORM},
	}
}

type RoleService struct {
	role.RoleDefaultServer
}

func (m *RoleService) RolePermissions(ctx context.Context, in *role.SysRole) (res *role.RoleMenu, err error) {
	out := &role.RoleMenu{}
	var ret *role.SysRoleORM
	if err = m.DB.Where("code = ?", in.Code).Preload("Menu").Find(&ret).Error; err != nil {
		return
	}
	for _, menu := range ret.Menu {
		out.Menus = append(out.Menus, menu.Id)
	}
	return out, nil
}
