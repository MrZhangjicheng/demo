package data

import (
	consul "github.com/MrZhangjicheng/kitdemo/contrib/registry/consul"
	"github.com/MrZhangjicheng/kitdemo/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewRegister)

type Data struct {
	db *gorm.DB
}

func NewRegister() registry.Registar {
	c := consulAPI.DefaultConfig()
	c.Address = "127.0.0.1:8500"
	c.Scheme = "http"
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}
