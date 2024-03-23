package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-atreus/atreus-server/internal/server/router"
	"github.com/go-atreus/protocol/admin/auth"
	"github.com/go-atreus/protocol/admin/menu"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewHTTPServer(authApi *router.AuthApi, logger log.Logger, tp *tracesdk.TracerProvider,
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
		),
		//http.ResponseEncoder(DefaultResponseEncoder),
	}
	opts = append(opts, http.Address("0.0.0.0:8000"))
	httpSrv := http.NewServer(opts...)
	auth.RegisterAuthHTTPServer(r, nil, authApi)
	menu.RegisterMenuHTTPServer(r, nil, authApi)
	httpSrv.HandlePrefix("/", r)
	return httpSrv
}
