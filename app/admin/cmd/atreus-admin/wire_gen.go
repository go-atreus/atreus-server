// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-atreus/atreus-server/app/admin/internal/server"
	"github.com/go-atreus/atreus-server/app/admin/internal/server/rpc"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init application.
func initApp(logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	authServer := rpc.NewAuthServer()
	menuServer := rpc.NewMenuServer(logger)
	grpcServer := server.NewGRPCServer(logger, authServer, menuServer)
	registrar := data.NewRegistrar()
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
	}, nil
}
