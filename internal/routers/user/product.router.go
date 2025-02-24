package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (p *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/list")
		productRouterPublic.GET("/detail/:id")
	}
	// private router
}