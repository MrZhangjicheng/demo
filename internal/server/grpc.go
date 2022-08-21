package server

import (
	v1 "demo/api/auth/service/v1"
	"demo/internal/service"

	"github.com/MrZhangjicheng/kitdemo/transport/grpc"
)

func NewGrpcServer(s *service.Auth) *grpc.Server {
	srv := grpc.NewServer()
	v1.RegisterAuthServer(srv, s)
	return srv
}
