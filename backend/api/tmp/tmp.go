package main

import (
	"context"
	"log"
	pb "my-posting-site/backend/common/protobuf/golang/helloWorld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:5300", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("connection:", err)
	}
	defer conn.Close()

	client := pb.NewHelloWorldClient(conn)

	pong, err := client.Greeting(ctx, &pb.RequestHello{MessageHello: "Saveliy"})
	if err != nil {
		log.Fatal("err:", err)
	}
	log.Println("success:", pong)
}
