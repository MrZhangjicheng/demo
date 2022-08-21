package service

import (
	pb "demo/api/auth/service/v1"

	"github.com/MrZhangjicheng/kitdemo/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuth)

// 该层是引用
type Auth struct {
	pb.UnimplementedAuthServer

	log *log.Helper
}

func NewAuth(logger log.Logger) *Auth {
	return &Auth{
		log: log.NewHelper(logger),
	}
}
