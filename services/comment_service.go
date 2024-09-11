package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/reposotories"
)

type ICommentService interface {
	Create(createCommentInput dto.CreateCommentInput, userId uint, hanabiId uint) (*models.Comment, error)
}

type CommentService struct {
	repository reposotories.ICommentRepository
}

func NewCommentService(repository reposotories.ICommentRepository) ICommentService {
	return &CommentService{repository: repository}
}

func (s *CommentService) Create(createCommentInput dto.CreateCommentInput, userId uint, hanabiId uint) (*models.Comment, error) {
	newComment := models.Comment{
		Content:  createCommentInput.Content,
		UserID:   userId,
		HanabiID: hanabiId,
	}

	return s.repository.Create(newComment)
}
