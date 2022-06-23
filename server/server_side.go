package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "grpc-demo/api/server_side"
)

type ServerSide struct {
	pb.UnimplementedServerSideServer
}

func (s *ServerSide) GetUser(req *pb.GetUserRequest, stream pb.ServerSide_GetUserServer) error {
	fmt.Println("GetUser")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		err := stream.Send(&pb.GetUserResponse{
			UserId: fmt.Sprintf("%d", i),
			Name:   "name" + fmt.Sprintf("%d", i),
			Email:  fmt.Sprintf("laravel_code@163.com"),
		})
		if err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, req, err)
		}
	}
	return nil
}

func main() {
	fmt.Println("start server")

	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServerSideServer(s, &ServerSide{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
