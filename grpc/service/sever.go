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
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	caPath     string
	server_crt string
	server_key string
)

func init() {
	fileDir, _ := os.Getwd()
	caPath = filepath.Join(fileDir, "grpc/client_ca_cert.pem")
	server_crt = filepath.Join(fileDir, "grpc/server_cert.pem")
	server_key = filepath.Join(fileDir, "grpc/server_key.pem")
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	rpcpb1.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *rpcpb1.HelloRequest) (*rpcpb1.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &rpcpb1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// 我们在用go 1.15版本以上，用gRPC通过TLS建立安全连接时，会出现证书报错问题：
// panic: rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate
// is not valid for any names, but wanted to match localhost"
func main() {

	certificate, err := tls.LoadX509KeyPair(server_crt, server_key)
	if err != nil {
		log.Panicf("could not load server key pair: %s", err)
	}

	certPool := x509.NewCertPool()
	bytes, err := os.ReadFile(caPath)
	if err != nil {
		log.Panicf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(bytes); !ok {
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

	// 双向tls
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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
