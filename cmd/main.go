package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/samucadutra/lab-cloud-run-goexpert/configs"
	"github.com/samucadutra/lab-cloud-run-goexpert/internal/infra/webserver"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	webserver := webserver.NewWebServer(config.WebServerPort)
	webserver.Router.Route("/weather", func(r chi.Router) {

		webserver.Router.Get("/current/{zipcode}", webserver.NewWeatherHandler(config.WeatherApiKey).GetWeather)
	})

}
