package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// path to .env
const envFilePath = ".env"

// config keys
const (
	APP_NAME string = "APP_NAME"
)

// load configs
func ReadConfigs() {
	viper.SetConfigFile(envFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Get(key string) string {
	return viper.GetString(key)
}
