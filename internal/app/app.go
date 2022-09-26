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

	//httpHandlers := http.NewHandler(tokenManager)
	//httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTP))
	//
	//go func() {
	//	if err := httpServer.Run(); err != nil {
	//		logrus.Errorf("error occurred while running http server: %s\n", err.Error())
	//	}
	//}()
}
