// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-atreus/atreus-server/app/admin/internal/conf"
	"github.com/go-atreus/atreus-server/app/admin/internal/data"
	"github.com/go-atreus/atreus-server/app/admin/internal/server"
	"github.com/go-atreus/atreus-server/app/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init application.
func initApp(logger log.Logger, tracerProvider *trace.TracerProvider, bootstrap *conf.Bootstrap, auth *conf.Auth, confData *conf.Data) (*kratos.App, func(), error) {
	db := data.NewOrm(logger, bootstrap)
	cmdable := data.NewRedisCmd(confData, logger)
	dataData, cleanup, err := data.NewData(db, cmdable, logger)
	if err != nil {
		return nil, nil, err
	}
	userService := service.NewUserServer(logger, dataData)
	roleService := service.NewRoleService(logger, dataData)
	dictService := service.NewDictService(logger, dataData)
	organizationService := service.NewOrganization(logger, dataData)
	userRepo := data.NewUserRepo(dataData, logger)
	authServer := service.NewAuthServer(auth, userRepo)
	menuServer := service.NewMenuServer(logger, dataData)
	grpcServer := server.NewGRPCServer(logger, bootstrap, auth, userService, roleService, dictService, organizationService, authServer, menuServer)
	registrar := data.NewRegistrar()
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
