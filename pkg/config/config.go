package config

import (
	"sync"

	"github.com/spf13/viper"
)

// Config is the struct that contains all the configuration variables
type Config struct {
	Port           string
	AllowedOrigins []string
}

var once sync.Once
var singleton *Config

func setViperDefaults() {
	viper.SetDefault("PORT", "9091")
	viper.SetDefault("ALLOWED_ORIGINS", []string{"*"})
}

func GetConfig() *Config {
	once.Do(
		func() {
			setViperDefaults()
			singleton = &Config{
				Port:           viper.GetString("PORT"),
				AllowedOrigins: viper.GetStringSlice("ALLOWED_ORIGINS"),
			}
		})

	return singleton
}
