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
	socialMedias := []models.SocialMedia{}

	result := db.Table("social_media").Scan(&socialMedias)
	if result.RowsAffected == 0 {
		return socialMedias, errors.New("socialmedias not found")
	}

	for i, p := range socialMedias {
		user := models.User{}
		err := db.Table("users").Select([]string{"id", "email", "username"}).Where("id = ?", p.UserID).Scan(&user).Error
		if err != nil {
			continue
		}
		socialMedias[i].User = &user
	}

	return socialMedias, nil
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
