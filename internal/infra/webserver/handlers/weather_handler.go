package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/samucadutra/lab-cloud-run-goexpert/internal/usecase"
	"net/http"
)

type WeatherHandler struct {
	WeatherUseCase usecase.WeatherUseCase
}

func NewWeatherHandler(apiKey string) *WeatherHandler {
	return &WeatherHandler{
		WeatherUseCase: usecase.NewGetWeatherByZipcodeUseCase(apiKey),
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	zipcode := chi.URLParam(r, "zipcode")

	response, err := h.WeatherUseCase.Execute(zipcode)
	if err != nil {
		if err.Error() == "cannot find zipcode" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
