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

func CommentRoutes(db *gorm.DB, router *gin.Engine) {
	validate := validator.New()

	commentRepository := repositories.NewCommentRepository()
	commentService := services.NewCommentService(db, commentRepository, validate)
	commentController := controllers.NewCommentController(commentService)

	commentRoutes := router.Group("/comments")
	{
		commentRoutes.Use(middlewares.Authentication())
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.GET("/", commentController.GetComments)
		commentRoutes.PUT("/:commentId", commentController.UpdateComment)
		commentRoutes.DELETE("/:commentId", commentController.DeleteCommentByID)
	}
}
