package usecase

// WeatherUseCase defines the interface for getting weather by zipcode
type WeatherUseCase interface {
	Execute(zipcode string) (map[string]float64, error)
}
