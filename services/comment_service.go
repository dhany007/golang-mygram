package services

import (
	"final/models"
	"final/params"
)

type CommentService interface {
	CreateComment(commentParam params.CreateComment) (models.Comment, error)
	GetComments() ([]models.Comment, error)
	UpdateComment(commentParam params.UpdateComment, commentId int) (models.Comment, error)
	DeleteCommentByID(commentId int) error
}
