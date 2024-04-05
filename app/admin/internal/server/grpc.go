package server

import (
	"context"
	"github.com/go-atreus/atreus-server/app/admin/api/auth"
	"github.com/go-atreus/atreus-server/app/admin/api/dict"
	"github.com/go-atreus/atreus-server/app/admin/api/menu"
	"github.com/go-atreus/atreus-server/app/admin/api/organization"
	"github.com/go-atreus/atreus-server/app/admin/api/role"
	"github.com/go-atreus/atreus-server/app/admin/api/user"
	"github.com/go-atreus/atreus-server/app/admin/internal/conf"
	"github.com/go-atreus/atreus-server/app/admin/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

const TestKey = "abc"

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/Atreus.auth.Auth/getUserToken"] = struct{}{}
	whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
func NewGRPCServer(logger log.Logger, bc *conf.Bootstrap, authConfig *conf.Auth,
	userSvr *service.UserService,
	roleSvr *service.RoleService,
	dictSvr *service.DictService,
	orgSvr *service.OrganizationService,
	authSvr *service.AuthServer, menuSvr *service.MenuServer) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(authConfig.Key), nil
				},
					jwt.WithSigningMethod(jwt2.SigningMethodHS256),
					jwt.WithClaims(func() jwt2.Claims {
						return jwt2.MapClaims{}
					})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
	}
	c := bc.Server
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	svr := grpc.NewServer(opts...)
	auth.RegisterAuthServer(svr, authSvr)
	menu.RegisterMenuServer(svr, menuSvr)
	user.RegisterUserServer(svr, userSvr)
	role.RegisterRoleServer(svr, roleSvr)
	organization.RegisterOrganizationServer(svr, orgSvr)
	dict.RegisterDictServer(svr, dictSvr)
	return svr
}
