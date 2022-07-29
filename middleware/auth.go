package middleware

import (
	"fmt"
	"net/http"
	"project/service"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				panic(err)
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			panic(err)
		}
		clientToken := cookie.Value
		claims := service.ValidateToken(clientToken)
		fmt.Print("------------------------------->>", claims.UserID)

		ctx.Set("UserID", claims.UserID)

		ctx.Next()
	}
}
