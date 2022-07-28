package middleware

import (
	"net/http"
	"project/service"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "token is missing"})
			ctx.Abort()
			return
		}

		claims := service.ValidateToken(clientToken)

		ctx.Set("UserID", claims.UserID)

		ctx.Next()
	}
}
