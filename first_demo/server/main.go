package main

import (
	"context"
	user "first_demo/pb"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Gender string
	Hobby  string
}

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (uss UserServiceServer) GetUserInfo(ctx context.Context, ur *user.UserRequest) (*user.UserResponse, error) {
	s := ur.GetName()
	db, err := initMysql()
	if err != nil {
		log.Fatalln("打开数据库失败")
	}

	//自动迁移，根据传入的指针建表
	createTable(db)

	var userResult User
	db.Where("name = ?", s).First(&User{}).Scan(&userResult)

	return &user.UserResponse{
		Age:    userResult.Age,
		Gender: userResult.Gender,
	}, nil
}

func main() {
	//监听端口
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("监听端口失败", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, UserServiceServer{})
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalln(err)
	}
}

var dsn string = "root:123456@tcp(127.0.0.1:3306)/dbtest1?charset=utf8mb4&parseTime=True&loc=Local"

func initMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func createTable(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&User{})
	if err != nil {
		log.Fatalln(err)
	}
}
