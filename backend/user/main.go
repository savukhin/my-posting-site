package main

import (
	"fmt"
	grpc_clients "my-posting-site/backend/common/grpc"
	user_controllers "my-posting-site/backend/user/controllers"
)

func main() {
	fmt.Println("User service starting...")
	server, err := grpc_clients.NewServer().
		AddPort(":3100").
		AddUserServer(&user_controllers.UserServer{}).
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
