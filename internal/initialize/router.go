package initialize

import (
	"github.com/gin-gonic/gin"
	c "go-ecommerce-be/internal/controller"
	"go-ecommerce-be/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// Use middleware
	r.Use(middlewares.AuthMiddleware())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewUserController().GetUserById)
	}
	return r
}
