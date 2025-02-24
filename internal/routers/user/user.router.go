package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	UserRouterPublic := Router.Group("/user")
	{
		UserRouterPublic.POST("/register")
		UserRouterPublic.POST("/otp")
	}
	// private router
	UserRouterPrivate := Router.Group("/user")
	// UserRouterPrivate.Use(middleware.limiter())
	// UserRouterPrivate.Use(middleware.Auth())
	// UserRouterPrivate.Use(middleware.Permission())
	{
		UserRouterPrivate.GET("/get_info")
	}
}