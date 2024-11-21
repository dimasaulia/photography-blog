package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var CONFIG *viper.Viper

func ReadConfigigurationFile() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read configuration file")
	}

	CONFIG = config
	return config
}
