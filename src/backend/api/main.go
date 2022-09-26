package main

import (
    // "context"
    "fmt"
    // "os"    
	// pb "common/protobuf/my-posting-site/api/helloworld"
    // "google.golang.org/grpc"
    // "google.golang.org/grpc/grpclog"
    "backend/api/tmp"
    // "tmp"

    // "backend/api/tmp"
    // "tmp"
)

func main() {
	// opts := []grpc.DialOption{
    //     grpc.WithInsecure(),
    // }

    // if (opts == nil) {
    //     fmt.Println("No opts")
    // }

    fmt.Println("Started")
    fmt.Println(tmp.Echo("Echo"))
}