package main

import (
	"examProject/cmd/userSrv/handel"
	"examProject/cmd/userSrv/proto/UserSrv"
	_ "examProject/pkg/initialization"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 地址
	addr := "127.0.0.1:8888"
	// 1.监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	UserSrv.RegisterUserServer(s, &handel.User{})
	s.Serve(listener)
}
