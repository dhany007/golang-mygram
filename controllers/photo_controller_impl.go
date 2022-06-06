package controllers

import (
	"final/helpers"
	"final/params"
	"final/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoControllerImpl struct {
	PhotoService services.PhotoService
}

func NewPhotoController(service services.PhotoService) PhotoController {
	return &PhotoControllerImpl{
		PhotoService: service,
	}
}

func (controller *PhotoControllerImpl) CreatePhoto(ctx *gin.Context) {
	request := params.CreateUpdatePhoto{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	photo, err := controller.PhotoService.CreatePhoto(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, photo)
}

func (controller *PhotoControllerImpl) GetPhotos(ctx *gin.Context) {
	photos, err := controller.PhotoService.GetPhotos()
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
	}

	ctx.JSON(http.StatusOK, photos)
}

func (controller *PhotoControllerImpl) UpdatePhoto(ctx *gin.Context) {
	request := params.CreateUpdatePhoto{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter photo id")
		return
	}

	photo, err := controller.PhotoService.UpdatePhoto(request, photoId)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (controller *PhotoControllerImpl) DeletePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter photo id")
		return
	}

	err = controller.PhotoService.DeletePhotoByID(photoId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your photo has been successfully deleted")
}
