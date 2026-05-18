package main

import (
	"captcha/handler"
	pb "captcha/proto/captcha"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "captcha"
	version = "latest"
)

func main() {
	//配置consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)
	srv.Init()

	// Register handler
	pb.RegisterCaptchaHandler(srv.Server(), new(handler.Captcha))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
