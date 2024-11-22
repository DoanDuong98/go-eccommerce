package global

import (
	"go-ecommerce-be/pkg/logger"
	"go-ecommerce-be/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
