package server

import (
	"github.com/go-atreus/atreus-server/app/admin/internal/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGRPCServer, service.NewAuthServer,
	service.NewMenuServer,
	service.NewUserServer)
