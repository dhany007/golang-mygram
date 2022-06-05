package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	UserRoutes(db, router)
	PhotoRoutes(db, router)
	CommentRoutes(db, router)

	router.Use(gin.Recovery())

	return router
}
