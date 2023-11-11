package main

import (
	"clientmetrics/pkg/server"

	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

func main() {
	viper.AutomaticEnv()
	server.Run()
}
