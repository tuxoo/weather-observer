package app

import (
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/idler/pkg/auth"
	"github.com/tuxoo/idler/pkg/db/mongo"
	"weather-observer/internal/config"
	"weather-observer/internal/controller/http"
	"weather-observer/internal/repository"
	"weather-observer/internal/server"
	"weather-observer/internal/service"
)

func Run(configPath string) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	tokenManager := auth.NewJWTTokenManager(cfg.Auth.JWT.SigningKey)

	mongoClient, err := mongo.NewMongoDb(cfg.Mongo)
	if err != nil {
		logrus.Fatalf("error initializing mongo client: %s", err.Error())
	}

	mongoDb := mongoClient.Database(cfg.Mongo.DB)

	repositories := repository.NewRepositories(mongoDb)

	services := service.NewServices(repositories)

	httpHandlers := http.NewHandler(services.UserService, tokenManager)
	httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTPConfig))

	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()
}
