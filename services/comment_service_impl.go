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

func NewCommentService(db *gorm.DB, repository repositories.CommentRepository, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		DB:                db,
		CommentRepository: repository,
		Validate:          validate,
	}
}

func (service *CommentServiceImpl) CreateComment(commentParam params.CreateComment) (models.Comment, error) {
	comment := models.Comment{}

	errValidate := service.Validate.Struct(commentParam)
	if errValidate != nil {
		return comment, errors.New(errValidate.Error())
	}

	comment.Message = commentParam.Message
	comment.UserID = commentParam.UserID
	comment.PhotoID = commentParam.PhotoID

	response, err := service.CommentRepository.CreateComment(service.DB, comment)
	if err != nil {
		return comment, errors.New(err.Error())
	}

	return response, nil
}

func (service *CommentServiceImpl) GetComments() ([]models.Comment, error) {
	comments := []models.Comment{}

	response, err := service.CommentRepository.GetComments(service.DB)

	if err != nil {
		return comments, errors.New(err.Error())
	}

	return response, nil
}

func (service *CommentServiceImpl) UpdateComment(commentParam params.UpdateComment, commentId int) (models.Comment, error) {
	comment := models.Comment{}

	errRequest := service.Validate.Struct(commentParam)
	if errRequest != nil {
		return comment, errors.New(errRequest.Error())
	}

	comment.Message = commentParam.Message

	response, err := service.CommentRepository.UpdateComment(service.DB, comment, commentId)

	if err != nil {
		return comment, errors.New(err.Error())
	}

	return response, nil
}

func (service *CommentServiceImpl) DeleteCommentByID(commentId int) error {
	comment := models.Comment{
		ID: uint(commentId),
	}

	err := service.CommentRepository.DeleteComment(service.DB, comment)
	if err != nil {
		return err
	}

	return nil
}
