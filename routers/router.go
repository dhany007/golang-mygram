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

func StartEngine(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	validate := validator.New()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.CreateUser)
		userRouter.POST("/login", userController.LoginUser)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/:userId", middlewares.UserAuthorization(), userController.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.UserAuthorization(), userController.DeleteUser)
	}

	photoRepository := repositories.NewPhotoRepository()
	photoService := services.NewPhotoService(db, photoRepository, validate)
	photoController := controllers.NewPhotoController(photoService)

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.GET("/", photoController.GetPhotos)
		photoRouter.PUT("/:photoId", photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoId", photoController.DeletePhoto)
	}

	router.Use(gin.Recovery())

	return router
}
