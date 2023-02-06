package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"protobuf/client/service"
)

func main() {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()

	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}

	productClient := service.NewProductServiceClient(conn)

	request := &service.ProductRequest{
		ProductId: 100,
	}

	stock, err := productClient.GetProductStock(context.Background(), request)

	if err != nil {
		log.Fatal("查询库存出错", err)
	}

	fmt.Println("查询成功", stock.ProductStock)
}
