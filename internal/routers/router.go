package routers

import (
	"github.com/gin-gonic/gin"
	c "go-ecommerce-be/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", c.NewUserController().GetUserById)
	}
	return r
}
