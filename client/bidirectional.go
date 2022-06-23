package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-demo/api/bidirectional"
)

func main() {
	fmt.Println("client start...")
	conn, err := grpc.Dial("localhost:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewBidirectionalServiceClient(conn)

	user, err := client.GetUser(context.Background())
	if err != nil {
		log.Fatalf("fail to get user: %v", err)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		err := user.Send(&pb.GetUserRequest{UserId: strconv.Itoa(i)})
		if err != nil {
			log.Fatalf("fail to send: %v", err)
		}

		resp, err := user.Recv()
		if err != nil {
			log.Fatalf("fail to recv: %v", err)
		}
		fmt.Printf("%v\n", resp)
	}

	err = user.CloseSend()
	if err != nil {
		log.Fatalf("fail to close send: %v", err)
	}
}
