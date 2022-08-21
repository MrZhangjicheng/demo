// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"demo/internal/data"
	"demo/internal/server"
	"demo/internal/service"
	"github.com/MrZhangjicheng/kitdemo"
	"github.com/MrZhangjicheng/kitdemo/log"
)

// Injectors from wire.go:

//  上面写依赖，下面会自动匹配
func initApp(logger log.Logger) *kitdemo.App {
	auth := service.NewAuth(logger)
	grpcServer := server.NewGrpcServer(auth)
	registar := data.NewRegister()
	app := newApp(logger, grpcServer, registar)
	return app
}
