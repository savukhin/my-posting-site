package main

import (
	"fmt"
	grpc_clients "my-posting-site/backend/common/grpc"
	post_controllers "my-posting-site/backend/post/controllers"
)

func main() {
	fmt.Println("Auth service starting...")
	server, err := grpc_clients.NewServer().
		AddPort(":3400").
		AddPostingServer(&post_controllers.PostServer{}).
		Build()

	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()

	fmt.Println("Starting to serve")
	server.Serve()
	fmt.Println("Stop serving")
}
