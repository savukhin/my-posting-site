package post_controllers

import (
	"context"
	"my-posting-site/backend/common/mappers"
	"my-posting-site/backend/common/models"
	pbPost "my-posting-site/common/protobuf/golang/post"
)

func generatePostError(err string) (*pbPost.PostResponse, error) {
	return &pbPost.PostResponse{
		Success: false,
		Error:   err,
	}, nil
}

func (server *PostServer) GetPost(ctx context.Context, req *pbPost.GetPostRequest) (*pbPost.PostResponse, error) {
	post, files, err := models.GetPostByID(int(req.Id))
	if err != nil {
		return generatePostError(err.Error())
	}

	if !post.Finished {
		return &pbPost.PostResponse{
			Success:  true,
			Finished: false,
		}, nil
	}

	postResponse, err := mappers.PostModelToPostResponse(post, files)
	if err != nil {
		return generatePostError(err.Error())
	}
	if !postResponse.Success {
		return generatePostError(postResponse.Error)
	}

	return postResponse, nil
}
