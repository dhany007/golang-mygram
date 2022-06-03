package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}
