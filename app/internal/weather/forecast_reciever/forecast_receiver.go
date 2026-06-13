package weather_forecast_reciever

import (
	"net/http"
)

type ForecastReceiver struct {
	c       *http.Client
	baseURL string
}

func New(cfg Config) ForecastReceiver {
	return ForecastReceiver{
		c: &http.Client{
			Timeout: cfg.Timeout,
		},
		baseURL: cfg.URL,
	}
}
