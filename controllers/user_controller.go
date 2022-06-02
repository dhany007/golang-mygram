package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}
