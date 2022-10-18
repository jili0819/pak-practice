package main

import (
	"context"
	"fmt"
	"github.com/jili/pkg-practice/grpc/rpcpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	rpcpb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *rpcpb.HelloRequest) (*rpcpb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &rpcpb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 80))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpcpb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
