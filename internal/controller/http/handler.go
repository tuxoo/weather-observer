package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tuxoo/idler/pkg/auth"
	"github.com/tuxoo/weather-observer/docs"
	"github.com/tuxoo/weather-observer/internal/config"
	"github.com/tuxoo/weather-observer/internal/service"
	"net/http"
	"time"
)

type Handler struct {
	userService  service.Users
	tokenManager auth.TokenManager
}

func NewHandler(userService service.Users, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		userService:  userService,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(cfg config.HTTPConfig) *gin.Engine {
	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(corsConfig),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("api/v1/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	//h.initMetrics(router)
	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h.initUserRoutes(api)
	}
}
