package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/grpclog"

	"github.com/ping-localhost/grpc-bug-report/server/logging"
)

const (
	port = ":50051"
)

type HelloService struct {
	pb.UnimplementedGreeterServer
}

func (s *HelloService) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(fmt.Errorf("failed to listen: %v", err))
	}

	certificate, err := tls.LoadX509KeyPair("../ssl/certificate.crt", "../ssl/certificate.key")
	if err != nil {
		panic("unable to load X509 key pair")
	}

	logger := logging.NewLogger()
	grpclog.SetLoggerV2(logging.NewGRPCLogger(logger))

	// Create the gRPC server
	server := grpc.NewServer([]grpc.ServerOption{
		// Sets the message size in bytes the server can receive
		grpc.MaxRecvMsgSize(int(10485760)),
		// Set TLS options
		grpc.Creds(credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{certificate},
		})),
	}...)

	pb.RegisterGreeterServer(server, &HelloService{})

	log.Print("Starting server....")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
