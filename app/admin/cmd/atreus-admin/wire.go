//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-atreus/atreus-server/app/admin/internal/conf"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-atreus/atreus-server/app/admin/internal/server"
	"github.com/go-atreus/atreus-server/app/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init application.
func initApp(log.Logger, *tracesdk.TracerProvider, *conf.Bootstrap, *conf.Auth) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
