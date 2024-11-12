package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-be/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailResponse(c, response.ErrCodeTokenInvalid, "")
			c.Abort()
			return
		}
		c.Next()
	}
}
