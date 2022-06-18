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

func NewCommentController(service services.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: service,
	}
}

func (c *CommentControllerImpl) CreateComment(ctx *gin.Context) {
	request := params.CreateComment{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	userId := ctx.MustGet("id").(float64)
	request.UserID = uint(userId)

	comment, err := c.CommentService.CreateComment(request)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

func (c *CommentControllerImpl) GetComments(ctx *gin.Context) {
	comments, err := c.CommentService.GetComments()
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c *CommentControllerImpl) UpdateComment(ctx *gin.Context) {
	request := params.UpdateComment{}
	requestValid := helpers.ReadFromRequestBody(ctx, &request)
	if !requestValid {
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter comment id")
		return
	}

	comment, err := c.CommentService.UpdateComment(request, commentId)
	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (c *CommentControllerImpl) DeleteCommentByID(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		helpers.FailedMessageResponse(ctx, "invalid parameter comment id")
		return
	}

	err = c.CommentService.DeleteCommentByID(commentId)

	if err != nil {
		helpers.FailedMessageResponse(ctx, err.Error())
		return
	}

	helpers.SuccessMessageResponse(ctx, "Your comment has been successfully deleted")
}
