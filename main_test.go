package main

import (
	"context"
	"io/ioutil"
	"log"
	"testing"

	"github.com/hashicorp/consul/api"

	"github.com/MrZhangjicheng/kitdemo/contrib/registry/consul"
	"github.com/MrZhangjicheng/kitdemo/transport/http"
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
	// srvs, _ := r.FindService(context.Background(), "auth")
	// if len(srvs) == 0 {
	// 	fmt.Println("没有服务")
	// }
	// addr := srvs[0].Endpoints[0][7:]
	// addr := "192.168.60.182:1038"

	// conn, _ := grpc.Dial(addr, grpc.WithInsecure())
	// conn, _ := grpc.Dial(context.Background(), grpc.WithEndPoint("discovery:///auth"), grpc.Withdiscovery(r))
	// conn, _ := grpc.Dial(context.Background(), grpc.WithEndPoint("192.168.60.182:12126"), grpc.Withdiscovery(r))
	// cAuth := v1.NewAuthClient(conn)
	// resp, errAuth := cAuth.Login(context.Background(), &v1.LoginReq{
	// 	Password: "2113",
	// 	Email:    "sssser",
	// })
	// if errAuth != nil {
	// 	fmt.Println(errAuth)
	// 	log.Fatal(errAuth)
	// }

	// log.Fatal(resp)

	// httpConn, _ := http.Newclient(context.Background(), http.WithEndpoint("192.168.60.182:12127"), http.Withdiscovery(r))
	httpConn, _ := http.Newclient(context.Background(), http.WithEndpoint("discovery:///auth"), http.Withdiscovery(r))

	resp, _ := httpConn.Do("GET", "/admin/say")
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("%s", data)

}
