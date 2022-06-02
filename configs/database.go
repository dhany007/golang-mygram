package configs

import (
	"final/helpers"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = ""
	DB_NAME = "postgres"
)

func StartDB() *gorm.DB {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	helpers.PanicIfError(err)

	log.Default().Println("connection db succcess")

	return db
}
