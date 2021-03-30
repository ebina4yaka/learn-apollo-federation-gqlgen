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

func (r *queryResolver) TopProducts(ctx context.Context, first *int) ([]*model.Product, error) {
	conn, _ := proto.GetConnection()
	defer conn.Close()
	c := protoGenerated.NewProductServiceClient(conn)
	response, err := c.TopProducts(ctx, &protoGenerated.TopProductsQuery{First: int64(*first)}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Printf("Error when calling FindUserByID: %s\n", err)
	}

	return convertProductsResponse(response), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
