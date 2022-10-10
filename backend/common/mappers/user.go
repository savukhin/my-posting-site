package mappers

import (
	"my-posting-site/backend/common/models"
	pbUser "my-posting-site/backend/common/protobuf/golang/user"
)

func UserToProfileResponse(req *models.User) *pbUser.ProfileResponse {
	return &pbUser.ProfileResponse{
		Id:        req.ID,
		Username:  req.Username,
		Success:   true,
		Error:     "",
		AvatarURL: "",
	}
}
