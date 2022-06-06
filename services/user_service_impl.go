package services

import (
	"errors"
	"final/helpers"
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

func NewUserService(repository repositories.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) CreateUser(userParams params.CreateUser) (models.User, error) {
	user := models.User{}

	err := service.Validate.Struct(userParams)
	if err != nil {
		return user, errors.New(err.Error())
	}

	user.Age = int(userParams.Age)
	user.Email = userParams.Email
	user.Password = userParams.Password
	user.Username = userParams.Username

	response, err := service.UserRepository.CreateUser(service.DB, user)
	if err != nil {
		return user, errors.New(err.Error())
	}

	return response, nil
}

func (service *UserServiceImpl) LoginUser(userParams params.LoginUser) (string, error) {
	err := service.Validate.Struct(userParams)
	if err != nil {
		return "", errors.New(err.Error())
	}

	user := models.User{
		Email: userParams.Email,
	}

	response, err := service.UserRepository.LoginUser(service.DB, user)
	if err != nil {
		return "", errors.New(err.Error())
	}

	passwordUser := response.Password

	comparePassword := helpers.ComparePassword([]byte(passwordUser), []byte(userParams.Password))
	if !comparePassword {
		return "", errors.New("password not match")
	}

	token := helpers.GenerateToken(int(response.ID), user.Email)

	return token, nil
}

func (service *UserServiceImpl) UpdateUser(userParams params.UpdateUser, userId int) (models.User, error) {
	user := models.User{}

	errRequest := service.Validate.Struct(userParams)
	if errRequest != nil {
		return user, errors.New(errRequest.Error())
	}

	user.Email = userParams.Email
	user.Username = userParams.Username

	response, err := service.UserRepository.UpdateUser(service.DB, user, userId)

	if err != nil {
		return user, errors.New(err.Error())
	}

	return response, nil
}

func (service *UserServiceImpl) DeleteUserByID(userId int) error {
	user := models.User{
		ID: uint(userId),
	}

	err := service.UserRepository.DeleteUserByID(service.DB, user)
	if err != nil {
		return err
	}

	return nil
}
