package main

import (
	"log"

	"github.com/oneshick/users-service/internal/database"
	transportgrpc "github.com/oneshick/users-service/internal/transport/grpc"
	"github.com/oneshick/users-service/internal/user"
)

func main() {
	database.InitDB()

	repo := user.NewRepository(database.DB)
	svc := user.NewService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
