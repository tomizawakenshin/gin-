package dto

type CreateHanabiInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Description string `json:"description" binding:"required,min=2,max=100"`
	PhotoURL    string `json:"photoURL" binding:"required"`
	Tag         string `json:"tag" binding:"required"`
}
