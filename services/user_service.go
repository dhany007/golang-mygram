package services

import (
	"final/models"
	"final/params"
)

type UserService interface {
	CreateUser(userParams params.CreateUser) (models.User, error)
	LoginUser(userParams params.LoginUser) (string, error)
}
