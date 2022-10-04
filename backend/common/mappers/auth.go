package mappers

import (
	"errors"
	"my-posting-site/backend/common/models"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
)

// func LoginRequestToUser(req *pbAuth.LoginRequest) (*models.User, error) {

// }

func RegisterRequestToUser(req *pbAuth.RegisterRequest) (*models.User, error) {
	if req.Password != req.Password2 {
		return nil, errors.New("passwords not equal")
	}

	return models.CreateUser(req.Username, req.Email, req.Password)
}
