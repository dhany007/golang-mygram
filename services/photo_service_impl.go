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

func NewPhotoService(db *gorm.DB, photoRepository repositories.PhotoRepository, validate *validator.Validate) PhotoService {
	return &PhotoServiceImpl{
		DB:              db,
		PhotoRepository: photoRepository,
		Validate:        validate,
	}
}

func (photoService *PhotoServiceImpl) CreatePhoto(photoParams params.CreateUpdatePhoto) (models.Photo, error) {
	photo := models.Photo{}

	errValidate := photoService.Validate.Struct(photoParams)
	if errValidate != nil {
		return photo, errors.New(errValidate.Error())
	}

	photo.Title = photoParams.Title
	photo.Caption = photoParams.Caption
	photo.PhotoUrl = photoParams.PhotoUrl
	photo.UserID = photoParams.UserID

	response, err := photoService.PhotoRepository.CreatePhoto(photoService.DB, photo)

	if err != nil {
		return photo, errors.New(err.Error())
	}

	return response, nil
}

func (photoService *PhotoServiceImpl) GetPhotos() ([]models.Photo, error) {
	photos := []models.Photo{}

	response, err := photoService.PhotoRepository.GetPhotos(photoService.DB)

	if err != nil {
		return photos, errors.New(err.Error())
	}

	return response, nil
}

func (photoService *PhotoServiceImpl) UpdatePhoto(photoParams params.CreateUpdatePhoto, photoId int) (models.Photo, error) {
	photo := models.Photo{}

	errRequest := photoService.Validate.Struct(photoParams)
	if errRequest != nil {
		return photo, errors.New(errRequest.Error())
	}

	photo.Caption = photoParams.Caption
	photo.Title = photoParams.Title
	photo.PhotoUrl = photoParams.PhotoUrl

	response, err := photoService.PhotoRepository.UpdatePhoto(photoService.DB, photo, photoId)

	if err != nil {
		return photo, errors.New(err.Error())
	}

	return response, nil
}

func (photoService *PhotoServiceImpl) DeletePhotoByID(photoId int) error {
	panic("implement me")
}
