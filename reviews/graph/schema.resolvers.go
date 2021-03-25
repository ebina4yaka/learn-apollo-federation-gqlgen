package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"learn-apollo-federation-gqlgen/reviews/graph/model"
)

func (r *productResolver) Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error) {
	var res []*model.Review

	for _, review := range reviews {
		if review.Product.Upc == obj.Upc {
			res = append(res, review)
		}
	}

	return res, nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	var res []*model.Review

	for _, review := range reviews {
		if review.Author.ID == obj.ID {
			res = append(res, review)
		}
	}

	return res, nil
}

type ProductResolver interface {
	Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error)
}

type UserResolver interface {
	Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error)
}

func (r *Resolver) Product() ProductResolver { return &productResolver{r} }

func (r *Resolver) User() UserResolver { return &userResolver{r} }

type productResolver struct{ *Resolver }
type userResolver struct{ *Resolver }