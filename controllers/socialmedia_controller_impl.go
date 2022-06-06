package controllers

import (
	"final/helpers"
	"final/params"
	"final/services"
	"net/http"
	"strconv"

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
	socialMedias, err := controller.SocialMediaService.GetSocialMedias()
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"social_medias": socialMedias,
	})
}

func (controller *SocialMediaControllerImpl) UpdateSocialMedia(ctx *gin.Context) {
	request := params.CreateUpdateSocialMedia{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter socialmedia id")
		return
	}

	socialMedia, err := controller.SocialMediaService.UpdateSocialMedias(request, socialMediaId)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func (controller *SocialMediaControllerImpl) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter socialmedia id")
		return
	}

	err = controller.SocialMediaService.DeleteSocialMediasByID(socialMediaId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your socialmedia has been successfully deleted")
}
