package graph

import (
	"learn-apollo-federation-gqlgen/reviews/graph/model"
	"learn-apollo-federation-gqlgen/reviews/proto/generated"
	"strconv"
)

func convertProduct(product *generated.Product) *model.Product {
	if product == nil {
		return nil
	}

	return &model.Product{
		Upc: product.Upc,
	}
}

func convertUser(user *generated.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:       strconv.FormatUint(user.Id, 10),
	}
}

func convertReviewsResponse(reviewsResponse *generated.ReviewsResponse) []*model.Review {
	if reviewsResponse == nil {
		return nil
	}

	var reviews []*model.Review

	for _, review := range reviewsResponse.Reviews {
		reviews = append(reviews, &model.Review{
			Body: review.Body,
			Author: convertUser(review.Author),
			Product: convertProduct(review.Product),
		})
	}

	return reviews
}

