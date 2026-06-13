package weather_transport_http

import "github.com/egor-kr/weather-app/internal/domain"

type GetForecastRequest struct {
	City string `json:"city"`
}

type GetForecastResponse struct {
	Status        string  `json:"status"`
	Temperature   float64 `json:"temperature"`
	FeelsLike     float64 `json:"feels_like"`
	WindSpeed     float64 `json:"wind_speed"`
	WindDirection string  `json:"wind_direction"`
	Pressure      int     `json:"pressure"`
	RainChance    int     `json:"rain_chance"`
}

func domainToResponse(forecast domain.Forecast) GetForecastResponse {
	return GetForecastResponse{
		Status:        forecast.Status,
		Temperature:   forecast.Temperature,
		FeelsLike:     forecast.FeelsLike,
		WindSpeed:     forecast.WindSpeed,
		WindDirection: forecast.WindDirection,
		Pressure:      forecast.Pressure,
		RainChance:    forecast.RainChance,
	}
}
