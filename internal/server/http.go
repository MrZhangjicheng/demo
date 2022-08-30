package server

import (
	"demo/internal/service"

	"github.com/MrZhangjicheng/kitdemo/transport/http"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(s *service.Auth) *http.Server {
	r := gin.New()
	admin := r.Group("/admin")
	{
		admin.GET("/say", service.Say)
	}
	serveroption := http.WithHandler(r)
	srv := http.NewServer(serveroption)
	return srv
}
