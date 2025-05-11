package client

import (
	"gateway/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	UserService UserService
	// Добавьте другие сервисы по аналогии
}

var Clients *GRPCClients

func InitGRPCClients(cfg *config.Config) {
	// Подключение к UserService
	userConn, err := grpc.Dial(
		cfg.UserServiceGRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to UserService: %v", err)
	}

	Clients = &GRPCClients{
		UserService: NewUserServiceClient(userConn),
	}
}