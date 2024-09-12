package reposotories

import (
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	CreateLike(like models.Like) error
	DeleteLike(userID uint, commentID uint) error
}

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) ILikeRepository {
	return &LikeRepository{db: db}
}

func (r *LikeRepository) CreateLike(like models.Like) error {
	return r.db.Create(&like).Error
}

func (r *LikeRepository) DeleteLike(userID uint, commentID uint) error {
	return r.db.Where("user_id = ? AND comment_id = ?", userID, commentID).Delete(&models.Like{}).Error
}
