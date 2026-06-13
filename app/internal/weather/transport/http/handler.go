package weather_transport_http

import (
	"context"

	"github.com/egor-kr/weather-app/internal/domain"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service WeatherService
}

func New(service WeatherService) HTTPHandler {
	return HTTPHandler{
		service: service,
	}
}

type WeatherService interface {
	GetForecast(ctx context.Context, city string) (domain.Forecast, error)
}

func (h HTTPHandler) GetRoutes() []gin.RouteInfo {
	return []gin.RouteInfo{
		{
			Method:      "GET",
			Path:        "/forecast/:city",
			HandlerFunc: h.GetForecast,
		},
	}
}
