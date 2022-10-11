package post_controllers

import (
	"context"
	kafka_messages "my-posting-site/backend/common/kafka"
	"my-posting-site/backend/common/mappers"
	pbAuth "my-posting-site/backend/common/protobuf/golang/auth"
	pbPost "my-posting-site/backend/common/protobuf/golang/post"
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

func (server *PostServer) CreatePost(ctx context.Context, req *pbPost.CreatePostRequest) (*pbPost.DefaultResponse, error) {
	client := (*Client.GetAuthClient())
	grpcResponse, err := client.CheckJWT(context.Background(), &pbAuth.JWTRequest{Token: req.Token})

	if err != nil {
		return generateError(err.Error())
	}
	if !grpcResponse.Success {
		return generateError(grpcResponse.Error)
	}

	post_model, files, err := mappers.CreatePostRequestToPost(req, int(grpcResponse.UserId), "post/uploads")
	if err != nil {
		return generateError(err.Error())
	}
	post_model.Finished = false
	err = post_model.Save(files)
	if err != nil {
		return generateError(err.Error())
	}

	post_kafka := mappers.GrpcPostToKafkaPost(req, int(grpcResponse.UserId), post_model.ID)
	Produce(post_kafka)

	return &pbPost.DefaultResponse{
		Success: true,
		PostId:  int32(post_model.ID),
	}, nil
}

func CreatePost(post *kafka_messages.Post) error {
	post_model, files, err := mappers.KafkaPostToPost(post, "post/uploads")
	if err != nil {
		return err
	}
	post_model.Finished = true

	return post_model.Save(files)
}
