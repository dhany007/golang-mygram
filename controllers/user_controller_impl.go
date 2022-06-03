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

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController *UserControllerImpl) CreateUser(ctx *gin.Context) {
	request := params.CreateUser{}
	helpers.ReadFromRequestBody(ctx, &request)

	user, err := userController.UserService.CreateUser(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (userController *UserControllerImpl) LoginUser(ctx *gin.Context) {
	request := params.LoginUser{}
	helpers.ReadFromRequestBody(ctx, &request)

	token, err := userController.UserService.LoginUser(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (userController *UserControllerImpl) UpdateUser(ctx *gin.Context) {
	request := params.UpdateUser{}
	helpers.ReadFromRequestBody(ctx, &request)

	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter user id")
		return
	}

	response, err := userController.UserService.UpdateUser(request, userId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (userController *UserControllerImpl) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter user id")
		return
	}

	err = userController.UserService.DeleteUserByID(userId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your account has been successfully deleted")
}
