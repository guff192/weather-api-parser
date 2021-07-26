package main

import (
	"weather-api-parser/web"
)

func main() {
	StaticServer := web.NewStaticServer()
	go StaticServer.Run()

	WeatherServer := web.NewWeatherServer()
	WeatherServer.Run()
}
