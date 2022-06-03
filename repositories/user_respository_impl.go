package repositories

import (
	"errors"
	"final/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) CreateUser(db *gorm.DB, user models.User) (models.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return user, errors.New(err.Error())
	}

	return user, nil
}

func (userRepository *UserRepositoryImpl) LoginUser(db *gorm.DB, user models.User) (models.User, error) {
	result := db.Where("email = ?", user.Email).First(&user)

	if result.RowsAffected == 0 {
		err := fmt.Sprintf("user with email %s not found", user.Email)
		return user, errors.New(err)
	}

	return user, nil
}

func (userRepository *UserRepositoryImpl) UpdateUser(db *gorm.DB, user models.User, userId int) (models.User, error) {
	tempRequest := user
	result := db.Where("id = ?", userId).First(&user)

	if result.RowsAffected == 0 {
		return user, errors.New("user not found")
	}

	err := db.Model(&user).Where("id = ?", userId).Updates(models.User{Email: tempRequest.Email, Username: tempRequest.Username}).Error

	if err != nil {
		return user, errors.New(err.Error())
	}

	user.ID = uint(userId)
	user.Email = tempRequest.Email
	user.Username = tempRequest.Username

	return user, nil
}

func (userRepository *UserRepositoryImpl) DeleteUserByID(db *gorm.DB, user models.User) error {
	err := db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}