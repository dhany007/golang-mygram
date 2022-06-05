package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func SuccessMessageResponse(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func FailedMessageResponse(ctx *gin.Context, err string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

func ReadFromRequestBody(ctx *gin.Context, result interface{}) bool {
	request := true
	err := ctx.ShouldBindJSON(result)
	if err != nil {
		FailedMessageResponse(ctx, err.Error())
		request = false
	}

	return request
}
