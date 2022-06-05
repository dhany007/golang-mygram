package repositories

import (
	"final/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(db *gorm.DB, comment models.Comment) (models.Comment, error)
	GetComments(db *gorm.DB) ([]models.Comment, error)
	UpdateComment(db *gorm.DB, comment models.Comment, commentId int) (models.Comment, error)
	DeleteComment(db *gorm.DB, comment models.Comment) error
}
