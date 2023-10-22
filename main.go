package main

import (
	"fmt"
	"log"
	"net"
	"os"

	productgrpc "github.com/dhrleandro/product-grpc-golang/application/grpc"
	pb "github.com/dhrleandro/product-grpc-golang/application/grpc/protofiles"
	"github.com/dhrleandro/product-grpc-golang/application/usecase"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/database"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/database/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db := database.ConnectDB(os.Getenv("env"))

	// Fica ouvindo conexões na porta 9000
	l, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// inicia servidor gRPC
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	// registra o service
	r := &repository.ProductRepositoryGORM{Db: db}
	productUseCase := &usecase.ProductUseCase{Pr: r}
	pb.RegisterProductServiceServer(server, productgrpc.NewService(productUseCase))
	reflection.Register(server)
	fmt.Println("Serve")

	// serve as conexões
	server.Serve(l)
}
