package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-demo/api/client_side"
)

type ClientSide struct {
	Name string
	pb.UnimplementedClientSideServer
}

func (c *ClientSide) GetUser(stream pb.ClientSide_GetUserServer) error {
	for {
		var res pb.GetUserResponse
		res.UserId = "123"
		res.Name = "hedeqiang"
		res.Email = "laravel_code@163.com"

		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&res)
		}
		if err != nil {
			log.Fatalf("%v.GetUser(_) = _, %v", c, err)
		}
		fmt.Println(recv.GetUserId())

	}
}

func main() {
	fmt.Println("start server")

	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterClientSideServer(s, &ClientSide{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
