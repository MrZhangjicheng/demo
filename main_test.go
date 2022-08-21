package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	v1 "demo/api/auth/service/v1"

	"github.com/hashicorp/consul/api"

	"github.com/MrZhangjicheng/kitdemo/contrib/registry/consul"
	"google.golang.org/grpc"
)

func Test(t *testing.T) {
	c := api.DefaultConfig()
	c.Address = "127.0.0.1:8500"
	c.Scheme = "http"
	cli, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	srvs, _ := r.FindService(context.Background(), "auth")
	if len(srvs) == 0 {
		fmt.Println("没有服务")
	}
	addr := srvs[0].Endpoints[0][7:]
	// addr := "192.168.60.182:1038"
	conn, _ := grpc.Dial(addr, grpc.WithInsecure())
	cAuth := v1.NewAuthClient(conn)
	resp, errAuth := cAuth.Login(context.Background(), &v1.LoginReq{
		Password: "2113",
		Email:    "sssser",
	})
	if errAuth != nil {
		fmt.Println(errAuth)
		log.Fatal(errAuth)
	}

	log.Fatal(resp)

}
