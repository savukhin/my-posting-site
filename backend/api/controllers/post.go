package api_controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"my-posting-site/backend/api/dto"
	api_utils "my-posting-site/backend/api/utils"
	pbPost "my-posting-site/common/protobuf/golang/post"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(client pbPost.PostingClient) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		jwt := req.Header.Get("Authorization")
		if jwt == "" {
			api_utils.ResponseError(res, errors.New("no token provided"), http.StatusBadRequest)
			return
		}

		data := &dto.Post{}

		photos := make(map[string][]byte)

		count, err := strconv.Atoi(req.FormValue("count"))
		if err != nil || count <= 0 {
			api_utils.ResponseError(res, errors.New("count is not valid number"), http.StatusBadRequest)
			return
		}

		data.Elements = make([]dto.PostElement, count)
		for i := 0; i < count; i++ {
			prefix := strconv.Itoa(i) + "_"
			isText, err := strconv.ParseBool(req.FormValue(prefix + "isText"))
			if err != nil {
				api_utils.ResponseError(res, errors.New("error in isText with prefix "+prefix), http.StatusBadRequest)
				return
			}

			data.Elements[i] = dto.PostElement{
				IsText: isText,
			}

			if isText {
				data.Elements[i].Text = req.FormValue(prefix + "text")
				continue
			}

			data.Elements[i].Title = req.FormValue(prefix + "title")
			data.Elements[i].PhotoFileName = prefix + "photo"

			key := data.Elements[i].PhotoFileName
			file, _, err := req.FormFile(key)
			if err != nil {
				continue
			}
			b, err := io.ReadAll(file)
			if err != nil {
				continue
			}

			photos[key] = b
		}

		protobuf := data.ToProtobuf(photos)

		grpcRequest := &pbPost.CreatePostRequest{Token: jwt, Post: protobuf}
		grpcResponse, err := client.CreatePost(context.Background(), grpcRequest)

		res.Header().Set("Content-Type", "application/json")

		if err != nil {
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}
		if !grpcResponse.Success {
			api_utils.ResponseError(res, errors.New(grpcResponse.Error), http.StatusBadRequest)
			return
		}

		response := dto.FromPostDefaultResponse(grpcResponse)

		b, _ := json.Marshal(response)
		res.Write(b)
	}
}

func GetPost(client pbPost.PostingClient) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("grpcResponse, err")
		vars := mux.Vars(req)
		post_id, err := strconv.Atoi(vars["post_id"])
		if err != nil {
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}

		grpcResponse, err := client.GetPost(context.Background(), &pbPost.GetPostRequest{Id: int32(post_id)})

		res.Header().Set("Content-Type", "application/json")

		if err != nil {
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}
		if !grpcResponse.Success {
			api_utils.ResponseError(res, errors.New(grpcResponse.Error), http.StatusBadRequest)
			return
		}

		b, _ := json.Marshal(grpcResponse)
		res.Write(b)
	}
}

func GetFile(client pbPost.PostingClient) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		filepath := vars["filepath"]
		if filepath == "" {
			api_utils.ResponseError(res, errors.New("no filepath provided"), http.StatusBadRequest)
			return
		}

		grpcResponse, err := client.GetFile(context.Background(), &pbPost.GetFileRequest{Path: filepath})

		res.Header().Set("Content-Type", "image/jpeg")

		if err != nil {
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}
		if !grpcResponse.Success {
			api_utils.ResponseError(res, errors.New(grpcResponse.Error), http.StatusBadRequest)
			return
		}

		b, _ := json.Marshal(grpcResponse.File)
		res.Write(b)
	}
}
