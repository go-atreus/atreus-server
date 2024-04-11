package biz

import (
	"context"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
)

type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*user.SysUserORM, error)
}
