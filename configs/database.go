package configs

import (
	"final/helpers"
	"final/models"
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
		"host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable",
		DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	helpers.PanicIfError(err)

	err = autoMigrate(db)
	helpers.PanicIfError(err)

	log.Default().Println("connection db succcess")

	return db
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Photo{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Comment{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.SocialMedia{}); err != nil {
		return err
	}

	return nil
}
