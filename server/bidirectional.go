package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-demo/api/bidirectional"
)

type BidirectionalService struct {
	pb.UnimplementedBidirectionalServiceServer
}

func (b *BidirectionalService) GetUser(stream pb.BidirectionalService_GetUserServer) (err error) {
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Printf("stream.Recv error: %v", err)
			return err
		}
		log.Printf("stream.Recv req: %v", req)
		err = stream.Send(&pb.GetUserResponse{
			UserId: "123",
			Name:   "hedeqiang",
			Email:  "laravel_code@163.com",
		})
		if err != nil {
			log.Printf("stream.Send error: %v", err)
			return err
		}
	}

}

func main() {
	fmt.Println("start server")

	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBidirectionalServiceServer(s, &BidirectionalService{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
