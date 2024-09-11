package controller

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICommentController interface {
	Create(ctx *gin.Context)
}

type CommentController struct {
	service services.ICommentService
}

func NewCommentController(service services.ICommentService) ICommentController {
	return &CommentController{service: service}
}

func (c *CommentController) Create(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var input dto.CreateCommentInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := user.(*models.User).ID

	hanabiId, err := strconv.ParseUint(ctx.Param("hanabiId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id "})
		return
	}

	newComment, err := c.service.Create(input, userId, uint(hanabiId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newComment})

}
