package helpers

import (
	"final/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func SuccessMessageResponse(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, models.Response{
		Message: message,
	})
}

func FailedMessageResponse(ctx *gin.Context, err string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
		Error: err,
	})
}

func ReadFromRequestBody(ctx *gin.Context, result interface{}) {
	err := ctx.ShouldBindJSON(result)
	if err != nil {
		FailedMessageResponse(ctx, err.Error())
		return
	}
}
