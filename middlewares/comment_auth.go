package middlewares

import (
	"final/helpers"
	"final/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentRepository := repositories.NewCommentRepository()

		paramCommentId, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			helpers.FailedMessageResponse(ctx, "invalid parameter comment id")
			return
		}

		tokenUserId := uint(ctx.MustGet("id").(float64))
		comment, err := commentRepository.GetCommentById(db, paramCommentId)
		if err != nil {
			helpers.FailedMessageResponse(ctx, err.Error())
			return
		}

		if tokenUserId != uint(comment.UserID) {
			helpers.FailedMessageResponse(ctx, "unauthorized")
			return
		}

		ctx.Next()
	}
}
