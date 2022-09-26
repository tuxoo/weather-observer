package app

import (
	"github.com/sirupsen/logrus"
	"weather-observer/internal/config"
)

func Run(configPath string) {
	_, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
}
