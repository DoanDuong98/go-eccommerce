package global

import (
	"go-ecommerce-be/pkg/logger"
	"go-ecommerce-be/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
