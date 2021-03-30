package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"learn-apollo-federation-gqlgen/accounts/graph/generated"
	"learn-apollo-federation-gqlgen/accounts/graph/model"
	"learn-apollo-federation-gqlgen/accounts/proto"
	protoGenerated "learn-apollo-federation-gqlgen/accounts/proto/generated"
	"log"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	conn, _ := proto.GetConnection()
	defer conn.Close()
	c := protoGenerated.NewUserServiceClient(conn)

	response, err := c.FindUserByID(ctx, &protoGenerated.UserQuery{Id: 1}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Printf("Error when calling FindUserByID in Me: %s\n", err)
	}

	return convertUser(response), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
