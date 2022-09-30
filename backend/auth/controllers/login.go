package auth_controllers

import (
	"context"
	"fmt"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
)

type AuthServer struct {
	pbAuth.UnimplementedAuthenticationServer
}

func (server *AuthServer) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.DefaultResponse, error) {
	fmt.Println("Got login request")

	output := "Hello world! Hello, " + string(req.Username)

	return &pbAuth.DefaultResponse{
		Token: output,
	}, nil
}
