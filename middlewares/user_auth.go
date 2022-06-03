package middlewares

import (
	"final/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paramUserId, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil {
			helpers.FailedMessageResponse(ctx, "invalid parameter user id")
			return
		}
		tokenUserId := uint(ctx.MustGet("id").(float64))

		if tokenUserId != uint(paramUserId) {
			helpers.FailedMessageResponse(ctx, "unauthorized")
			return
		}

		ctx.Next()
	}
}
