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
	IncrementCommentCount(hanabiId uint) error
}

type HanabiRepository struct {
	db *gorm.DB
}

func NewHanabiRepository(db *gorm.DB) IHanabiRepository {
	return &HanabiRepository{db: db}
}

func (r *HanabiRepository) FindAll() (*[]models.Hanabi, error) {
	var hanabis []models.Hanabi

	// Hanabiをcreated_atで降順に並べ替え
	//result := r.db.Preload("User").Order("created_at DESC").Find(&hanabis)
	result := r.db.Order("created_at DESC").Find(&hanabis)
	if result.Error != nil {
		return nil, result.Error
	}

	// 各Hanabiに対してCommentCountを計算
	for i := range hanabis {
		var commentCount int64
		result = r.db.Model(&models.Comment{}).Where("hanabi_id = ?", hanabis[i].ID).Count(&commentCount)
		if result.Error != nil {
			return nil, errors.New("コメント数の取得に失敗しました")
		}
		hanabis[i].CommentCount = uint(commentCount)
	}

	return &hanabis, nil
}

func (r *HanabiRepository) FindByID(hanabiID uint) (*models.Hanabi, error) {
	var hanabi models.Hanabi
	result := r.db.Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		First(&hanabi, "id = ?", hanabiID)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("hanabi not found")
		}
		return nil, result.Error
	}

	var commentCount int64
	result = r.db.Model(&models.Comment{}).Where("hanabi_id = ?", hanabiID).Count(&commentCount)
	if result.Error != nil {
		return nil, errors.New("指定されたhanabiのコメントが取得できませんでした")
	}
	hanabi.CommentCount = uint(commentCount)

	return &hanabi, nil
}

// ここに IncrementCommentCount 関数を追加します
func (r *HanabiRepository) IncrementCommentCount(hanabiId uint) error {
	return r.db.Model(&models.Hanabi{}).Where("id = ?", hanabiId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
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
