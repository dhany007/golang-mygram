package tests

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
	DB_HOST_TEST = "localhost"
	DB_PORT_TEST = "5432"
	DB_USER_TEST = "koinworks"
	DB_PASS_TEST = ""
	DB_NAME_TEST = "koinworks"
)

func StartDBTest() *gorm.DB {
	dataSourceNameTest := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable",
		DB_HOST_TEST, DB_USER_TEST, DB_PASS_TEST, DB_NAME_TEST, DB_PORT_TEST,
	)

	db, err := gorm.Open(postgres.Open(dataSourceNameTest), &gorm.Config{})
	helpers.PanicIfError(err)

	err = db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	helpers.PanicIfError(err)

	log.Default().Println("connection db succcess")

	return db
}
