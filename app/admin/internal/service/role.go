package service

import (
	"context"
	"github.com/go-atreus/atreus-server/app/admin/api/menu"
	"github.com/go-atreus/atreus-server/app/admin/api/role"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (m *RoleService) GetRole(ctx context.Context, sysRole *role.SysRole) (*role.SysRole, error) {
	//TODO implement me
	panic("implement me")
}

func (m *RoleService) RoleSelect(ctx context.Context, empty *emptypb.Empty) (res *role.RoleSelectResp, err error) {
	var roles []*role.SysRoleORM
	if err = m.DB.Find(&roles).Error; err != nil {
		return
	}
	res = &role.RoleSelectResp{}
	for _, r := range roles {
		res.Options = append(res.Options, &role.RoleSelectOption{Name: r.Name, Value: r.Code})
	}
	return
}

func (m *RoleService) RolePermissions(ctx context.Context, in *role.SysRole) (res *role.RoleMenu, err error) {
	out := &role.RoleMenu{}
	var ret *role.SysRoleORM
	if err = m.DB.Where("code = ?", in.Code).Preload("Menu").Find(&ret).Error; err != nil {
		return
	}
	for _, mm := range ret.Menu {
		out.Menus = append(out.Menus, mm.Id)
	}
	return out, nil
}

func (m *RoleService) RolePermissionsPut(ctx context.Context, in *role.RolePermReq) (res *role.RoleMenu, err error) {
	out := &role.RoleMenu{}
	var menus []*menu.SysMenuORM
	if err = m.DB.Where("id in (?)", in.GetIds()).Find(&menus).Error; err != nil {
		return
	}
	var roleOrm *role.SysRoleORM
	if err = m.DB.Model(&role.SysRoleORM{}).Where("code = ?", in.Code).First(&roleOrm).Error; err != nil {
		return
	}
	if err = m.DB.Model(&roleOrm).Association("Menu").Replace(menus); err != nil {
		return
	}
	return out, nil
}
