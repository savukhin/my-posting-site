package api_controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"my-posting-site/backend/api/dto"
	api_utils "my-posting-site/backend/api/utils"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
	"net/http"
)

func Login(client pbAuth.AuthenticationClient) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		data := &dto.Login{}
		err := json.NewDecoder(req.Body).Decode(data)

		if err != nil {
			fmt.Println("Error parsing json: ", err)
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}

		grpcResponse, err := client.Login(context.Background(), data.ToProtobuf())

		res.Header().Set("Content-Type", "application/json")

		response := dto.DefaultJWTResponse{
			DefaultResponse: dto.DefaultResponse{},
			Token:           "",
		}

		if err != nil {
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		} else {
			response.Token = grpcResponse.Token
			response.Msg = "Succesfully login"
		}

		b, _ := json.Marshal(response)
		res.Write(b)
	}
}
