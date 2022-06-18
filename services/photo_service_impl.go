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

func (s *PhotoServiceImpl) CreatePhoto(photoParams params.CreateUpdatePhoto) (models.Photo, error) {
	photo := models.Photo{}

	errValidate := s.Validate.Struct(photoParams)
	if errValidate != nil {
		return photo, errors.New(errValidate.Error())
	}

	photo.Title = photoParams.Title
	photo.Caption = photoParams.Caption
	photo.PhotoUrl = photoParams.PhotoUrl
	photo.UserID = photoParams.UserID

	response, err := s.PhotoRepository.CreatePhoto(s.DB, photo)

	if err != nil {
		return photo, errors.New(err.Error())
	}

	return response, nil
}

func (s *PhotoServiceImpl) GetPhotos() ([]models.Photo, error) {
	photos := []models.Photo{}

	response, err := s.PhotoRepository.GetPhotos(s.DB)

	if err != nil {
		return photos, errors.New(err.Error())
	}

	return response, nil
}

func (s *PhotoServiceImpl) UpdatePhoto(photoParams params.CreateUpdatePhoto, photoId int) (models.Photo, error) {
	photo := models.Photo{}

	errRequest := s.Validate.Struct(photoParams)
	if errRequest != nil {
		return photo, errors.New(errRequest.Error())
	}

	photo.Caption = photoParams.Caption
	photo.Title = photoParams.Title
	photo.PhotoUrl = photoParams.PhotoUrl

	response, err := s.PhotoRepository.UpdatePhoto(s.DB, photo, photoId)

	if err != nil {
		return photo, errors.New(err.Error())
	}

	return response, nil
}

func (s *PhotoServiceImpl) DeletePhotoByID(photoId int) error {
	photo := models.Photo{
		ID: uint(photoId),
	}

	err := s.PhotoRepository.DeletePhoto(s.DB, photo)
	if err != nil {
		return err
	}

	return nil
}
