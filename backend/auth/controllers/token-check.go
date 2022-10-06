package auth_controllers

import (
	"context"
	"my-posting-site/backend/common/utils"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
)

func generateUserError(err string) (*pbAuth.UserResponse, error) {
	return &pbAuth.UserResponse{
		Success: false,
		Error:   err,
	}, nil
}

func (server *AuthServer) JWTCheck(ctx context.Context, req *pbAuth.JWTRequest) (*pbAuth.UserResponse, error) {
	token, err := utils.UnpackJWT(req.Token)
	if err != nil {
		return generateUserError("No user with this login/password")
	}

	return &pbAuth.UserResponse{
		Success:    true,
		Token:      req.Token,
		Expiration: token.TimeExp.String(),
		UserId:     int32(token.UserID),
	}, nil
}
