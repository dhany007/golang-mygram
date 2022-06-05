package repositories

import (
	"errors"
	"final/models"

	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (commentRepository *CommentRepositoryImpl) CreateComment(db *gorm.DB, comment models.Comment) (models.Comment, error) {
	err := db.Create(&comment).Error
	if err != nil {
		return comment, errors.New(err.Error())
	}

	commentCreated := models.Comment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}

	return commentCreated, nil
}
func (commentRepository *CommentRepositoryImpl) GetComments(db *gorm.DB) ([]models.Comment, error) {
	panic("implement me")
}
func (commentRepository *CommentRepositoryImpl) UpdateComment(db *gorm.DB, comment models.Comment, commentId int) (models.Comment, error) {
	panic("implement me")
}
func (commentRepository *CommentRepositoryImpl) DeleteComment(db *gorm.DB, comment models.Comment) error {
	panic("implement me")
}
