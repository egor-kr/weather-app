package weather_service

import (
	"context"

	"github.com/egor-kr/weather-app/internal/domain"
)

func (s WeatherService) GetForecast(ctx context.Context, city string) (domain.Forecast, error) {
	return s.receiver.GetForecast(ctx, city)
}
