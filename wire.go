// +build wireinject

//+build wireinject

package main

import (
	"demo/internal/data"
	"demo/internal/server"
	"demo/internal/service"

	"github.com/MrZhangjicheng/kitdemo"
	"github.com/MrZhangjicheng/kitdemo/log"
	"github.com/google/wire"
)

//  上面写依赖，下面会自动匹配
func initApp(log.Logger) *kitdemo.App {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
