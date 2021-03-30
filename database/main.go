package main

import (
	"fmt"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"learn-apollo-federation-gqlgen/database/proto"
	"learn-apollo-federation-gqlgen/database/proto/generated"
	"log"
	"net"
	"os"
)

const defaultPort = "9000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listen to :%s\n", port)

	s := proto.Server{}
	grpcServer := grpc.NewServer()

	generated.RegisterUserServiceServer(grpcServer, &s)
	generated.RegisterProductServiceServer(grpcServer, &s)
	generated.RegisterReviewServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
