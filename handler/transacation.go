package handler

import (
	"coin-batam/entities"
	"net/http"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddTransaction(c *gin.Context) {
	claims, err := ValidateToken(TokenResult, c)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{
			"statusError": true,
			"message":     err.Error(),
		})
		return
	}
	var transaction entities.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	transaction.Transaction_id = nanoid.New()
	transaction.User_id = claims.ID
	transaction.Created_at = time.Now().String()
	transaction.Updated_at = time.Now().String()
	err = h.Service.AddTransaction(transaction)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
}
