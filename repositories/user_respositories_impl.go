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
	err := db.Create(&user).Error
	if err != nil {
		return user, errors.New(err.Error())
	}

	return user, nil
}

func (userRepository *UserRepostoryImpl) LoginUser(db *gorm.DB, user models.User) (models.User, error) {
	result := db.Where("email = ?", user.Email).First(&user)

	if result.RowsAffected == 0 {
		err := fmt.Sprintf("user with email %s not found", user.Email)
		return user, errors.New(err)
	}

	return user, nil
}
