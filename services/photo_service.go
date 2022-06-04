package services

import (
	"final/models"
	"final/params"
)

type PhotoService interface {
	CreatePhoto(photoParams params.CreateUpdatePhoto) (models.Photo, error)
	GetPhotos() ([]models.Photo, error)
	UpdatePhoto(photoParams params.CreateUpdatePhoto, photoId int) (models.Photo, error)
	DeletePhotoByID(photoId int) error
}
