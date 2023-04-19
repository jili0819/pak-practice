package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/jili/pkg-practice/grpc/rpcpb/rpcpb1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	tlsDir = "./tls-config"

	ca         = tlsDir + "/ca.crt"
	server_crt = tlsDir + "/server.crt"
	server_key = tlsDir + "/server.key"
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
	certificate, err := tls.LoadX509KeyPair(server_crt, server_key)
	if err != nil {
		log.Panicf("could not load server key pair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Panicf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Panic("failed to append client certs")
	}
	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	s := grpc.NewServer(grpc.Creds(creds))
	rpcpb1.RegisterGreeterServer(s, &server{})

	// grpcui -plaintext 127.0.0.1:8090
	reflection.Register(s) // grpcui本地测试调用grpc服务
	go func() {
		if err = s.Serve(lis); err != nil {
			errChan <- err
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	defer func() {
		s.GracefulStop()
	}()
	select {
	case <-errChan:
	case <-stopChan:
	}
}
