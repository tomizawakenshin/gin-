package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	UserID    uint   `gorm:"not null"` // 外部キーとしてUserIDを参照
	User      User   // リレーション（1対多の関係：Commentは1人のUserに属する）
	HanabiID  uint   // 外部キーとしてHanabiIDを追加
	Hanabi    Hanabi // リレーション（1対多の関係：Commentは1つのHanabiに属する）
	Likes     []Like // コメントについたいいね（1対多の関係）
	LikeCount uint
	HasLiked  bool
}
