package dto

import (
	pbAuth "my-posting-site/backend/common/protobuf/golang/auth"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password2"`
}

type User struct {
	HasError bool   `json:"has_error"`
	Msg      string `json:"msg"`
	ID       int    `json:"id"`
}

func (data *Login) ToProtobuf() *pbAuth.LoginRequest {
	return &pbAuth.LoginRequest{
		Username: data.Username,
		Password: data.Password,
	}
}

func (data *Register) ToProtobuf() *pbAuth.RegisterRequest {
	return &pbAuth.RegisterRequest{
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		Password2: data.PasswordRepeat,
	}
}
