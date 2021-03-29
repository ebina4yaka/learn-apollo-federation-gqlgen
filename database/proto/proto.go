package proto

import (
	"context"
	"learn-apollo-federation-gqlgen/database/mapping"
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
	db, _ := mapping.GetConnection()
	tx := db.WithContext(ctx)
	var products []mapping.Product
	tx.Order("upc").Limit(int(query.First)).Find(&products)
	var results []*generated.Product
	for _, product := range products {
		results = append(results, &generated.Product{
			Id:    uint64(product.ID),
			Upc:   product.Upc,
			Price: int64(product.Price),
		})
	}

	return &generated.ProductsResponse{
		Products: results,
	}, nil
}

func (s *Server) FindUserByID(ctx context.Context, query *generated.UserQuery) (*generated.User, error) {
	db, _ := mapping.GetConnection()
	tx := db.WithContext(ctx)
	var user mapping.User
	tx.First(&user, query.Id)
	result := generated.User{
		Id:       uint64(user.ID),
		Username: user.UserName,
	}
	return &result, nil
}
