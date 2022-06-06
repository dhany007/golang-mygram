package routers

import (
	"final/controllers"
	"final/middlewares"
	"final/repositories"
	"final/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SocialMediaRoutes(db *gorm.DB, router *gin.Engine) {
	validate := validator.New()

	socialMediaRepository := repositories.NewSocialMediaRepository()
	socialMediaService := services.NewSocialMediaService(validate, socialMediaRepository, db)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	socialMediaRoutes := router.Group("/socialmedias")
	{
		socialMediaRoutes.Use(middlewares.Authentication())
		socialMediaRoutes.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaRoutes.GET("/", socialMediaController.GetSocialMedias)
		socialMediaRoutes.PUT("/:socialMediaId", socialMediaController.UpdateSocialMedia)
		socialMediaRoutes.DELETE("/:socialMediaId", socialMediaController.DeleteSocialMedia)
	}
}
