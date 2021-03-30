package proto

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

const defaultPort = "9000"
const defaultHost = "localhost"

func GetConnection() (*grpc.ClientConn, error) {
	port := os.Getenv("GRPC_PORT")
	host := os.Getenv("GRPC_HOST")
	if port == "" {
		port = defaultPort
	}
	if host == "" {
		host = defaultHost
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn, nil
}
