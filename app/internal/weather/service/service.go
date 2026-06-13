package weather_service

import (
	"context"

	"github.com/egor-kr/weather-app/internal/domain"
)

type WeatherService struct {
	receiver ForecastReceiver
}

func New(receiver ForecastReceiver) WeatherService {
	return WeatherService{
		receiver: receiver,
	}
}

type ForecastReceiver interface {
	GetForecast(ctx context.Context, city string) (domain.Forecast, error)
}
