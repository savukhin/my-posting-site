package user_controllers

import (
	"context"

	"my-posting-site/backend/common/mappers"
	"my-posting-site/backend/common/models"
	pbUser "my-posting-site/backend/common/protobuf/golang/user"
)

type UserServer struct {
	pbUser.UnimplementedUserServer
}

func generateError(err string) (*pbUser.ProfileResponse, error) {
	return &pbUser.ProfileResponse{
		Success: false,
		Error:   err,
	}, nil
}

func (server *UserServer) GetProfile(ctx context.Context, req *pbUser.ProfileRequest) (*pbUser.ProfileResponse, error) {
	user, err := models.GetUserById(req.Id)
	if err != nil {
		return generateError("No user with this id")
	}

	return mappers.UserToProfileResponse(user), nil
}
