package main

import (
	"context"
	user "first_demo/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	userServiceClient := user.NewUserServiceClient(conn)
	userResponse, _ := userServiceClient.GetUserInfo(context.Background(), &user.UserRequest{
		Name: "test_one",
	})
	fmt.Println(userResponse, userResponse.Age, userResponse.Gender)
	userResponse, _ = userServiceClient.GetUserInfo(context.Background(), &user.UserRequest{
		Name: "test_two",
	})
	fmt.Println(userResponse, userResponse.Age, userResponse.Gender)
}
