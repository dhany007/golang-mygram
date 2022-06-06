package middlewares

import (
	"final/helpers"
	"final/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SocialMediaAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		socialMediaRepository := repositories.NewSocialMediaRepository()

		paramSocialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))
		if err != nil {
			helpers.FailedMessageResponse(ctx, "invalid parameter socialmedia id")
			return
		}

		tokenUserId := uint(ctx.MustGet("id").(float64))
		socialMedia, err := socialMediaRepository.GetSocialMediaById(db, paramSocialMediaId)
		if err != nil {
			helpers.FailedMessageResponse(ctx, err.Error())
			return
		}

		if tokenUserId != uint(socialMedia.UserID) {
			helpers.FailedMessageResponse(ctx, "unauthorized")
			return
		}

		ctx.Next()
	}
}
