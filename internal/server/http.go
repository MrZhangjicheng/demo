package server

import (
	"demo/internal/service"

	"github.com/MrZhangjicheng/kitdemo/transport/http"
)

func NewHTTPServer(s *service.Auth) *http.Server {
	srv := http.NewServer()
	return srv
}
