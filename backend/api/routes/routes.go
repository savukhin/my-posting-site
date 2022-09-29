package routes

import (
	"context"
	"encoding/json"
	"fmt"
	pb "my-posting-site/common/protobuf/golang/helloWorld"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router, client pb.HelloWorldClient) *mux.Router {
	router.HandleFunc("/hello_world", HelloWorld(client))
	return router
}

type Message struct {
	Msg string `json:"msg"`
}

func HelloWorld(client pb.HelloWorldClient) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		msg := &Message{}
		err := json.NewDecoder(req.Body).Decode(msg)

		if err != nil {
			fmt.Println("Error parsing json: ", err)
		}

		gret, err := client.Greeting(context.Background(), &pb.RequestHello{MessageHello: msg.Msg})

		res.Header().Set("Content-Type", "application/json")

		response := struct {
			Err     string `json:"error"`
			Has_err bool   `json:"has_error"`
			Msg     string `json:"msg"`
		}{
			Err:     "",
			Has_err: false,
			Msg:     "",
		}

		if err != nil {
			response.Err = err.Error()
			response.Has_err = true
		} else {
			response.Msg = gret.GetMessageHello()
		}

		b, _ := json.Marshal(response)
		res.Write(b)
	}
}
