package routes

import (
	"errors"
	api_controllers "my-posting-site/backend/api/controllers"
	grpc_clients "my-posting-site/backend/common/grpc"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
	pbPost "my-posting-site/common/protobuf/golang/post"
	pbUser "my-posting-site/common/protobuf/golang/user"
	"net/http"

	"github.com/gorilla/mux"
)

func authRoutes(router *mux.Router, client pbAuth.AuthenticationClient) *mux.Router {
	router.HandleFunc("/login", api_controllers.Login(client)).Methods(http.MethodPost)
	router.HandleFunc("/register", api_controllers.Register(client)).Methods(http.MethodPost)

	return router
}

func userRoutes(router *mux.Router, client pbUser.UserClient) *mux.Router {
	router.HandleFunc("/get_profile/{user_id:[0-9]+}", api_controllers.GetProfile(client)).Methods(http.MethodGet)

	return router
}

func postRoutes(router *mux.Router, client pbPost.PostingClient) *mux.Router {
	router.HandleFunc("/create_post", api_controllers.CreatePost(client)).Methods(http.MethodPost)
	router.HandleFunc("/get_post/{post_id:[0-9]+}", api_controllers.GetPost(client)).Methods(http.MethodGet)
	router.HandleFunc("/get_file/{filepath:.*}", api_controllers.GetFile(client)).Methods(http.MethodGet)

	return router
}

func Routes(router *mux.Router, client *grpc_clients.GRPCClients) (*mux.Router, error) {
	authClient := client.GetAuthClient()
	userClient := client.GetUserClient()
	postingClient := client.GetPostingClient()
	if authClient == nil {
		return nil, errors.New("no auth client")
	}
	if userClient == nil {
		return nil, errors.New("no user client")
	}
	if postingClient == nil {
		return nil, errors.New("no posting client")
	}

	authRoutes(router.PathPrefix("/auth").Subrouter(), *authClient)
	userRoutes(router.PathPrefix("/user").Subrouter(), *userClient)
	postRoutes(router.PathPrefix("/post").Subrouter(), *postingClient)
	// authRoutes(router, *authClient)
	return router, nil
}
