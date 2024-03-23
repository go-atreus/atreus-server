package server

import (
	"github.com/go-atreus/atreus-server/app/admin/internal/server/rpc"
	"github.com/google/wire"
)

var ServerProvider = wire.NewSet(NewGRPCServer, rpc.NewAuthServer, rpc.NewMenuServer)
