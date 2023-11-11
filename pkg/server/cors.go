package server

import (
	"clientmetrics/pkg/config"

	"github.com/gin-contrib/cors"
)

// corsRouterConfig configures cors policy for cors.New gin middleware.
func corsRouterConfig() cors.Config {
	cfg := config.GetConfig()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowWildcard = true
	corsConfig.AllowOrigins = cfg.AllowedOrigins

	return corsConfig
}
