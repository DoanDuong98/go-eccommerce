package initialize

import (
	"fmt"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	err := r.Run(":8002")
	if err != nil {
		fmt.Print("has err")
		return
	}
}
