package grpc_clients

import (
	"errors"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
	pbHelloWorld "my-posting-site/common/protobuf/golang/helloWorld"
	pbUser "my-posting-site/common/protobuf/golang/user"
)

type GRPCServerBuilder struct {
	address          string
	serverType       ServerType
	helloWorldServer *pbHelloWorld.HelloWorldServer
	authServer       *pbAuth.AuthenticationServer
	userServer       *pbUser.UserServer
}

func NewServer() *GRPCServerBuilder {
	return &GRPCServerBuilder{}
}

func (builder *GRPCServerBuilder) AddPort(address string) *GRPCServerBuilder {
	builder.address = address
	return builder
}

func (builder *GRPCServerBuilder) AddServerType(serverType ServerType) *GRPCServerBuilder {
	builder.serverType = serverType
	return builder
}

func (builder *GRPCServerBuilder) AddHelloWorldServer(server pbHelloWorld.HelloWorldServer) *GRPCServerBuilder {
	builder.serverType = HelloWorld
	builder.helloWorldServer = &server
	return builder
}

func (builder *GRPCServerBuilder) AddAuthServer(server pbAuth.AuthenticationServer) *GRPCServerBuilder {
	builder.serverType = Auth
	builder.authServer = &server
	return builder
}

func (builder *GRPCServerBuilder) AddUserServer(server pbUser.UserServer) *GRPCServerBuilder {
	builder.serverType = User
	builder.userServer = &server
	return builder
}

func (builder *GRPCServerBuilder) Valid() error {
	if builder.address == "" {
		return errors.New("no address provided")
	}
	switch builder.serverType {
	case HelloWorld:
		if builder.helloWorldServer == nil {
			return errors.New("no helloWorld server provided")
		}
	case Auth:
		if builder.authServer == nil {
			return errors.New("no auth server provided")
		}
	case User:
		if builder.userServer == nil {
			return errors.New("no user server provided")
		}
	}

	return nil
}

func (builder *GRPCServerBuilder) Build() (*GRPCServer, error) {
	if err := builder.Valid(); err != nil {
		return nil, err
	}

	server, err := buildGRPCServer(
		builder.address,
		builder.serverType,
	)

	if err != nil {
		return nil, err
	}

	server.addHelloWorldServer(builder.helloWorldServer)
	server.addAuthServer(builder.authServer)
	server.addUserServer(builder.userServer)

	err = server.registerServer()

	return server, err
}
