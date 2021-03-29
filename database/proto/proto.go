package proto

import (
	"context"
	"learn-apollo-federation-gqlgen/database/proto/generated"
)

type Server struct{}

func (s *Server) ProductReviews(ctx context.Context, query *generated.ProductReviewsQuery) (*generated.ReviewsResponse, error) {
	panic("implement me")
}

func (s *Server) UserReviews(ctx context.Context, query *generated.UserReviewsQuery) (*generated.ReviewsResponse, error) {
	panic("implement me")
}

func (s *Server) FindProductByUpc(ctx context.Context, query *generated.ProductQuery) (*generated.Product, error) {
	panic("implement me")
}

func (s *Server) TopProducts(ctx context.Context, query *generated.TopProductsQuery) (*generated.ProductsResponse, error) {
	panic("implement me")
}

func (s *Server) FindUserByID(ctx context.Context, query *generated.UserQuery) (*generated.User, error) {
	panic("implement me")
}



