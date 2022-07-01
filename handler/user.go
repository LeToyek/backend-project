package handler

import (
	"coin-batam/entities"
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
