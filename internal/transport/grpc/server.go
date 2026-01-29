package transportgrpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userpb "github.com/oneshick/project-protos/proto/user"
	"github.com/oneshick/users-service/internal/user"
)

func RunGRPC(svc user.Service) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	reflection.Register(grpcServer)

	log.Println("gRPC server started on port 50051")

	return grpcServer.Serve(lis)
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("gRPC method: %s", info.FullMethod)
	return handler(ctx, req)
}
