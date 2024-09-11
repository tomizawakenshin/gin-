package dto

type CreateCommentInput struct {
	Content string `json:"content" binding:"required,max=100"`
}
