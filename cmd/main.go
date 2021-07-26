package main

import (
	"fmt"
	"weather-api-parser/web"
)

func main() {
	client := web.NewClient()
	resp, err := client.Get("Saint Petersburg", "f5adafaa4825669c205ae3547c92e61a")
	if err != nil {
		panic("error with WeatherAPI: " + err.Error())
	}
	fmt.Printf("%s", resp)
}
