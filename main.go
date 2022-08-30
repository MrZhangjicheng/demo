package main

import (
	"os"

	"github.com/MrZhangjicheng/kitdemo"
	zapcli "github.com/MrZhangjicheng/kitdemo/contrib/log/zap"
	"github.com/MrZhangjicheng/kitdemo/log"
	"github.com/MrZhangjicheng/kitdemo/registry"
	"github.com/MrZhangjicheng/kitdemo/transport/grpc"
	"github.com/MrZhangjicheng/kitdemo/transport/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Name = "auth"
)

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, rr registry.Registar) *kitdemo.App {
	return kitdemo.New(
		kitdemo.Name(Name),
		kitdemo.Logger(logger),
		kitdemo.Server(
			gs,
			hs,
		),
		kitdemo.Registar(rr),
	)
}

func initZapLogger() *log.Logger {
	encoderCfg := zapcore.EncoderConfig{
		// MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
	zlogger := zap.New(core).WithOptions()
	logger := zapcli.NewLogger(zlogger)
	t := log.With(logger, "ts", log.DefaultTimestamp)
	t = log.With(t, "caller", log.DefaultCaller)

	return &t

}

func main() {
	// 初始化一份日志 zap
	logger := initZapLogger()

	app := initApp(*logger)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
