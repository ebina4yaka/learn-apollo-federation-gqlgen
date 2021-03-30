package mapping

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

const defaultDbPort = "5432"
const defaultDbHost = "localhost"
const defaultDbUser = "apollo_federation"
const defaultDbPassword = "password"
const defaultDbName = "apollo_federation_development"

var count = 0

func GetConnection() (*gorm.DB, error) {
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbPort == "" {
		dbPort = defaultDbPort
	}
	if dbHost == "" {
		dbHost = defaultDbHost
	}
	if dbUser == "" {
		dbUser = defaultDbUser
	}
	if dbPassword == "" {
		dbPassword = defaultDbPassword
	}
	if dbName == "" {
		dbName = defaultDbName
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Not ready. Retry connecting...")
		time.Sleep(time.Second)
		count++
		log.Println(count)
		if count > 30 {
			panic("Failed to connect database.")
		}
		return GetConnection()
	}

	return db, nil
}
