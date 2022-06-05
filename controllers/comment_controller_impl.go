package controllers

import (
	"final/helpers"
	"final/params"
	"final/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentControllerImpl struct {
	CommentService services.CommentService
}

func NewCommentController(commentService services.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (commentController *CommentControllerImpl) CreateComment(ctx *gin.Context) {
	request := params.CreateComment{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	comment, err := commentController.CommentService.CreateComment(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (commentController *CommentControllerImpl) GetComments(ctx *gin.Context) {
	comments, err := commentController.CommentService.GetComments()
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
	}

	ctx.JSON(http.StatusOK, comments)
}

func (commentController *CommentControllerImpl) UpdateComment(ctx *gin.Context) {
	request := params.UpdateComment{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter photo id")
		return
	}

	comment, err := commentController.CommentService.UpdateComment(request, commentId)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (commentController *CommentControllerImpl) DeleteCommentByID(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter comment id")
		return
	}

	err = commentController.CommentService.DeleteCommentByID(commentId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your comment has been successfully deleted")
}
