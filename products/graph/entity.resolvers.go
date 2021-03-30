package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"learn-apollo-federation-gqlgen/products/graph/generated"
	"learn-apollo-federation-gqlgen/products/graph/model"
	"learn-apollo-federation-gqlgen/products/proto"
	protoGenerated "learn-apollo-federation-gqlgen/products/proto/generated"
	"log"
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

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
