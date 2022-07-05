package handler

import (
	"coin-batam/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddTransaction(c *gin.Context) {

	var transaction entities.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	err := h.Service.AddTransaction(transaction)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
}
