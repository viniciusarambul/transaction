package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Environments struct {
	APIPort        string
	DatabaseSchema string
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("API_PORT", "8080")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("error to find or read configuration file: %s", err))
	}

	return &Environments{
		APIPort: viper.GetString("API_PORT"),
	}
}
