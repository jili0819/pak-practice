package main

import (
	"context"
	"fmt"
	"github.com/jili/pkg-practice/grpc/rpcpb/rpcpb1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	rpcpb1.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *rpcpb1.HelloRequest) (*rpcpb1.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &rpcpb1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpcpb1.RegisterGreeterServer(s, &server{})
	reflection.Register(s) // grpcui本地测试调用grpc服务
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
