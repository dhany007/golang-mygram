package repositories

import (
	"final/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(db *gorm.DB, socialMedia models.SocialMedia) (models.SocialMedia, error)
	GetSocialMedias(db *gorm.DB) ([]models.SocialMedia, error)
	UpdateSocialMedia(db *gorm.DB) (models.SocialMedia, error)
	DeleteSocialMedia(db *gorm.DB) error
	GetSocialMediaById(db *gorm.DB) (models.SocialMedia, error)
}
