package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TryIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, isExist := ctx.Get("UserID")
		if isExist != true {
			ctx.JSON(500, gin.H{"error": "User ID is not exist"})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "index success",
			"UserID":  userID,
		})
	}
}
