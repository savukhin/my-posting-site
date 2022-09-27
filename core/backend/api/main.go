package main

import (
	// "context"

	"fmt"

	// "os"
	pb "core/common/protobuf/my-posting-site/api/helloworld"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/grpclog"
	"core/backend/api/tmp"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/grpclog"
	// "tmp"
	// "backend/api/tmp"
	// "tmp"
)

func main() {
	// opts := []grpc.DialOption{
	// 	grpc.WithInsecure(),
	// }
	// args := os.Args
	// conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	// if err != nil {
	// 	grpclog.Fatalf("fail to dial: %v", err)
	// }

	// defer conn.Close()

	// client := pb.NewReverseClient(conn)
	// request := &pb.Request{
	// 	Message: args[1],
	// }
	// response, err := client.Do(context.Background(), request)

	// if err != nil {
	// 	grpclog.Fatalf("fail to dial: %v", err)
	// }

	req := &pb.Request{}
	if req != nil {
		fmt.Println("req is not nil")
	}
	fmt.Println("Started")
	fmt.Println(tmp.Echo("asdf"))
}
