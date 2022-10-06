package auth_controllers

import (
	pbPost "my-posting-site/common/protobuf/golang/post"
)

type PostServer struct {
	pbPost.UnimplementedPostingServer
}

func generateError(err string) (*pbPost.DefaultResponse, error) {
	return &pbPost.DefaultResponse{
		Success: false,
		Error:   err,
	}, nil
}

func generatePostError(err string) (*pbPost.PostResponse, error) {
	return &pbPost.PostResponse{
		Success: false,
		Error:   err,
	}, nil
}
