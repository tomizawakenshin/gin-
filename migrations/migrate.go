package main

import (
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Hanabi{}, &models.Like{}); err != nil {
		panic("Failed to migrate db")
	}
}
