package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Environments struct {
	APIPort          string
	DatabaseSchema   string
	DatabaseName     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	LogLevel         string
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("API_PORT", "8080")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("error to find or read configuration file: %s", err))
	}

	return &Environments{
		APIPort:          viper.GetString("API_PORT"),
		DatabaseSchema:   viper.GetString("DATABASE_SCHEMA"),
		DatabaseName:     viper.GetString("DATABASE_NAME"),
		DatabaseUsername: viper.GetString("DATABASE_USERNAME"),
		DatabasePassword: viper.GetString("DATABASE_PASSWORD"),
		DatabaseHost:     viper.GetString("DATABASE_HOST"),
		DatabasePort:     viper.GetString("DATABASE_PORT"),
		LogLevel:         viper.GetString("LOG_LEVEL"),
	}
}
