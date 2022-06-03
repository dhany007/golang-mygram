package repositories

import (
	"errors"
	"final/models"

	"gorm.io/gorm"
)

type PhotoRepositoryImpl struct {
}

func NewPhotoRepository() PhotoRepository {
	return &PhotoRepositoryImpl{}
}

func (photoRepsitory *PhotoRepositoryImpl) CreatePhoto(db *gorm.DB, photo models.Photo) (models.Photo, error) {
	err := db.Create(&photo).Error
	if err != nil {
		return photo, errors.New(err.Error())
	}

	return photo, nil
}

func (photoRepsitory *PhotoRepositoryImpl) GetPhotos(db *gorm.DB) ([]models.Photo, error) {
	panic("implement me")
}

func (photoRepsitory *PhotoRepositoryImpl) UpdatePhoto(db *gorm.DB, photo models.Photo, photoId int) (models.Photo, error) {
	panic("implement me")
}

func (photoRepsitory *PhotoRepositoryImpl) DeletePhoto(db *gorm.DB, photo models.Photo) error {
	panic("implement me")
}