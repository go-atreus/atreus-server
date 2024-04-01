package server

import (
	"github.com/go-atreus/atreus-server/app/api/internal/server/router"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, router.NewAuthApi)
