package grpc_clients

import (
	"errors"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
	pbHelloWorld "my-posting-site/common/protobuf/golang/helloWorld"
	pbPost "my-posting-site/common/protobuf/golang/post"
	pbUser "my-posting-site/common/protobuf/golang/user"
	"net"

	"google.golang.org/grpc"
)

type ServerType string

const (
	HelloWorld ServerType = "helloWorld"
	Auth       ServerType = "auth"
	User       ServerType = "user"
	Post       ServerType = "post"
)

type GRPCServer struct {
	address          string
	listener         net.Listener
	grpcServer       *grpc.Server
	serverType       ServerType
	helloWorldServer *pbHelloWorld.HelloWorldServer
	authServer       *pbAuth.AuthenticationServer
	userServer       *pbUser.UserServer
	postingServer    *pbPost.PostingServer
}

// type HelloWorldServer struct {
// 	pbHelloWorld.UnimplementedHelloWorldServer
// }

// type AuthServer struct {
// 	pbAuth.UnimplementedAuthenticationServer
// }

func buildGRPCServer(address string, serverType ServerType) (*GRPCServer, error) {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return nil, errors.New("fatal: failed to listen: " + err.Error())
	}

	opts := []grpc.ServerOption{}

	server := &GRPCServer{
		address:    address,
		listener:   listener,
		grpcServer: grpc.NewServer(opts...),
		serverType: serverType,
	}

	return server, nil
}

func (server *GRPCServer) addAuthServer(serv *pbAuth.AuthenticationServer) {
	server.authServer = serv
}

func (server *GRPCServer) addHelloWorldServer(serv *pbHelloWorld.HelloWorldServer) {
	server.helloWorldServer = serv
}

func (server *GRPCServer) addUserServer(serv *pbUser.UserServer) {
	server.userServer = serv
}

func (server *GRPCServer) addPostingServer(serv *pbPost.PostingServer) {
	server.postingServer = serv
}

func (server *GRPCServer) registerServer() error {
	switch server.serverType {
	case HelloWorld:
		if server.helloWorldServer == nil {
			return errors.New("no hello world server provided")
		}
		pbHelloWorld.RegisterHelloWorldServer(server.grpcServer, *server.helloWorldServer)

	case Auth:
		if server.authServer == nil {
			return errors.New("no auth server provided")
		}
		pbAuth.RegisterAuthenticationServer(server.grpcServer, *server.authServer)
	case User:
		if server.userServer == nil {
			return errors.New("no user server provided")
		}
		pbUser.RegisterUserServer(server.grpcServer, *server.userServer)
	case Post:
		if server.postingServer == nil {
			return errors.New("no user server provided")
		}
		pbPost.RegisterPostingServer(server.grpcServer, *server.postingServer)
	}

	return nil
}

func (server *GRPCServer) Serve() error {
	return server.grpcServer.Serve(server.listener)
}

func (server *GRPCServer) Close() error {
	return server.listener.Close()
}
