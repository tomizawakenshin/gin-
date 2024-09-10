package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	IconPhoto string
	Hanabis   []Hanabi  // Userが作成したスレッド（1対多の関係）
	Comments  []Comment // Userが書いたコメント（1対多の関係）
	Likes     []Like    // Userがつけたいいね（1対多の関係）
	Items     []Item    `gorm:"constraint:OnDelete:CASCADE"`
}
