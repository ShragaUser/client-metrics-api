package server

import (
	"clientmetrics/pkg/clientmetrics"
	"clientmetrics/pkg/config"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.GetConfig()
	registerGinDefaults()
	r := gin.New()
	registerGinRouterDefaults(r)
	registerRoutes(r)
	err := endless.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r)
	if err != nil {
		slog.Error("shutting down server", "err", err.Error())
	}
}

func isAliveHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func registerRoutes(r *gin.Engine) {
	m := clientmetrics.GetMonitor()
	m.SetMetricPath(metricsRoute)
	m.Use(r) // uses /metrics endpoint
	r.GET(isAliveRoute, isAliveHandler)
	r.POST(addMetricRoute, clientmetrics.PostMetricHandler)
}

func registerGinRouterDefaults(r *gin.Engine) {
	r.Use(
		gin.Recovery(),
		cors.New(corsRouterConfig()),
		gin.LoggerWithWriter(gin.DefaultErrorWriter, isAliveRoute),
	)
}

func registerGinDefaults() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
}
