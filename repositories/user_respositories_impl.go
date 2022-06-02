package repositories

import (
	"errors"
	"final/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepostoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepostoryImpl{}
}

func (userRepository *UserRepostoryImpl) CreateUser(db *gorm.DB, user models.User) (models.User, error) {
	fmt.Println("user", user)
	err := db.Create(&user).Error
	if err != nil {
		return user, errors.New(err.Error())
	}

	return user, nil
}
