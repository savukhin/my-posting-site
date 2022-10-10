package auth_controllers

import (
	"context"
	"my-posting-site/backend/common/mappers"
	"my-posting-site/backend/common/models"
	pbAuth "my-posting-site/backend/common/protobuf/golang/auth"
	"my-posting-site/backend/common/utils"
)

type AuthServer struct {
	pbAuth.UnimplementedAuthenticationServer
}

func generateError(err string) (*pbAuth.DefaultResponse, error) {
	return &pbAuth.DefaultResponse{
		Success: false,
		Error:   err,
	}, nil
}

func (server *AuthServer) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.DefaultResponse, error) {
	user, err := models.GetUserByUsername(req.Username)
	if err != nil {
		return generateError("No user with this login/password")
	}

	if err := user.ComparePassword(req.Password); err != nil {
		return generateError("No user with this login/password")
	}

	return &pbAuth.DefaultResponse{
		Success: true,
		Token:   utils.GenerateJWT(user),
	}, nil
}

func (server *AuthServer) Register(ctx context.Context, req *pbAuth.RegisterRequest) (*pbAuth.DefaultResponse, error) {
	user, err := mappers.RegisterRequestToUser(req)

	if err := user.IsNotExists(); err != nil {
		return generateError(err.Error())
	}

	if err != nil {
		return generateError(err.Error())
	}

	_, err = user.Save()
	if err != nil {
		return generateError(err.Error())
	}

	return &pbAuth.DefaultResponse{
		Success: true,
		Token:   utils.GenerateJWT(user),
	}, nil
}
