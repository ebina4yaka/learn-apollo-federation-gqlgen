package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learn-apollo-federation-gqlgen/database/mapping"
	"learn-apollo-federation-gqlgen/database/proto"
	"learn-apollo-federation-gqlgen/database/proto/generated"
	"log"
	"net"
	"os"
)

const defaultPort = "9000"

func main() {
	db, _ := mapping.GetConnection()
	err := db.AutoMigrate(&mapping.User{}, &mapping.Product{}, &mapping.Review{})
	if err != nil {
		panic("Failed to migrate database.")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := proto.Server{}
	grpcServer := grpc.NewServer()

	generated.RegisterUserServiceServer(grpcServer, &s)
	generated.RegisterProductServiceServer(grpcServer, &s)
	generated.RegisterReviewServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
