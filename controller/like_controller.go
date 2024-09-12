package controller

import (
	"gin-fleamarket/models"
	"gin-fleamarket/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ILikeController interface {
	Like(ctx *gin.Context)
	Unlike(ctx *gin.Context)
}

type LikeController struct {
	services services.ILikeService
}

func NewLikeController(service services.ILikeService) ILikeController {
	return &LikeController{services: service}
}

func (c *LikeController) Like(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*models.User).ID
	commentId, err := strconv.ParseUint(ctx.Param("commentId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id "})
		return
	}

	err = c.services.Like(uint(userId), uint(commentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Liked Successfully"})
}

func (c *LikeController) Unlike(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*models.User).ID
	commentId, err := strconv.ParseUint(ctx.Param("commentId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id "})
		return
	}

	err = c.services.Unlike(uint(userId), uint(commentId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Unliked Successfully"})
}
