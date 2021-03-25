package graph

import (
	"context"
	"learn-apollo-federation-gqlgen/reviews/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

type ProductResolver interface {
	Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error)
}

type UserResolver interface {
	Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error)
}
