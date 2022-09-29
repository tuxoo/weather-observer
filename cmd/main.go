package main

import "github.com/tuxoo/weather-observer/internal/app"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
