package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	service "protobuf/server/service"
)

func main() {
	//创建一个GRPC服务
	rpcServer := grpc.NewServer()
	//注册接口
	service.RegisterProductServiceServer(rpcServer, service.ProductService)

	//自定义TCP协议
	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听出错", err)
	}

	err = rpcServer.Serve(listen)
	if err != nil {
		log.Fatal("启动服务出错", err)
	}
}
