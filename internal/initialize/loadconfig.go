package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go-ecommerce-be/global"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	// read config
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("Server Port::", v.GetInt("server.port"))

	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
}
