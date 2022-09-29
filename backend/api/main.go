package main

import (
	"fmt"
	pb "my-posting-site/common/protobuf/golang/helloWorld"
	"net"

	// "github.com/gorilla/handlers"
	// "github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	defer listener.Close()

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterHelloWorldServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct {
	pb.UnimplementedHelloWorldServer
}

func (s *server) Greeting(c context.Context, request *pb.RequestHello) (response *pb.ResponseHello, err error) {
	fmt.Println("Got greeting request")

	output := "Hello world! Hello, " + string(request.MessageHello)

	return &pb.ResponseHello{
		MessageHello: output,
	}, nil
}
