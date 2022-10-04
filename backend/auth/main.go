package main

import (
	"fmt"
	auth_controllers "my-posting-site/backend/auth/controllers"
	grpc_clients "my-posting-site/backend/common/grpc"
)

func main() {
	fmt.Println("Auth service starting...")
	server, err := grpc_clients.NewServer().
		AddPort(":5300").
		AddAuthServer(&auth_controllers.AuthServer{}).
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
