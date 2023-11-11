package main

import (
	"clientmetrics/pkg/clientmetrics"
	"clientmetrics/pkg/server"
	"log/slog"
	"os"

	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

func main() {
	viper.AutomaticEnv()
	if err := clientmetrics.Init(); err != nil {
		slog.Error("could not initialize client metrics", "err", err.Error())
		os.Exit(1)
	}
	server.Run()
}
