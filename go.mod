module demo

go 1.16

require (
	github.com/MrZhangjicheng/kitdemo v0.0.0
	github.com/MrZhangjicheng/kitdemo/contrib/log/zap v0.0.0-00010101000000-000000000000
	github.com/MrZhangjicheng/kitdemo/contrib/registry/consul v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.8.1
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/api v1.14.0
	go.uber.org/zap v1.22.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gorm.io/gorm v1.23.8
)

replace github.com/MrZhangjicheng/kitdemo => ../kitdemo

replace github.com/MrZhangjicheng/kitdemo/contrib/registry/consul => ../kitdemo/contrib/registry/consul

replace github.com/MrZhangjicheng/kitdemo/contrib/log/zap => ../kitdemo/contrib/log/zap
