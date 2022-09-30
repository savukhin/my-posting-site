package main

import (
	"fmt"
	"log"
	"my-posting-site/backend/api/routes"
	grpc_clients "my-posting-site/backend/common/grpc"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func createRESTSerer(port string, client *grpc_clients.GRPCClients) {
	r := mux.NewRouter()
	// router := r
	router := r.PathPrefix("/api").Subrouter()
	_, err := routes.Routes(router, client)
	if err != nil {
		fmt.Println(err)
		return
	}

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "enctype"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("REST server starting...")

	log.Fatal(http.ListenAndServe(
		":"+port,
		handlers.CORS(credentials, methods, origins, ttl, headers)(r),
	))
}

func main() {
	opts := make([]grpc_clients.GRPCConnectionOpt, 0)
	// opts = append(opts, grpc_clients.GRPCConnectionOpt{Url: "localhost:5300", ServerType: grpc_clients.HelloWorld})
	opts = append(opts, grpc_clients.GRPCConnectionOpt{Url: "127.0.0.1:5300", ServerType: grpc_clients.Auth})

	client, err := grpc_clients.BuildGRPCClients(opts)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	createRESTSerer("3000", client)
}
