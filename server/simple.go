package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-demo/api/simple"
)

type Simple struct {
	Name string
	pb.UnimplementedSimpleServiceServer
}

func (s *Simple) Get(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	name := req.GetName()
	return &pb.SimpleResponse{
		Message: "Hello " + name,
	}, nil
}

func main() {
	fmt.Println("start server")

	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSimpleServiceServer(s, &Simple{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
