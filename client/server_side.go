package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-demo/api/server_side"
)

func main() {
	fmt.Println("client start...")
	conn, err := grpc.Dial("localhost:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewServerSideClient(conn)

	user, err := client.GetUser(context.Background(), &pb.GetUserRequest{UserId: "1"})
	if err != nil {
		log.Fatalf("fail to get user: %v", err)
	}
	for {
		resp, err := user.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("fail to recv: %v", err)
		}
		fmt.Println(resp)
	}
}
