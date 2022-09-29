package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/idler/pkg/auth"
	"github.com/tuxoo/idler/pkg/db/mongo"
	"github.com/tuxoo/idler/pkg/hash"
	"github.com/tuxoo/weather-observer/internal/config"
	"github.com/tuxoo/weather-observer/internal/controller/http"
	"github.com/tuxoo/weather-observer/internal/repository"
	"github.com/tuxoo/weather-observer/internal/server"
	"github.com/tuxoo/weather-observer/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	tokenManager := auth.NewJWTTokenManager(cfg.Auth.SigningKey)

	mongoClient, err := mongo.NewMongoDb(cfg.Mongo)
	if err != nil {
		logrus.Fatalf("error initializing mongo client: %s", err.Error())
	}

	mongoDb := mongoClient.Database(cfg.Mongo.DB)

	repositories := repository.NewRepositories(mongoDb)

	services := service.NewServices(repositories, hasher, tokenManager, cfg)

	httpHandlers := http.NewHandler(services.UserService, tokenManager)
	httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTP))

	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Print("Weather observer application has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Weather observer application shutting down")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on http server shutting down: %s", err.Error())
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logrus.Errorf("error occured on mongo connection close: %s", err.Error())
	}
}
