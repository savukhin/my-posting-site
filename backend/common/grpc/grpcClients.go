package grpc_clients

import (
	"context"
	pbAuth "my-posting-site/common/protobuf/golang/auth"
	pbHelloWorld "my-posting-site/common/protobuf/golang/helloWorld"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCConnection struct {
	conn                 *grpc.ClientConn
	serverType           ServerType
	ctx                  context.Context
	helloWorldClient     pbHelloWorld.HelloWorldClient
	authenticationClient pbAuth.AuthenticationClient
}

type GRPCConnectionOpt struct {
	Url        string
	ServerType ServerType
}

func createGRPCConnection(opt GRPCConnectionOpt) (*GRPCConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		opt.Url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	grpcConnection := &GRPCConnection{
		conn:       conn,
		serverType: opt.ServerType,
		ctx:        ctx,
	}

	switch opt.ServerType {
	case HelloWorld:
		grpcConnection.helloWorldClient = pbHelloWorld.NewHelloWorldClient(conn)
	case Auth:
		grpcConnection.authenticationClient = pbAuth.NewAuthenticationClient(conn)
	}

	return grpcConnection, err
}

type GRPCClients struct {
	clients              []*GRPCConnection
	helloWorldClient     *pbHelloWorld.HelloWorldClient
	authenticationClient *pbAuth.AuthenticationClient
}

func (clients *GRPCClients) addClient(url string, serverType ServerType) error {
	conn, err := createGRPCConnection(GRPCConnectionOpt{url, serverType})
	if err == nil {
		clients.clients = append(clients.clients, conn)
		switch conn.serverType {
		case HelloWorld:
			clients.helloWorldClient = &conn.helloWorldClient
		case Auth:
			clients.authenticationClient = &conn.authenticationClient
		}
	}

	return err
}

func (clients *GRPCClients) AddHelloWorldConnection(conn *GRPCConnection) *GRPCClients {
	clients.clients = append(clients.clients, conn)
	return clients
}

func BuildGRPCClients(opts []GRPCConnectionOpt) (*GRPCClients, error) {
	response := &GRPCClients{
		clients: make([]*GRPCConnection, 0),
	}
	connections := make([]*GRPCConnection, 0)
	for _, opt := range opts {
		// conn, err := createGRPCConnection(opt)
		err := response.addClient(opt.Url, opt.ServerType)
		if err != nil {
			return nil, err
		}
		// connections = append(connections, conn)

		// fmt.Println(opt)
	}

	response.clients = append(response.clients, connections...)

	return response, nil
}

func (clients *GRPCClients) FindType(serverType ServerType) *GRPCConnection {
	for _, conn := range clients.clients {
		if conn.serverType == serverType {
			return conn
		}
	}
	return nil
}

func (clients *GRPCClients) GetHelloWorldClient() *pbHelloWorld.HelloWorldClient {
	return clients.helloWorldClient
}

func (clients *GRPCClients) GetAuthClient() *pbAuth.AuthenticationClient {
	return clients.authenticationClient
}

func (clients *GRPCClients) Close() {
	for _, client := range clients.clients {
		client.conn.Close()
	}
}
