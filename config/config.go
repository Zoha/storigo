package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// path to .env
const envFilePath = ".env"

// config keys
const (
	// general configs
	APP_NAME string = "APP_NAME"

	// server configs
	APP_HOST_AND_PORT string = "APP_HOST_AND_PORT"
	GIN_MODE          string = "GIN_MODE"
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
