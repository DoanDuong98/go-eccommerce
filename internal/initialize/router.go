package initialize

import (
	"go-ecommerce-be/global"
	"go-ecommerce-be/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Setting.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	// r.Use(gin.Logger()) // logger middleware
	// r.Use(middleware.Cors()) // cors middleware
	// r.Use(gin.Recovery()) // recovery middleware
	manageRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/ping")
	}
	{
		userRouter.UserRouter.InitUserRouter(MainGroup)
		userRouter.ProductRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.UserRouter.InitUserRouter(MainGroup)
		manageRouter.AdminRouter.InitAdminRouter(MainGroup)
	}
	return r
}
