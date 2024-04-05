package service

import (
	"context"
	admin "github.com/go-atreus/atreus-server/app/admin/api"
	"github.com/go-atreus/atreus-server/app/admin/api/auth"
	"github.com/go-atreus/atreus-server/app/admin/api/menu"
	"github.com/go-atreus/atreus-server/app/admin/api/role"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
	"github.com/go-atreus/atreus-server/app/interface/internal/conf"
	"github.com/go-atreus/tools/http2rpc"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAdminInterface)

var _ user.UserServer = (*AdminInterface)(nil)

type AdminInterface struct {
	*user.UserImpl
	*auth.AuthImpl
	*menu.MenuImpl
	*role.RoleImpl

	log *log.Helper
}

func NewAdminInterface(logger log.Logger, authConfig *conf.Auth, dis registry.Discovery) *AdminInterface {

	address := "discovery:///" + admin.ServerName

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(address),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			tracing.Client(),
			logging.Client(logger),
			http2rpc.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(authConfig.Key), nil
			},
				jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	)

	if err != nil {
		log.NewHelper(logger).Error(err)
		return nil
	}

	return &AdminInterface{
		log:      log.NewHelper(log.With(logger, "module", "service/interface")),
		UserImpl: user.NewUserImpl(conn),
		AuthImpl: auth.NewAuthImpl(conn),
		MenuImpl: menu.NewMenuImpl(conn),
		RoleImpl: role.NewRoleImpl(conn),
	}
}
