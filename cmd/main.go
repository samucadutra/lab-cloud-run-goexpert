package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samucadutra/lab-cloud-run-goexpert/configs"
	"github.com/samucadutra/lab-cloud-run-goexpert/internal/infra/webserver"
	"github.com/samucadutra/lab-cloud-run-goexpert/internal/infra/webserver/handlers"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	webserver := webserver.NewWebServer(config.WebServerPort)
	webserver.Router.Use(middleware.Logger)
	webserver.Router.Route("/weather", func(r chi.Router) {
		r.Get("/current/{zipcode}", handlers.NewWeatherHandler(config.WeatherApiKey).GetWeather)
	})

	webserver.Start()
}
