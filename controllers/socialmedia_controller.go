package controllers

import "github.com/gin-gonic/gin"

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}
