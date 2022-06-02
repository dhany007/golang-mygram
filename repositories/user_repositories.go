package repositories

import (
	"final/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(db *gorm.DB, user models.User) (models.User, error)
	LoginUser(db *gorm.DB, user models.User) (models.User, error)
}
