package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		Username string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Hostname string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
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

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server Port:", config.Server.Port)
	for _, database := range config.Databases {
		fmt.Println("Database:", database.Username, database.Password, database.Hostname)
	}
}
