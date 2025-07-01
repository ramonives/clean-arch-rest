package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"clean-arch-rest/internal/delivery/graphql"
	grpcHandler "clean-arch-rest/internal/delivery/grpc"
	httpHandler "clean-arch-rest/internal/delivery/http"
	"clean-arch-rest/internal/infrastructure/database"
	"clean-arch-rest/internal/infrastructure/repository"
	"clean-arch-rest/internal/usecase"
	pb "clean-arch-rest/proto/order"
)

func main() {
	// Conecta com o banco de dados PostgreSQL
	db := database.NewPostgresConnection()

	// Inicializa o repositório
	orderRepo := repository.NewOrderRepository(db)

	// Inicializa o caso de uso
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// Inicia o servidor HTTP REST
	go func() {
		router := httpHandler.SetupRouter(orderUseCase)
		fmt.Println("Servidor HTTP rodando na porta 8080")
		log.Fatal(router.Run(":8080"))
	}()

	// Inicia o servidor GraphQL
	go func() {
		graphqlHandler := graphql.NewGraphQLHandler(orderUseCase)
		http.Handle("/graphql", graphqlHandler)
		fmt.Println("Servidor GraphQL rodando na porta 8081")
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	// Inicia o servidor gRPC
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderServer := grpcHandler.NewOrderServer(orderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderServer)
	reflection.Register(grpcServer)

	fmt.Println("Servidor gRPC rodando na porta 9090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
