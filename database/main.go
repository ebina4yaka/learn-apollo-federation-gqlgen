package main

import "learn-apollo-federation-gqlgen/database/mapping"

func main() {
	db, _ := mapping.GetConnection()
	err := db.AutoMigrate(&mapping.User{}, &mapping.Product{}, &mapping.Review{})
	if err != nil {
		panic("Failed to migrate database.")
	}
}
