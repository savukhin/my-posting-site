package api_controllers

import (
	"context"
	"encoding/json"
	"errors"
	"my-posting-site/backend/api/dto"
	api_utils "my-posting-site/backend/api/utils"
	pbUser "my-posting-site/backend/common/protobuf/golang/user"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProfile(client pbUser.UserClient) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		user_id, err := strconv.Atoi(vars["user_id"])
		if err != nil {
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}

		grpcResponse, err := client.GetProfile(context.Background(), &pbUser.ProfileRequest{Id: int32(user_id)})

		res.Header().Set("Content-Type", "application/json")

		if err != nil {
			api_utils.ResponseError(res, err, http.StatusBadRequest)
			return
		}
		if !grpcResponse.Success {
			api_utils.ResponseError(res, errors.New(grpcResponse.Error), http.StatusBadRequest)
			return
		}

		response := dto.ProfileFromProtobuf(grpcResponse)

		b, _ := json.Marshal(response)
		res.Write(b)
	}
}
