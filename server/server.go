package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"

	pb "grpc-demo/api/user"
	"grpc-demo/register"
)

var (
	scheme      = "user"
	serviceName = "gprc.demo.user"
	addr        = "127.0.0.1:8012"
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
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			UserServerInterceptor,
			WorldInterceptor,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, &User{})
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func UserServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("你好")
	resp, err := handler(ctx, req)
	log.Println("再见")
	return resp, err
}

func WorldInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("你好，世界")
	resp, err := handler(ctx, req)
	log.Println("再见，世界")
	return resp, err
}

func init() {
	resolver.Register(register.NewResolverBuilder(scheme, serviceName, addr))
}
