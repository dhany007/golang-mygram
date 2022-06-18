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

func (s *CommentServiceImpl) CreateComment(commentParam params.CreateComment) (models.Comment, error) {
	comment := models.Comment{}

	errValidate := s.Validate.Struct(commentParam)
	if errValidate != nil {
		return comment, errors.New(errValidate.Error())
	}

	comment.Message = commentParam.Message
	comment.UserID = commentParam.UserID
	comment.PhotoID = commentParam.PhotoID

	response, err := s.CommentRepository.CreateComment(s.DB, comment)
	if err != nil {
		return comment, errors.New(err.Error())
	}

	return response, nil
}

func (s *CommentServiceImpl) GetComments() ([]models.Comment, error) {
	comments := []models.Comment{}

	response, err := s.CommentRepository.GetComments(s.DB)

	if err != nil {
		return comments, errors.New(err.Error())
	}

	return response, nil
}

func (s *CommentServiceImpl) UpdateComment(commentParam params.UpdateComment, commentId int) (models.Comment, error) {
	comment := models.Comment{}

	errRequest := s.Validate.Struct(commentParam)
	if errRequest != nil {
		return comment, errors.New(errRequest.Error())
	}

	comment.Message = commentParam.Message

	response, err := s.CommentRepository.UpdateComment(s.DB, comment, commentId)

	if err != nil {
		return comment, errors.New(err.Error())
	}

	return response, nil
}

func (s *CommentServiceImpl) DeleteCommentByID(commentId int) error {
	comment := models.Comment{
		ID: uint(commentId),
	}

	err := s.CommentRepository.DeleteComment(s.DB, comment)
	if err != nil {
		return err
	}

	return nil
}
