package services

import (
	"final/models"
	"final/params"
)

type SocialMediaService interface {
	CreateSocialMedia(socialMediaParams params.CreateUpdateSocialMedia) (models.SocialMedia, error)
	GetSocialMedias() ([]models.SocialMedia, error)
	UpdateSocialMedias(socialMediaParam params.CreateUpdateSocialMedia, socialMediaId int) (models.SocialMedia, error)
	DeleteSocialMediasByID(socialMediaId int) error
}
