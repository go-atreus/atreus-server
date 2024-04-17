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

func (m *RoleService) RolePermissions(ctx context.Context, in *role.SysRole) (*role.ListRoleResp, error) {
	out := &role.ListRoleResp{}
	return out, nil
}
