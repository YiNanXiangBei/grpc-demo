package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-demo/api/client_side"
)

func main() {
	fmt.Println("client start...")
	conn, err := grpc.Dial("localhost:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewClientSideClient(conn)

	user, err := client.GetUser(context.Background())
	if err != nil {
		return
	}
	if err != nil {
		log.Fatalf("fail to get user: %v", err)
	}
	i := 0
	for {
		if i > 10 {
			break
		}
		i++
		time.Sleep(time.Second)
		err := user.Send(&pb.GetUserRequest{
			UserId: strconv.Itoa(i),
		})
		fmt.Println("send:", i)
		if err != nil {
			log.Fatalf("fail to send: %v", err)
		}
	}

	resp, err := user.CloseAndRecv()
	if err != nil {
		log.Fatalf("fail to close: %v", err)
	}
	fmt.Println(resp)
}
