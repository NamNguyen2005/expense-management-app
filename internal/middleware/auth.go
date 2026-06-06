package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	// Authentication logic would go here
	return func(ctx *gin.Context) {
		// For now, just pass through
		ctx.Next()
	}
}