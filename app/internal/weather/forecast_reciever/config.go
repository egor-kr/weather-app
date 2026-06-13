package weather_forecast_reciever

import (
	"fmt"
	"time"

	"github.com/egor-kr/weather-app/internal/utils"
)

type Config struct {
	URL     string        `envconfig:"URL" required:"true"`
	Timeout time.Duration `envconfig:"TIMEOUT" default:"5s"`
}

func NewConfigMust() Config {
	cfg, err := utils.NewEnvConfig[Config]("FORECAST_RECEIVER")

	if err != nil {
		err = fmt.Errorf("get forecast receiver config: %w", err)
		panic(err)
	}

	return cfg
}
