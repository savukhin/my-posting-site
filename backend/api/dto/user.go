package dto

import (
	pbUser "my-posting-site/backend/common/protobuf/golang/user"
)

type UserReq struct {
	ID int32 `json:"user_id"`
}

type Profile struct {
	Has_err   bool   `json:"has_error"`
	Msg       string `json:"msg"`
	Id        int32  `json:"user_id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

func (data *UserReq) ToProtobuf() *pbUser.ProfileRequest {
	return &pbUser.ProfileRequest{
		Id: data.ID,
	}
}

func ProfileFromProtobuf(src *pbUser.ProfileResponse) *Profile {
	profile := &Profile{
		Id:        src.Id,
		Username:  src.Username,
		AvatarURL: src.AvatarURL,
	}

	if src.Success {
		profile.Has_err = false
		profile.Msg = "Succesfully got profile"
	} else {
		profile.Has_err = true
		profile.Msg = src.Error
	}

	return profile
}
