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
	"strconv"
)

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
