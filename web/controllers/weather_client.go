package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherClient interface {
	Get(city, key string) error
	Parse() (weatherAPIResponse, error)
}

func NewClient() WeatherClient {
	return &weatherClient{
		client:   http.DefaultClient,
		response: []byte{},
	}
}

type weatherClient struct {
	client   *http.Client
	response []byte
}

func (wc *weatherClient) Get(city, key string) error {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&lang=RU&units=metric",
		city, key)
	resp, err := wc.client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bytesResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	wc.response = bytesResp
	return nil
}

type weatherStr struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type mainWeather struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
}

type weatherAPIResponse struct {
	Weather []weatherStr `json:"weather"`
	Main    mainWeather  `json:"main"`
	Name    string       `json:"name"`
}

func (wc *weatherClient) Parse() (weatherAPIResponse, error) {
	weather := weatherAPIResponse{}

	if err := json.Unmarshal(wc.response, &weather); err != nil {
		return weatherAPIResponse{}, err
	}
	return weather, nil
}
