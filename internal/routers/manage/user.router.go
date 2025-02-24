package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// private router
	UserRouterPrivate := Router.Group("/admin/user")
	// UserRouterPrivate.Use(middleware.limiter())
	// UserRouterPrivate.Use(middleware.Auth())
	// UserRouterPrivate.Use(middleware.Permission())
	{
		UserRouterPrivate.GET("/active_user")
	}
}