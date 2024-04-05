package service

import (
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
