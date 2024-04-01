package server

import (
	"github.com/go-atreus/atreus-server/app/admin/internal/server/rpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGRPCServer, rpc.NewAuthServer,
	rpc.NewMenuServer,
	rpc.NewUserServer)
