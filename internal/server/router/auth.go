package router

import (
	context "context"
	"fmt"
	"github.com/go-atreus/atreus-server/internal/conf"
	"github.com/go-atreus/protocol/admin"
	"github.com/go-atreus/protocol/admin/auth"
	"github.com/go-atreus/protocol/admin/menu"
	"github.com/go-atreus/protocol/admin/user"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v5"
)

type AuthApi struct {
	*auth.AuthImpl
	*menu.MenuImpl
	*user.UserImpl
}

var _ auth.AuthHTTPServer = (*AuthApi)(nil)
var _ menu.MenuHTTPServer = (*AuthApi)(nil)
var _ user.UserHTTPServer = (*AuthApi)(nil)

func Client(keyProvider jwtv4.Keyfunc, opts ...jwt.Option) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if keyProvider == nil {
				return nil, jwt.ErrNeedTokenProvider
			}
			var fromContext jwtv4.Claims
			var ok bool
			if fromContext, ok = jwt.FromContext(ctx); !ok {
				fromContext = &jwtv4.MapClaims{}
			}
			token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, fromContext)
			key, err := keyProvider(token)
			if err != nil {
				return nil, jwt.ErrGetKey
			}
			tokenStr, err := token.SignedString(key)
			if err != nil {
				return nil, jwt.ErrSignToken
			}
			if clientContext, ok := transport.FromClientContext(ctx); ok {
				clientContext.RequestHeader().Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))
				return handler(ctx, req)
			}
			return nil, jwt.ErrWrongContext
		}
	}
}

func NewAuthApi(logger log.Logger, authConfig *conf.Auth, dis registry.Discovery) *AuthApi {
	address := "discovery:///" + admin.ServerName

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(address),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			tracing.Client(),
			logging.Client(logger),
			Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(authConfig.Key), nil
			},
				jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	)

	if err != nil {
		log.NewHelper(logger).Error(err)
		return nil
	}
	return &AuthApi{
		AuthImpl: auth.NewAuthImpl(conn),
		MenuImpl: menu.NewMenuImpl(conn),
		UserImpl: user.NewUserImpl(conn),
	}
}
