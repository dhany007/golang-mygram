package services

import (
	"errors"
	"final/models"
	"final/params"
	"final/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PhotoServiceImpl struct {
	DB              *gorm.DB
	PhotoRepository repositories.PhotoRepository
	Validate        *validator.Validate
}

func NewPhotoService(db *gorm.DB, repository repositories.PhotoRepository, validate *validator.Validate) PhotoService {
	return &PhotoServiceImpl{
		DB:              db,
		PhotoRepository: repository,
		Validate:        validate,
	}
}

func (service *PhotoServiceImpl) CreatePhoto(photoParams params.CreateUpdatePhoto) (models.Photo, error) {
	photo := models.Photo{}

	errValidate := service.Validate.Struct(photoParams)
	if errValidate != nil {
		return photo, errors.New(errValidate.Error())
	}

	photo.Title = photoParams.Title
	photo.Caption = photoParams.Caption
	photo.PhotoUrl = photoParams.PhotoUrl
	photo.UserID = photoParams.UserID

	response, err := service.PhotoRepository.CreatePhoto(service.DB, photo)

	if err != nil {
		return photo, errors.New(err.Error())
	}

	return response, nil
}

func (service *PhotoServiceImpl) GetPhotos() ([]models.Photo, error) {
	photos := []models.Photo{}

	response, err := service.PhotoRepository.GetPhotos(service.DB)

	if err != nil {
		return photos, errors.New(err.Error())
	}

	return response, nil
}

func (service *PhotoServiceImpl) UpdatePhoto(photoParams params.CreateUpdatePhoto, photoId int) (models.Photo, error) {
	photo := models.Photo{}

	errRequest := service.Validate.Struct(photoParams)
	if errRequest != nil {
		return photo, errors.New(errRequest.Error())
	}

	photo.Caption = photoParams.Caption
	photo.Title = photoParams.Title
	photo.PhotoUrl = photoParams.PhotoUrl

	response, err := service.PhotoRepository.UpdatePhoto(service.DB, photo, photoId)

	if err != nil {
		return photo, errors.New(err.Error())
	}

	return response, nil
}

func (service *PhotoServiceImpl) DeletePhotoByID(photoId int) error {
	photo := models.Photo{
		ID: uint(photoId),
	}

	err := service.PhotoRepository.DeletePhoto(service.DB, photo)
	if err != nil {
		return err
	}

	return nil
}
