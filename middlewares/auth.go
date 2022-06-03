package middlewares

import (
	"final/helpers"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			helpers.FailedMessageResponse(ctx, "token not found")
			return
		}

		bearerToken := strings.HasPrefix(token, "Bearer")
		if !bearerToken {
			helpers.FailedMessageResponse(ctx, "no bearer token")
			return
		}

		tokenString := strings.Split(token, "Bearer ")[1]
		if tokenString == "" {
			helpers.FailedMessageResponse(ctx, "token not found after bearer")
			return
		}

		claims, err := helpers.VerifyToken(tokenString)
		if err != nil {
			helpers.FailedMessageResponse(ctx, err.Error())
		}

		data := claims.(jwt.MapClaims)

		ctx.Set("id", data["id"])
		ctx.Set("email", data["email"])

		ctx.Next()
	}
}

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
