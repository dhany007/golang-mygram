package controllers

import "github.com/gin-gonic/gin"

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}
