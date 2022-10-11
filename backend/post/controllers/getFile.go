package post_controllers

import (
	"context"
	pbPost "my-posting-site/backend/common/protobuf/golang/post"
	"os"
)

func generateFileError(err string) (*pbPost.GetFileResponse, error) {
	return &pbPost.GetFileResponse{
		Success: false,
		Error:   err,
	}, nil
}

func (server *PostServer) GetFile(ctx context.Context, req *pbPost.GetFileRequest) (*pbPost.GetFileResponse, error) {
	file, err := os.ReadFile(req.Path)
	if err != nil {
		return generateFileError(err.Error())
	}

	return &pbPost.GetFileResponse{
		Success: true,
		File:    file,
	}, nil
}
