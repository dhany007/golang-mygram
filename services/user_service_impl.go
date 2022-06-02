package services

import (
	"errors"
	"final/models"
	"final/params"
	"final/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repositories.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (userService *UserServiceImpl) CreateUser(userParams params.CreateUser) (models.User, error) {
	err := userService.Validate.Struct(userParams)

	user := models.User{}

	if err != nil {
		return user, errors.New(err.Error())
	}

	user.Age = int(userParams.Age)
	user.Email = userParams.Email
	user.Password = userParams.Password
	user.Username = userParams.Username

	response, err := userService.UserRepository.CreateUser(userService.DB, user)
	if err != nil {
		return user, errors.New(err.Error())
	}

	return response, nil
}
