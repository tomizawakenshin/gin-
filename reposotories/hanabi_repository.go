package reposotories

import (
	"errors"
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type IHanabiRepository interface {
	FindAll() (*[]models.Hanabi, error)
	FindByID(hanabiID uint) (*models.Hanabi, error)
	Create(newItem models.Hanabi) (*models.Hanabi, error)
	PreloadUser(hanabi *models.Hanabi) error
}

type HanabiRepository struct {
	db *gorm.DB
}

func NewHanabiRepository(db *gorm.DB) IHanabiRepository {
	return &HanabiRepository{db: db}
}

func (r *HanabiRepository) FindAll() (*[]models.Hanabi, error) {
	var hanabis []models.Hanabi
	// created_at カラムで降順に並べ替える
	result := r.db.Preload("User").Order("created_at DESC").Find(&hanabis)
	if result.Error != nil {
		return nil, result.Error
	}
	return &hanabis, nil
}

func (r *HanabiRepository) FindByID(hanabiID uint) (*models.Hanabi, error) {
	var hanabi models.Hanabi
	result := r.db.Preload("User").First(&hanabi, "id = ?", hanabiID)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("hanabi not found")
		}
		return nil, result.Error
	}
	return &hanabi, nil
}

func (r *HanabiRepository) Create(newItem models.Hanabi) (*models.Hanabi, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *HanabiRepository) PreloadUser(hanabi *models.Hanabi) error {
	return r.db.Preload("User").First(hanabi).Error
}
