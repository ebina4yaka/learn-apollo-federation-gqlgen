package proto

import (
	"context"
	"gorm.io/gorm"
	"learn-apollo-federation-gqlgen/database/mapping"
	"learn-apollo-federation-gqlgen/database/proto/generated"
)

type Server struct{}

func getAuthorsFromReviews(tx *gorm.DB, reviews []mapping.Review) []mapping.User {
	var authors []mapping.User
	var authorIds []uint
	for _, review := range reviews {
		authorIds = append(authorIds, review.AuthorID)
	}
	tx.Find(&authors, authorIds)
	return authors
}

func getProductsFromReviews(tx *gorm.DB, reviews []mapping.Review) []mapping.Product {
	var products []mapping.Product
	var productIds []uint
	for _, review := range reviews {
		productIds = append(productIds, review.ProductID)
	}
	tx.Find(&products, productIds)
	return products
}

func searchUserFromSlice(id uint, users []mapping.User) *mapping.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func searchProductFromSlice(id uint, products []mapping.Product) *mapping.Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func convertUser(user *mapping.User) *generated.User {
	if user == nil {
		return nil
	}
	return &generated.User{
		Id:       uint64(user.ID),
		Username: user.UserName,
	}
}

func convertProduct(product *mapping.Product) *generated.Product {
	if product == nil {
		return nil
	}
	return &generated.Product{
		Id:    uint64(product.ID),
		Upc:   product.Upc,
		Name:   product.Name,
		Price: int64(product.Price),
	}
}

func convertReview(review *mapping.Review, author *generated.User, product *generated.Product) *generated.Review {
	if review == nil {
		return nil
	}
	return &generated.Review{
		Id:      uint64(review.ID),
		Body:    review.Body,
		Author:  author,
		Product: product,
	}
}

func (s *Server) ProductReviews(ctx context.Context, query *generated.ProductReviewsQuery) (*generated.ReviewsResponse, error) {
	db, _ := mapping.GetConnection()
	tx := db.WithContext(ctx)
	var reviews []mapping.Review
	tx.Table("reviews").
		Select("reviews.id, reviews.body, reviews.author_id, reviews.product_id").
		Joins("left join products on products.id = reviews.product_id").
		Where("products.upc = ?", query.Upc).
		Find(&reviews)
	authors := getAuthorsFromReviews(tx, reviews)
	products := getProductsFromReviews(tx, reviews)
	var results []*generated.Review
	for _, review := range reviews {
		author := searchUserFromSlice(review.AuthorID, authors)
		product := searchProductFromSlice(review.ProductID, products)
		results = append(results, convertReview(&review, convertUser(author), convertProduct(product)))
	}

	return &generated.ReviewsResponse{
		Reviews: results,
	}, nil
}

func (s *Server) UserReviews(ctx context.Context, query *generated.UserReviewsQuery) (*generated.ReviewsResponse, error) {
	db, _ := mapping.GetConnection()
	tx := db.WithContext(ctx)
	var reviews []mapping.Review
	tx.Where("author_id = ?", query.Id).Find(&reviews)
	authors := getAuthorsFromReviews(tx, reviews)
	products := getProductsFromReviews(tx, reviews)
	var results []*generated.Review
	for _, review := range reviews {
		author := searchUserFromSlice(review.AuthorID, authors)
		product := searchProductFromSlice(review.ProductID, products)
		results = append(results, convertReview(&review, convertUser(author), convertProduct(product)))
	}

	return &generated.ReviewsResponse{
		Reviews: results,
	}, nil
}

func (s *Server) FindProductByUpc(ctx context.Context, query *generated.ProductQuery) (*generated.Product, error) {
	db, _ := mapping.GetConnection()
	tx := db.WithContext(ctx)
	var product mapping.Product
	tx.Where("upc = ?", query.Upc).First(&product)
	return convertProduct(&product), nil
}

func (s *Server) TopProducts(ctx context.Context, query *generated.TopProductsQuery) (*generated.ProductsResponse, error) {
	db, _ := mapping.GetConnection()
	tx := db.WithContext(ctx)
	var products []mapping.Product
	tx.Order("upc").Limit(int(query.First)).Find(&products)
	var results []*generated.Product
	for _, product := range products {
		results = append(results, convertProduct(&product))
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
	return convertUser(&user), nil
}
