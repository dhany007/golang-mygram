package repositories

import (
	"final/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(db *gorm.DB, photo models.Photo) (models.Photo, error)
	GetPhotos(db *gorm.DB) ([]models.Photo, error)
	UpdatePhoto(db *gorm.DB, photo models.Photo, photoId int) (models.Photo, error)
	DeletePhoto(db *gorm.DB, photo models.Photo) error
	GetPhotoById(db *gorm.DB, photoId int) (models.Photo, error)
}
