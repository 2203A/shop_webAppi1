package main

import (
	_ "SX1/shop_webAppi/userserve/cmd/initialize"
	"SX1/shop_webAppi/userserve/handle"
	UserPDS "SX1/shop_webAppi/userserve/initnal/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 地址
	addr := "127.0.0.1:8081"
	// 1.监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	UserPDS.RegisterUserServer(s, &handle.User{})
	// 4.启动服务端
	s.Serve(listener)
}
