package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/jili/pkg-practice/grpc/rpcpb/rpcpb1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	caPath     string
	client_crt string
	client_key string
)

func init() {
	fileDir, _ := os.Getwd()
	caPath = filepath.Join(fileDir, "grpc/ca_cert.pem")
	client_crt = filepath.Join(fileDir, "grpc/client_cert.pem")
	client_key = filepath.Join(fileDir, "grpc/client_key.pem")
}

func main() {
	//从输入的证书文件中为客户端构造TLS凭证
	cert, err := tls.LoadX509KeyPair(client_crt, client_key)
	if err != nil {
		log.Fatalf("Failed to LoadX509KeyPair %v", err)
	}
	caPool := x509.NewCertPool()
	bytes, err := os.ReadFile(caPath)
	if err != nil {
		log.Fatalf("Failed to os.ReadFile %v", err)
	}

	if ok := caPool.AppendCertsFromPEM(bytes); !ok {
		log.Fatalf("Failed to AppendCertsFromPEM %v", err)
	}
	tlsConfig := &tls.Config{
		ServerName:   "www.example.com", // NOTE: this is required!
		Certificates: []tls.Certificate{cert},
		RootCAs:      caPool,
	}

	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	clientInfo := rpcpb1.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := clientInfo.SayHello(ctx, &rpcpb1.HelloRequest{Name: "name21431"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
