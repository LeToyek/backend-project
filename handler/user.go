package handler

import "github.com/gin-gonic/gin"

func (h *Handler) AddUser(c *gin.Context) {

	h.Service.AddUser()
}
