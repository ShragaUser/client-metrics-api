package config

import (
	"log/slog"
	"sync"

	"github.com/spf13/viper"
)

// Config is the struct that contains all the configuration variables
type Config struct {
	Port                        string
	AllowedOrigins              []string
	ManualConfigurationFilePath string
	logLevel                    string
}

var once sync.Once
var singleton *Config

func setViperDefaults() {
	viper.SetDefault("PORT", "9091")
	viper.SetDefault("ALLOWED_ORIGINS", []string{"*"})
	viper.SetDefault("CONFIGURATION_FILE_PATH", "")
	viper.SetDefault("LOG_LEVEL", "INFO")
}

func GetConfig() *Config {
	once.Do(
		func() {
			setViperDefaults()
			singleton = &Config{
				Port:                        viper.GetString("PORT"),
				AllowedOrigins:              viper.GetStringSlice("ALLOWED_ORIGINS"),
				ManualConfigurationFilePath: viper.GetString("CONFIGURATION_FILE_PATH"),
				logLevel:                    viper.GetString("LOG_LEVEL"),
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

func (c *Config) GetLogLevel() slog.Level {
	levelvar := slog.LevelVar{}
	if err := levelvar.UnmarshalText([]byte(c.logLevel)); err != nil {
		slog.Error("could not parse log level", "err", err.Error())
		return slog.LevelInfo
	}

	return levelvar.Level()
}
