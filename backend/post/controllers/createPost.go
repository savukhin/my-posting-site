package auth_controllers

import (
	"context"
	"log"
	grpc_clients "my-posting-site/backend/common/grpc"
	"my-posting-site/backend/common/mappers"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
	pbPost "my-posting-site/common/protobuf/golang/post"
)

type PostServer struct {
	pbPost.UnimplementedPostingServer
}

var Client *grpc_clients.GRPCClients

func init() {
	opts := make([]grpc_clients.GRPCConnectionOpt, 0)
	opts = append(opts, grpc_clients.GRPCConnectionOpt{Url: "127.0.0.1:3400", ServerType: grpc_clients.Auth})

	var err error
	Client, err = grpc_clients.BuildGRPCClients(opts)
	if err != nil {
		log.Fatal(err)
		return
	}
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

func (server *PostServer) CreatePost(ctx context.Context, req *pbPost.CreatePostRequest) (*pbPost.DefaultResponse, error) {
	client := (*Client.GetAuthClient())
	grpcResponse, err := client.CheckJWT(context.Background(), &pbAuth.JWTRequest{Token: req.Token})

	if err != nil {
		return generateError(err.Error())
	}
	if !grpcResponse.Success {
		return generateError(grpcResponse.Error)
	}

	post, files, err := mappers.CreatePostRequestToPost(req, int(grpcResponse.UserId), "post/uploads")
	if err != nil {
		return generateError(err.Error())
	}

	post.Save(files)

	return &pbPost.DefaultResponse{
		Success: true,
		PostId:  int32(post.ID),
	}, nil
}
