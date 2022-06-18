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

func (c *PhotoControllerImpl) CreatePhoto(ctx *gin.Context) {
	request := params.CreateUpdatePhoto{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	photo, err := c.PhotoService.CreatePhoto(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, photo)
}

func (c *PhotoControllerImpl) GetPhotos(ctx *gin.Context) {
	photos, err := c.PhotoService.GetPhotos()
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
	}

	ctx.JSON(http.StatusOK, photos)
}

func (c *PhotoControllerImpl) UpdatePhoto(ctx *gin.Context) {
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

	photo, err := c.PhotoService.UpdatePhoto(request, photoId)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (c *PhotoControllerImpl) DeletePhoto(ctx *gin.Context) {
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter photo id")
		return
	}

	err = c.PhotoService.DeletePhotoByID(photoId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your photo has been successfully deleted")
}
