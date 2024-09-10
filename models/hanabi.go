package models

import (
	"gorm.io/gorm"
)

type Hanabi struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Description  string
	Photo        string
	UserID       uint      `gorm:"not null"` // 外部キーとしてUserIDを参照
	User         User      // リレーションを明示（1対多の関係：Threadは1人のUserに属する）
	Tag          string    `gorm:"not null"`
	CommentCount uint      `gorm:"default:0"` // コメント数のカウント
	Comments     []Comment // Threadに関連するコメント（1対多の関係）
}
