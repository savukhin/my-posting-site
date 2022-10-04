package routes

import (
	"errors"
	api_controllers "my-posting-site/backend/api/controllers"
	grpc_clients "my-posting-site/backend/common/grpc"
	pbAuth "my-posting-site/common/protobuf/golang/auth"

	"github.com/gorilla/mux"
)

func authRoutes(router *mux.Router, client pbAuth.AuthenticationClient) *mux.Router {
	router.HandleFunc("/login", api_controllers.Login(client))
	router.HandleFunc("/register", api_controllers.Register(client))

	return router
}

func Routes(router *mux.Router, client *grpc_clients.GRPCClients) (*mux.Router, error) {
	authClient := client.GetAuthClient()
	if authClient == nil {
		return nil, errors.New("no auth client")
	}

	authRoutes(router.PathPrefix("/auth").Subrouter(), *authClient)
	// authRoutes(router, *authClient)
	return router, nil
}
