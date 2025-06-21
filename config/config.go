package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
}

type Server struct {
	Host string
	Port int
}

func GetConfig() Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	return Config{
		Server: Server{
			Host: viper.GetString("SRV_HOST"),
			Port: viper.GetInt("SRV_PORT"),
		},
	}
}
