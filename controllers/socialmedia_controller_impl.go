package controllers

import (
	"final/helpers"
	"final/params"
	"final/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialMediaControllerImpl struct {
	SocialMediaService services.SocialMediaService
}

func NewSocialMediaController(service services.SocialMediaService) SocialMediaController {
	return &SocialMediaControllerImpl{
		SocialMediaService: service,
	}
}

func (controller *SocialMediaControllerImpl) CreateSocialMedia(ctx *gin.Context) {
	request := params.CreateUpdateSocialMedia{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	socialMedia, err := controller.SocialMediaService.CreateSocialMedia(request)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, socialMedia)
}

func (controller *SocialMediaControllerImpl) GetSocialMedias(ctx *gin.Context) {
	panic("implement me")
}

func (controller *SocialMediaControllerImpl) UpdateSocialMedia(ctx *gin.Context) {
	panic("implement me")
}

func (controller *SocialMediaControllerImpl) DeleteSocialMedia(ctx *gin.Context) {
	panic("implement me")
}
