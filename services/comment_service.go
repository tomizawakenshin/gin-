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
	repository       reposotories.ICommentRepository
	hanabiRepository reposotories.IHanabiRepository
}

func NewCommentService(repository reposotories.ICommentRepository, hanabiRepository reposotories.IHanabiRepository) ICommentService {
	return &CommentService{
		repository:       repository,
		hanabiRepository: hanabiRepository, // 初期化
	}
}

func (s *CommentService) Create(createCommentInput dto.CreateCommentInput, userId uint, hanabiId uint) (*models.Comment, error) {
	newComment := models.Comment{
		Content:  createCommentInput.Content,
		UserID:   userId,
		HanabiID: hanabiId,
	}

	// ポインタ型でCreate関数を呼び出し
	newCommentPtr, err := s.repository.Create(newComment)
	if err != nil {
		return nil, err
	}

	// コメント作成後に Hanabi の comment_count をインクリメント
	err = s.hanabiRepository.IncrementCommentCount(hanabiId)
	if err != nil {
		return nil, err
	}

	return newCommentPtr, nil // ポインタ型を返す
}
