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

func PhotoRoutes(db *gorm.DB, router *gin.Engine) {
	validate := validator.New()

	photoRepository := repositories.NewPhotoRepository()
	photoService := services.NewPhotoService(db, photoRepository, validate)
	photoController := controllers.NewPhotoController(photoService)

	photoRoutes := router.Group("/photos")
	{
		photoRoutes.Use(middlewares.Authentication())
		photoRoutes.POST("/", photoController.CreatePhoto)
		photoRoutes.GET("/", photoController.GetPhotos)
		photoRoutes.PUT("/:photoId", middlewares.PhotoAuthorization(db), photoController.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", middlewares.PhotoAuthorization(db), photoController.DeletePhoto)
	}
}
