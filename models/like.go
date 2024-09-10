package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID    uint    `gorm:"not null"` // 外部キーとしてUserIDを参照
	User      User    // リレーション（1対多の関係：Likeは1人のUserに属する）
	CommentID uint    `gorm:"not null"` // 外部キーとしてCommentIDを参照
	Comment   Comment // リレーション（1対多の関係：Likeは1つのCommentに属する）
}
