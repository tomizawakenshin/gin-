package reposotories

import (
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(newItem models.Comment) (*models.Comment, error)
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentMemoryRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(newComment models.Comment) (*models.Comment, error) {
	result := r.db.Create(&newComment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newComment, nil
}
