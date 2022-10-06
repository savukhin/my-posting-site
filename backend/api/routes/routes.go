package routes

import (
	"errors"
	api_controllers "my-posting-site/backend/api/controllers"
	grpc_clients "my-posting-site/backend/common/grpc"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
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

func Routes(router *mux.Router, client *grpc_clients.GRPCClients) (*mux.Router, error) {
	authClient := client.GetAuthClient()
	userClient := client.GetUserClient()
	if authClient == nil {
		return nil, errors.New("no auth client")
	}
	if userClient == nil {
		return nil, errors.New("no user client")
	}

	authRoutes(router.PathPrefix("/auth").Subrouter(), *authClient)
	userRoutes(router.PathPrefix("/user").Subrouter(), *userClient)
	// authRoutes(router, *authClient)
	return router, nil
}
