package server

import (
	"coin-batam/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Handler *handler.Handler
}

func (s *Server) StartServer() {
	s.Router.POST("/test", s.Handler.JustTest)
}
