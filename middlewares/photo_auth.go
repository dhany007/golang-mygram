package middlewares

import (
	"final/helpers"
	"final/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		photoRepository := repositories.NewPhotoRepository()

		paramPhotoId, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			helpers.FailedMessageResponse(ctx, "invalid parameter photo id")
			return
		}

		tokenUserId := uint(ctx.MustGet("id").(float64))
		photo, err := photoRepository.GetPhotoById(db, paramPhotoId)
		if err != nil {
			helpers.FailedMessageResponse(ctx, err.Error())
			return
		}

		if tokenUserId != uint(photo.UserID) {
			helpers.FailedMessageResponse(ctx, "unauthorized")
			return
		}

		ctx.Next()
	}
}
