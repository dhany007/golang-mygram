package services

import (
	"errors"
	"final/models"
	"final/params"
	"final/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMediaServiceImpl struct {
	Validate              *validator.Validate
	SocialMediaRepository repositories.SocialMediaRepository
	DB                    *gorm.DB
}

func NewSocialMediaService(validate *validator.Validate, repository repositories.SocialMediaRepository, db *gorm.DB) SocialMediaService {
	return &SocialMediaServiceImpl{
		Validate:              validate,
		SocialMediaRepository: repository,
		DB:                    db,
	}
}

func (s *SocialMediaServiceImpl) CreateSocialMedia(socialMediaParams params.CreateUpdateSocialMedia) (models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}

	errValidate := s.Validate.Struct(socialMediaParams)
	if errValidate != nil {
		return socialMedia, errors.New(errValidate.Error())
	}

	socialMedia.Name = socialMediaParams.Name
	socialMedia.SocialMediaUrl = socialMediaParams.SocialMediaUrl
	socialMedia.UserID = socialMediaParams.UserID

	response, err := s.SocialMediaRepository.CreateSocialMedia(s.DB, socialMedia)

	if err != nil {
		return socialMedia, errors.New(err.Error())
	}

	return response, nil
}

func (s *SocialMediaServiceImpl) GetSocialMedias() ([]models.SocialMedia, error) {
	socialMedias := []models.SocialMedia{}

	response, err := s.SocialMediaRepository.GetSocialMedias(s.DB)

	if err != nil {
		return socialMedias, errors.New(err.Error())
	}

	return response, nil
}

func (s *SocialMediaServiceImpl) UpdateSocialMedias(socialMediaParam params.CreateUpdateSocialMedia, socialMediaId int) (models.SocialMedia, error) {
	socialMedia := models.SocialMedia{}

	errRequest := s.Validate.Struct(socialMediaParam)
	if errRequest != nil {
		return socialMedia, errors.New(errRequest.Error())
	}

	socialMedia.Name = socialMediaParam.Name
	socialMedia.SocialMediaUrl = socialMediaParam.SocialMediaUrl
	socialMedia.UserID = socialMediaParam.UserID

	response, err := s.SocialMediaRepository.UpdateSocialMedia(s.DB, socialMedia, socialMediaId)

	if err != nil {
		return socialMedia, errors.New(err.Error())
	}

	return response, nil
}

func (s *SocialMediaServiceImpl) DeleteSocialMediasByID(socialMediaId int) error {
	socialMedia := models.SocialMedia{
		ID: uint(socialMediaId),
	}

	err := s.SocialMediaRepository.DeleteSocialMedia(s.DB, socialMedia)
	if err != nil {
		return err
	}

	return nil
}
