package handler

import (
	"net/http"
	"project/entities"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entities.User

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Password = HashPassword(user.Password)

		timeNow, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Created_at = timeNow.String()
		user.Updated_at = timeNow.String()
		user.User_id = "uuid.New"

		c.JSON(200, user)
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
