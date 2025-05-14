package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type WeatherData struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
}

type GetWeatherByZipcodeUseCase struct {
	WeatherApiKey string
}

func NewGetWeatherByZipcodeUseCase(apiKey string) *GetWeatherByZipcodeUseCase {
	return &GetWeatherByZipcodeUseCase{WeatherApiKey: apiKey}
}

func (uc *GetWeatherByZipcodeUseCase) Execute(zipcode string) (map[string]float64, error) {
	if !isValidZipCode(zipcode) {
		return nil, fmt.Errorf("invalid zipcode")
	}

	location, err := fetchLocation(zipcode)
	if err != nil {
		return nil, fmt.Errorf("cannot find zipcode: %w", err)
	}

	weather, err := fetchWeather(location, uc.WeatherApiKey)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	tempK := weather.TempC + 273.15

	return map[string]float64{
		"temp_C": weather.TempC,
		"temp_F": weather.TempF,
		"temp_K": tempK,
	}, nil
}

func isValidZipCode(zipcode string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(zipcode)
}

func fetchLocation(zipcode string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch location")
	}

	var data struct {
		Localidade string `json:"localidade"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", fmt.Errorf("can not find zipcode")
	}

	return data.Localidade, nil
}

func fetchWeather(location, apiKey string) (*WeatherData, error) {
	baseURL := "http://api.weatherapi.com/v1/current.json"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", apiKey)
	q.Add("q", location)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data")
	}

	var data struct {
		Current WeatherData `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data.Current, nil
}
