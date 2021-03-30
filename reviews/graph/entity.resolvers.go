package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"learn-apollo-federation-gqlgen/reviews/graph/generated"
	"learn-apollo-federation-gqlgen/reviews/graph/model"
	"learn-apollo-federation-gqlgen/reviews/proto"
	protoGenerated "learn-apollo-federation-gqlgen/reviews/proto/generated"
	"log"
	"strconv"
)

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {
	conn, _ := proto.GetConnection()
	defer conn.Close()
	c := protoGenerated.NewProductServiceClient(conn)
	response, err := c.FindProductByUpc(ctx, &protoGenerated.ProductQuery{Upc: upc}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Printf("Error when calling FindUserByID: %s\n", err)
	}

	return convertProduct(response), nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	conn, _ := proto.GetConnection()
	defer conn.Close()
	c := protoGenerated.NewUserServiceClient(conn)

	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Printf("can not convert id to uint: %s\n", err)
	}

	response, err := c.FindUserByID(ctx, &protoGenerated.UserQuery{Id: uintId}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Printf("Error when calling FindUserByID: %s\n", err)
	}

	return convertUser(response), nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
