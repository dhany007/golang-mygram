package controllers

import "github.com/gin-gonic/gin"

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetComments(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteCommentByID(ctx *gin.Context)
}
