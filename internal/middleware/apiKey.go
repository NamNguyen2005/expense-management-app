package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware() gin.HandlerFunc {
	successKey := os.Getenv("API_KEY")
	if successKey == "" {
		successKey = "test-case"
	}
	return func(ctx *gin.Context){
		currentKey := ctx.GetHeader("X-API-Key")
		if currentKey == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest , gin.H {
				"err" : "missing api key ???",
			})
			return 
		}
		if currentKey != successKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized , gin.H{
				"err" :"Unauthorized" , 
			})
			return 
		}
		ctx.Next()
	}
}