package services

import (
	"fmt"
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/reposotories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(createInputment dto.CreateItemInput, userId uint) (*models.Item, error)
	Restore(itemId uint) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput, userId uint) (*models.Item, error)
	Delete(itemId uint, userId uint) error
}

type ItemService struct {
	repository reposotories.IItemRepository
}

func NewItemService(repository reposotories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint, userId uint) (*models.Item, error) {
	return s.repository.FindById(itemId, userId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput, userId uint) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		Soldout:     false,
		UserID:      userId,
	}

	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput, userId uint) (*models.Item, error) {
	targetItem, err := s.FindById(itemId, userId)
	if err != nil {
		fmt.Println("指定のアイテムが見つかりませんでした")
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}

	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}

	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}

	if updateItemInput.Soldout != nil {
		targetItem.Soldout = *updateItemInput.Soldout
	}

	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint, userId uint) error {
	return s.repository.Delete(itemId, userId)
}

func (s *ItemService) Restore(itemId uint) (*models.Item, error) {
	return s.repository.Restore(itemId)
}
