package handler

import (
	"coin-batam/entities"
	"net/http"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddUser(c *gin.Context) {
	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	user.User_id = nanoid.New()
	user.Created_at = time.Now().String()
	user.Updated_at = time.Now().String()

	h.Service.AddUser(user)
}
func (h *Handler) Login(c *gin.Context) {
	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}
	userId, err := h.Service.UseUser(user.Email, user.Password)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"statusError": true,
			"message":     err.Error(),
		})
		return
	}
	token, expirationTime := GenerateAllTokens(userId)
	c.IndentedJSON(http.StatusOK, gin.H{
		"statusError": false,
		"token":       token,
		"expired":     expirationTime,
	})
}
func (h *Handler) GetUserById(c *gin.Context) {
	claims, err := ValidateToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{
			"statusError": true,
			"message":     err.Error(),
		})
		return
	}
	user, err := h.Service.GetUserById(claims.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{
			"statusError": true,
			"message":     err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}
