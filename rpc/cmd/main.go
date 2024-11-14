package main

import (
	"examProject/internal/tools/consul"
	"examProject/pkg/golbal"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//记录启动日志
	zap.S().Infow(
		"Started ApiServer：：：",
		zap.String("Host", "127.0.0.1"),
		zap.Int("Port", 8888),
	)
	consulClient := consul.NewConsulClient(golbal.ApiConfig.Consul.Host, golbal.ApiConfig.Consul.Port)
	// 注册到 consul  如果每次注册服务ID不一样 那么则需要进行注册前的判断
	consulClient.Register(golbal.ApiConfig.Server.Host, golbal.ApiConfig.Server.Port, golbal.ApiConfig.Server.Name, golbal.ApiConfig.Server.Tags, uuid.NewString())
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8888))
	if err != nil {
		fmt.Println()
		log.Fatalf("failed to listen: %v", err)
	}
	//建立grpc服务连接
	//var opts []grpc.ServerOption
	grpcServer := grpc.NewServer()
	//__.RegisterUserServer(grpcServer, &handle.User{})
	grpcServer.Serve(lis)

}
