package server

import (
	"github.com/go-atreus/atreus-server/app/admin/internal/server/rpc"
	"github.com/go-atreus/protocol/admin/auth"
	"github.com/go-atreus/protocol/admin/menu"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(logger log.Logger, authSvr *rpc.AuthServer, menuSvr *rpc.MenuServer) *grpc.Server {

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
		),
	}
	//opts = append(opts, grpc.Network("0.0.0.0"))
	opts = append(opts, grpc.Address("127.0.0.1:8081"))
	svr := grpc.NewServer(opts...)
	auth.RegisterAuthServer(svr, authSvr)
	menu.RegisterMenuServer(svr, menuSvr)
	return svr
}
