package graph

import (
	"learn-apollo-federation-gqlgen/products/graph/model"
	"learn-apollo-federation-gqlgen/products/proto/generated"
)

func convertProduct(product *generated.Product) *model.Product {
	if product == nil {
		return nil
	}

	return &model.Product{
		Upc: product.Upc,
		Name: product.Name,
		Price: int(product.Price),
	}
}

func convertProductsResponse(productsResponse *generated.ProductsResponse) []*model.Product {
	if productsResponse == nil {
		return nil
	}

	var products []*model.Product

	for _, product := range productsResponse.Products {
		products = append(products, &model.Product{
			Upc: product.Upc,
			Name: product.Name,
			Price: int(product.Price),
		})
	}

	return products
}
