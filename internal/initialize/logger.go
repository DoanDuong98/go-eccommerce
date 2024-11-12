package initialize

import (
	"go-ecommerce-be/global"
	"go-ecommerce-be/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
