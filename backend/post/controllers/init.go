package post_controllers

import (
	"log"
	grpc_clients "my-posting-site/backend/common/grpc"
)

var Client *grpc_clients.GRPCClients

func init() {
	opts := make([]grpc_clients.GRPCConnectionOpt, 0)
	opts = append(opts, grpc_clients.GRPCConnectionOpt{Url: "127.0.0.1:3400", ServerType: grpc_clients.Auth})

	var err error
	Client, err = grpc_clients.BuildGRPCClients(opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	Consume()
}
