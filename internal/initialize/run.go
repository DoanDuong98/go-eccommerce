package initialize

import (
	"fmt"
	"go-ecommerce-be/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Config oke!")
	InitMysql()
	InitRedis()

	r := InitRouter()
	err := r.Run(":8082")
	if err != nil {
		fmt.Print("has err")
		return
	}
}
