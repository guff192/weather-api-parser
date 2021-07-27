package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleCurrentWeatherByCityName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city, ok := vars["city"]
	if !ok {
		http.Error(w, "missing city name", http.StatusBadRequest)
		return
	}

	weatherClient := NewClient()

	if err := weatherClient.Get(city, "f5adafaa4825669c205ae3547c92e61a"); err != nil {
		http.Error(w, "error with WeatherAPI: "+err.Error(), http.StatusInternalServerError)
	}

	weather, err := weatherClient.Parse()
	if err != nil {
		http.Error(w, "error parsing weather: "+err.Error(), http.StatusInternalServerError)
	}
	weatherJSON, _ := json.Marshal(weather)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(weatherJSON)
}
