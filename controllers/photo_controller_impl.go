package controllers

import (
	"final/helpers"
	"final/params"
	"final/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoControllerImpl struct {
	PhotoService services.PhotoService
}

func NewPhotoController(photoService services.PhotoService) PhotoController {
	return &PhotoControllerImpl{
		PhotoService: photoService,
	}
}

func (photoController *PhotoControllerImpl) CreatePhoto(ctx *gin.Context) {
	// panic("implement me")
	request := params.CreateUpdatePhoto{}
	helpers.ReadFromRequestBody(ctx, &request)

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	photo, err := photoController.PhotoService.CreatePhoto(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (photoController *PhotoControllerImpl) GetPhotos(ctx *gin.Context) {
	panic("implement me")
}

func (photoController *PhotoControllerImpl) UpdatePhoto(ctx *gin.Context) {
	panic("implement me")
}

func (photoController *PhotoControllerImpl) DeletePhoto(ctx *gin.Context) {
	panic("implement me")
}
