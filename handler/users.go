package handler

import (
	"project/entities"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user entities.User
	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": user,
	})
}
