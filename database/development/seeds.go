package main

import (
	"learn-apollo-federation-gqlgen/database/mapping"
)

func main() {
	db, _ := mapping.GetConnection()
	err := db.AutoMigrate(&mapping.User{}, &mapping.Product{}, &mapping.Review{})
	if err != nil {
		panic("Failed to migrate database.")
	}
	db.Delete(&mapping.User{})
	db.Delete(&mapping.Product{})
	db.Delete(&mapping.Review{})
	userMe := mapping.User{UserName: "Me"}
	userOther := mapping.User{UserName: "Other"}
	db.Create(&userMe)
	db.Create(&userOther)
	product1 := mapping.Product{Upc: "top-1", Name: "Trilby", Price: 11}
	product2 := mapping.Product{Upc: "top-2", Name: "Fedora", Price: 22}
	product3 := mapping.Product{Upc: "top-3", Name: "Boater", Price: 33}
	db.Create(&product1)
	db.Create(&product2)
	db.Create(&product3)
	review1 := mapping.Review{
		Body: "A highly effective form of birth control.",
		ProductID: product1.ID,
		AuthorID:  userMe.ID,
	}
	review2 := mapping.Review{
		Body: "Fedoras are one of the most fashionable hats around and can look great with a variety of outfits.",
		ProductID: product2.ID,
		AuthorID:  userMe.ID,
	}
	review3 := mapping.Review{
		Body: "A highly effective form of birth control.",
		ProductID: product2.ID,
		AuthorID:  userOther.ID,
	}
	db.Create(&review1)
	db.Create(&review2)
	db.Create(&review3)
}
