package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	pb "grpc-demo/api/user"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             1 * time.Second,  // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func main() {
	fmt.Println("client start")
	// 通过设置超时时间，然后抓包之后可以明确看到，超时时间和链接的复用存在着一定关系
	conn, err := grpc.Dial("127.0.0.1:8012", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	md := map[string][]string{
		"laowang": {"niupi"},
	}
	con := metadata.NewOutgoingContext(context.Background(), md)
	//con := metadata.NewIncomingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(con, 60*time.Second)
	defer cancel()
	md1, _ := metadata.FromIncomingContext(con)
	fmt.Println(md1)
	md2, _ := metadata.FromOutgoingContext(con)
	fmt.Println(md2)

	// incoming 数据只有 incoming 才能接收；outgoing 数据也是只有 outgoing 才能接收
	// 写入到outgoing 的数据可以传到外部去，写入到 incoming 的数据不可以传入到外部
	// 但是外部只能使用 incoming 才能接收到数据

	// 创建用户
	createUserRequest := &pb.CreateUserRequest{
		Name:     "zhangsan",
		Email:    "hedeqiang@88.com",
		Password: "123456",
	}
	user, err := c.CreateUser(ctx, createUserRequest)
	if err != nil {
		log.Fatalf("create user failed: %v", err)
	}

	log.Printf("create user success: %v\n", user)

	time.Sleep(30 * time.Second)

	// 获取用户信息 GetUser
	getUserRequest := &pb.GetUserRequest{
		UserId: "1",
	}
	getUser, err := c.GetUser(ctx, getUserRequest)
	if err != nil {
		log.Fatalf("get user failed: %v", err)
	}
	log.Printf("get user success: %v\n", getUser)

}
