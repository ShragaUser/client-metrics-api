package config

import (
	"sync"

	"github.com/spf13/viper"
)

// Config is the struct that contains all the configuration variables
type Config struct {
	Port                        string
	AllowedOrigins              []string
	ManualConfigurationFilePath string
}

var once sync.Once
var singleton *Config

func setViperDefaults() {
	viper.SetDefault("PORT", "9091")
	viper.SetDefault("ALLOWED_ORIGINS", []string{"*"})
	viper.SetDefault("CONFIGURATION_FILE_PATH", "")
}

func GetConfig() *Config {
	once.Do(
		func() {
			setViperDefaults()
			singleton = &Config{
				Port:                        viper.GetString("PORT"),
				AllowedOrigins:              viper.GetStringSlice("ALLOWED_ORIGINS"),
				ManualConfigurationFilePath: viper.GetString("CONFIGURATION_FILE_PATH"),
			}
		})

	return singleton
}

func (c *Config) ConfigFileSupported() bool {
	return c.ManualConfigurationFilePath != ""
}

func (c *Config) GetPreDefinedCustomMetricsConfig() (*InputFile, error) {
	return ParseConfigFileFromFilePath(c.ManualConfigurationFilePath)
}
