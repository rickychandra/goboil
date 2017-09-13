package main

import (
	"github.com/rickychandra/goboil/common/config"
	"github.com/rickychandra/goboil/common/logger"
	"github.com/rickychandra/goboil/common/server"
	"github.com/rickychandra/goboil/common/stringutil"
	"github.com/rickychandra/goboil/services/myservice/api"
	pb2 "github.com/rickychandra/goboil/shared/pb"
	"google.golang.org/grpc"
)

func main() {
	logger.Init()

	lis := server.CreateTCPListener("", config.Get(config.MyServicePort))
	grpcServer := grpc.NewServer()
	pb2.RegisterMyServiceServer(grpcServer, api.NewMyServiceServer())
	logger.Info(nil, stringutil.Sprintf(
		"Service listening at %v", config.Get(config.MyServicePort)))
	grpcServer.Serve(lis)
}
