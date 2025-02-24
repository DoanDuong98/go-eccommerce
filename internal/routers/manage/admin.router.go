package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (p *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	AdminRouterPublic := Router.Group("/admin")
	{
		AdminRouterPublic.POST("/login")
	}
	// private router
	AdminRouterPrivate := Router.Group("/admin/user")
	// AdminRouterPrivate.Use(middleware.limiter())
	// AdminRouterPrivate.Use(middleware.Auth())
	// AdminRouterPrivate.Use(middleware.Permission())
	{
		AdminRouterPrivate.GET("/active_user")
	}
}