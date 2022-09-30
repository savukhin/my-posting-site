package api_utils

import (
	"encoding/json"
	"my-posting-site/backend/api/dto"
	"net/http"
)

func ResponseEmptySucess(res http.ResponseWriter, status int) {
	response := &dto.DefaultResponse{
		Has_err: false,
		Msg:     "Success",
	}
	b, _ := json.Marshal(response)

	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func ResponseError(res http.ResponseWriter, message error, status int) error {
	response := &dto.DefaultResponse{
		Has_err: true,
		Msg:     message.Error(),
	}
	b, err := json.Marshal(response)

	if err != nil {
		return err
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(b)
	return nil
}
