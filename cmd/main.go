package main

import "weather-observer/internal/app"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
