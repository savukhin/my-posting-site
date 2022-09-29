package main

import (
	"log"
	"my-posting-site/backend/api/routes"
	pb "my-posting-site/common/protobuf/golang/helloWorld"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createGRPSClient(port string) (pb.HelloWorldClient, *grpc.ClientConn, error) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:5300", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("connection:", err)
		return nil, nil, err
	}

	var client pb.HelloWorldClient = pb.NewHelloWorldClient(conn)

	return client, conn, nil
}

func createRESTSerer(port string, client pb.HelloWorldClient) {
	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()
	routes.Routes(router, client)

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "enctype"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(
		":"+port,
		handlers.CORS(credentials, methods, origins, ttl, headers)(r),
	))
}

func main() {
	client, conn, err := createGRPSClient("5300")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	createRESTSerer("3000", client)
}
