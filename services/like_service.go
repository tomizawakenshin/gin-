package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/reposotories"
)

type ILikeService interface {
	Like(userID uint, commentID uint) error
	Unlike(userID uint, commentID uint) error
}

type LikeService struct {
	likeRepository reposotories.ILikeRepository
}

func NewLikeService(likeRepository reposotories.ILikeRepository) ILikeService {
	return &LikeService{
		likeRepository: likeRepository,
	}
}

func (s *LikeService) Like(userID uint, commentID uint) error {
	like := models.Like{
		UserID:    userID,
		CommentID: commentID,
	}
	return s.likeRepository.CreateLike(like)
}

func (s *LikeService) Unlike(userID uint, commentID uint) error {
	return s.likeRepository.DeleteLike(userID, commentID)
}
