package controllers

import (
	"final/helpers"
	"final/params"
	"final/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &UserControllerImpl{
		UserService: service,
	}
}

func (controller *UserControllerImpl) CreateUser(ctx *gin.Context) {
	request := params.CreateUser{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	user, err := controller.UserService.CreateUser(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (controller *UserControllerImpl) LoginUser(ctx *gin.Context) {
	request := params.LoginUser{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	token, err := controller.UserService.LoginUser(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (controller *UserControllerImpl) UpdateUser(ctx *gin.Context) {
	request := params.UpdateUser{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter user id")
		return
	}

	response, err := controller.UserService.UpdateUser(request, userId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *UserControllerImpl) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter user id")
		return
	}

	err = controller.UserService.DeleteUserByID(userId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your account has been successfully deleted")
}
