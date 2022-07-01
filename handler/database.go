package handler

import (
	"coin-batam/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) JustTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"test": "success",
	})
}
