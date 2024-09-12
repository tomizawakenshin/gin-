package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/reposotories"
)

type IHanabiService interface {
	FindAll(date string) (*[]models.Hanabi, error)
	FindByID(hanabiID uint, userID uint) (*models.Hanabi, error)
	Create(createInputment dto.CreateHanabiInput, userId uint) (*models.Hanabi, error)
	PreloadUser(hanabi *models.Hanabi) error
}

type HanabiService struct {
	repository reposotories.IHanabiRepository
}

func NewHanabiService(repository reposotories.IHanabiRepository) IHanabiService {
	return &HanabiService{repository: repository}
}

func (s *HanabiService) FindAll(date string) (*[]models.Hanabi, error) {
	return s.repository.FindAll(date)
}
func (s *HanabiService) FindByID(itemID uint, userID uint) (*models.Hanabi, error) {
	return s.repository.FindByID(itemID, userID)
}
func (s *HanabiService) Create(createHanabiInput dto.CreateHanabiInput, userId uint) (*models.Hanabi, error) {
	newHanabi := models.Hanabi{
		Name:         createHanabiInput.Name,
		Description:  createHanabiInput.Description,
		Photo:        createHanabiInput.PhotoURL,
		UserID:       userId,
		Tag:          createHanabiInput.Tag,
		CommentCount: 0,
	}

	return s.repository.Create(newHanabi)
}

func (s *HanabiService) PreloadUser(hanabi *models.Hanabi) error {
	return s.repository.PreloadUser(hanabi)
}
