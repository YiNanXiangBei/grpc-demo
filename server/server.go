package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "grpc-demo/api/user"
)

type User struct {
	pb.UnimplementedUserServiceServer
}

func (u *User) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// 这里可以做一些业务逻辑处理。为了演示方便，这里直接返回一个用户信息
	return &pb.GetUserResponse{
		UserId:    "1",
		Name:      "hedeqiang",
		Email:     "hedeqiang@88.com",
		CreatedAt: "2019-01-01",
		UpdatedAt: "2019-01-01",
	}, nil
}

func (u *User) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println(req)
	md1, _ := metadata.FromIncomingContext(ctx)
	md2, _ := metadata.FromOutgoingContext(ctx)
	fmt.Println("md1: ", md1)
	fmt.Println("md2: ", md2)
	// 创建用户，理论上应该是把用户信息写入数据库，这里为了演示方便直接返回一个用户信息
	return &pb.CreateUserResponse{
		UserId: "1",
		Name:   req.Name,
		Email:  req.Email,
	}, nil
}

func main() {
	fmt.Println("start server")
	listen, err := net.Listen("tcp", ":8012")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &User{})
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
