package repositories

import (
	"errors"
	"final/models"

	"gorm.io/gorm"
)

type SocialMediaRepositoryImpl struct {
}

func NewSocialMediaRepository() SocialMediaRepository {
	return &SocialMediaRepositoryImpl{}
}

func (repository *SocialMediaRepositoryImpl) CreateSocialMedia(db *gorm.DB, socialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := db.Create(&socialMedia).Error
	if err != nil {
		return socialMedia, errors.New(err.Error())
	}

	socialMediaCreated := models.SocialMedia{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt,
	}

	return socialMediaCreated, nil
}

func (repository *SocialMediaRepositoryImpl) GetSocialMedias(db *gorm.DB) ([]models.SocialMedia, error) {
	panic("implement me")
}

func (repository *SocialMediaRepositoryImpl) UpdateSocialMedia(db *gorm.DB) (models.SocialMedia, error) {
	panic("implement me")
}

func (repository *SocialMediaRepositoryImpl) DeleteSocialMedia(db *gorm.DB) error {
	panic("implement me")
}

func (repository *SocialMediaRepositoryImpl) GetSocialMediaById(db *gorm.DB) (models.SocialMedia, error) {
	panic("implement me")
}
