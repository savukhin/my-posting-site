package main

import (
	pb "my-posting-site/common/protobuf/golang/helloWorld"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterHelloWorldServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct {
	pb.HelloWorldServer
}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	output := string(request.Message)
	response = &pb.Response{
		Message: output,
	}

	return response, nil
}
