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

func (r *productResolver) Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error) {
	conn, _ := proto.GetConnection()
	defer conn.Close()
	c := protoGenerated.NewReviewServiceClient(conn)
	response, err := c.ProductReviews(ctx, &protoGenerated.ProductReviewsQuery{Upc: obj.Upc}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Printf("Error when calling FindUserByID: %s\n", err)
	}

	return convertReviewsResponse(response), nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	conn, _ := proto.GetConnection()
	defer conn.Close()
	c := protoGenerated.NewReviewServiceClient(conn)

	uintId, err := strconv.ParseUint(obj.ID, 10, 64)
	if err != nil {
		log.Printf("can not convert id to uint: %s\n", err)
	}

	response, err := c.UserReviews(ctx, &protoGenerated.UserReviewsQuery{Id: uintId}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Printf("Error when calling FindUserByID: %s\n", err)
	}

	return convertReviewsResponse(response), nil
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type productResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
