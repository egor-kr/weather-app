package domain

type Forecast struct {
	Status        string
	Temperature   float64
	FeelsLike     float64
	WindSpeed     float64
	WindDirection string
	Pressure      int
	RainChance    int
}
