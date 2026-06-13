package weather_forecast_reciever

import "github.com/egor-kr/weather-app/internal/domain"

type ForecastResponse struct {
	Current Current `json:"current"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	Condition    Condition `json:"condition"`
	TempC        float64   `json:"temp_c"`
	FeelsLikeC   float64   `json:"feelslike_c"`
	WindKph      float64   `json:"wind_kph"`
	WindDir      string    `json:"wind_dir"`
	PressureMb   float64   `json:"pressure_mb"`
	ChanceOfRain int       `json:"chance_of_rain"`
}

func (r ForecastResponse) toDomain() domain.Forecast {
	return domain.Forecast{
		Status:        r.Current.Condition.Text,
		Temperature:   r.Current.TempC,
		FeelsLike:     r.Current.FeelsLikeC,
		WindSpeed:     r.Current.WindKph,
		WindDirection: r.Current.WindDir,
		Pressure:      int(r.Current.PressureMb),
		RainChance:    r.Current.ChanceOfRain,
	}
}
