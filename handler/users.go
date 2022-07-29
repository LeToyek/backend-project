package handler

import (
	"fmt"
	"net/http"
	"project/entities"
	"project/service"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users []entities.User

func init() {
	users = append(users, entities.User{
		User_id:    "012809uje0-12dj0-21",
		First_name: "rijal",
		Last_name:  "christoph",
		Email:      "rijal@gmail.com",
		Password:   "matamu123",
		Phone:      "192038012938",
		Created_at: time.Now().Local().String(),
		Updated_at: time.Now().Local().String()})
	users = append(users, entities.User{
		User_id:    "012809uje0-12dj0-21",
		First_name: "michael",
		Last_name:  "mulyono",
		Email:      "mulyono@gmail.com",
		Password:   "$2a$10$slhx1XeUybh.dCLBtM3MYOd5BsgiG7pgPBOvBVgAD3SHLfRid7I5C",
		Phone:      "112312341235",
		Created_at: time.Now().Local().String(),
		Updated_at: time.Now().Local().String()})
}

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

		users = append(users, user)
		c.JSON(200, user)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user entities.User
		var foundUser entities.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, u := range users {
			if u.Email == user.Email {
				foundUser = u
				break
			}
			fmt.Print(i)
		}
		passwordIsCorrect, msg := verifyPassword(user.Password, foundUser.Password)
		if passwordIsCorrect == false {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, _, err := service.GenerateTokens(user.User_id)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:  "token",
			Value: token,
		})
		if err != nil {
			panic(err)
		}

	}
}
func verifyPassword(userPass string, foundUserPass string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(foundUserPass), []byte(userPass))
	check := true
	var msg string
	if err != nil {
		check = false
		msg = err.Error()
		return check, msg
	}
	return check, msg
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
