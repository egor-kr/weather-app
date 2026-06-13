package weather_forecast_reciever

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/egor-kr/weather-app/internal/domain"
)

func (r ForecastReceiver) GetForecast(
	ctx context.Context,
	city string,
) (domain.Forecast, error) {
	url := fmt.Sprintf("%s&q=%s", r.baseURL, city)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.Forecast{}, err
	}
	fmt.Printf("URL: %s", url)
	resp, err := r.c.Do(req)
	if err != nil {
		return domain.Forecast{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Forecast{}, fmt.Errorf("external api error: status %d", resp.StatusCode)
	}

	var result ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return domain.Forecast{}, err
	}

	fmt.Printf("%v\n", result)

	domainForecast := result.toDomain()

	return domainForecast, nil
}
