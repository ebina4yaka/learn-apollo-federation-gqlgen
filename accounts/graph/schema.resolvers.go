package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"learn-apollo-federation-gqlgen/accounts/graph/generated"
	"learn-apollo-federation-gqlgen/accounts/graph/model"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID:       "1234",
		Username: "Me",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }