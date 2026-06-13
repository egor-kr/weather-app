package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/egor-kr/weather-app/internal/http_server"
	"github.com/egor-kr/weather-app/internal/logger"
	weather_forecast_reciever "github.com/egor-kr/weather-app/internal/weather/forecast_reciever"
	weather_service "github.com/egor-kr/weather-app/internal/weather/service"
	weather_transport_http "github.com/egor-kr/weather-app/internal/weather/transport/http"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cancel()
	loggerCfg := logger.NewConfigMust()
	log, err := logger.New(loggerCfg)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to initialize logger: %w", err))
		os.Exit(1)
	}
	defer log.Close()

	slog.SetDefault(log.Logger)

	slog.Debug("Initializing gin router")
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard

	eng := gin.New()
	eng.Use(sloggin.New(log.Logger), gin.Recovery())

	receiverCfg := weather_forecast_reciever.NewConfigMust()
	receiver := weather_forecast_reciever.New(receiverCfg)
	weatherService := weather_service.New(receiver)
	weatherTransport := weather_transport_http.New(weatherService)

	slog.Debug("Registering routes")
	routes := weatherTransport.GetRoutes()
	for _, route := range routes {
		eng.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	eng.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	srvCfg := http_server.NewConfigMust()
	srv := http_server.New(srvCfg, eng.Handler(), log.Logger)

	slog.Debug("Starting server")
	if err := srv.Run(ctx); err != nil {
		log.Error("server stopped", slog.String("error", err.Error()))
	}
}
