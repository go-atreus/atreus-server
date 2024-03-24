package server

import (
	"context"
	sj "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-atreus/atreus-server/internal/conf"
	"github.com/go-atreus/atreus-server/internal/server/router"
	"github.com/go-atreus/protocol/admin/auth"
	"github.com/go-atreus/protocol/admin/menu"
	"github.com/go-atreus/protocol/admin/user"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v5"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	shttp "net/http"
	"strings"
)

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
func NewHTTPServer(authApi *router.AuthApi, logger log.Logger,
	authConfig *conf.Auth,
	tp *tracesdk.TracerProvider,
) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	var opts = []http.ServerOption{
		http.Middleware(
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
						return &jwt2.MapClaims{}
					})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
		http.ResponseEncoder(JsonResponseEncoder),
	}
	opts = append(opts, http.Address("0.0.0.0:8000"))
	httpSrv := http.NewServer(opts...)
	auth.RegisterAuthHTTPServer(httpSrv, authApi)
	menu.RegisterMenuHTTPServer(httpSrv, authApi)
	user.RegisterUserHTTPServer(httpSrv, authApi)
	httpSrv.HandlePrefix("/", r)
	return httpSrv
}

// JsonResponseEncoder encodes the object to the HTTP response.
func JsonResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	codec, _ := http.CodecForRequest(r, "Accept")
	resBytes, err := codec.Marshal(v)
	if err != nil {
		return err
	}
	data, err := sj.Marshal(&Res{
		Code: 0, Message: "", Data: sj.RawMessage(resBytes),
	})
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
	w.WriteHeader(shttp.StatusOK)
	_, _ = w.Write(data)
	return nil
}

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
