package main

import (
	"log"

	"github.com/oneshick/users-service/internal/database"
	transportgrpc "github.com/oneshick/users-service/internal/transport/grpc"
	"github.com/oneshick/users-service/internal/user"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Создание репозитория и сервиса
	repo := user.NewRepository(database.DB)
	svc := user.NewService(repo)

	// Запуск gRPC сервера
	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
