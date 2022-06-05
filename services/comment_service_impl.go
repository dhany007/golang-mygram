package services

import (
	"errors"
	"final/models"
	"final/params"
	"final/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CommentServiceImpl struct {
	DB                *gorm.DB
	CommentRepository repositories.CommentRepository
	Validate          *validator.Validate
}

func NewCommentService(db *gorm.DB, commentRepository repositories.CommentRepository, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		DB:                db,
		CommentRepository: commentRepository,
		Validate:          validate,
	}
}

func (commentService *CommentServiceImpl) CreateComment(commentParam params.CreateComment) (models.Comment, error) {
	comment := models.Comment{}

	errValidate := commentService.Validate.Struct(commentParam)
	if errValidate != nil {
		return comment, errors.New(errValidate.Error())
	}

	comment.Message = commentParam.Message
	comment.UserID = commentParam.UserID
	comment.PhotoID = commentParam.PhotoID

	response, err := commentService.CommentRepository.CreateComment(commentService.DB, comment)
	if err != nil {
		return comment, errors.New(err.Error())
	}

	return response, nil
}

func (commentService *CommentServiceImpl) GetComments() ([]models.Comment, error) {
	comments := []models.Comment{}

	response, err := commentService.CommentRepository.GetComments(commentService.DB)

	if err != nil {
		return comments, errors.New(err.Error())
	}

	return response, nil
}

func (commentService *CommentServiceImpl) UpdateComment(commentParam params.UpdateComment, commentId int) (models.Comment, error) {
	panic("implement me")
}

func (commentService *CommentServiceImpl) DeleteCommentByID(commentId int) error {
	panic("implement me")
}
