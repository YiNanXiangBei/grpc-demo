package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-demo/api/simple"
)

func main() {
	fmt.Println("client start...")
	conn, err := grpc.Dial("localhost:8989", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSimpleServiceClient(conn)

	response, err := client.Get(context.Background(), &pb.SimpleRequest{Name: "hedeqiang"})
	if err != nil {
		log.Fatalf("fail to call: %v", err)
	}
	fmt.Println(response.GetMessage())
}
