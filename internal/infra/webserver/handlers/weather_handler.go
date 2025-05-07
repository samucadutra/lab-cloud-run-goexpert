package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"regexp"
)

type WeatherData struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
}

type WeatherHandler struct {
	WeatherApiKey string
}

func NewWeatherHandler(apiKey string) *WeatherHandler {
	return &WeatherHandler{WeatherApiKey: apiKey}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	zipcode := chi.URLParam(r, "zipcode")

	if !isValidZipCode(zipcode) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := fetchLocation(zipcode)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := fetchWeather(location, h.WeatherApiKey)
	if err != nil {
		http.Error(w, "failed to fetch weather data", http.StatusInternalServerError)
		return
	}

	tempK := weather.TempC + 273.15

	response := map[string]float64{
		"temp_C": weather.TempC,
		"temp_F": weather.TempF,
		"temp_K": tempK,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
		return "", fmt.Errorf("location not found")
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
