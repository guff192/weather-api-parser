package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherClient interface {
	Get(city, key string) ([]byte, error)
}

func NewClient() WeatherClient {
	return weatherClient{
		client: &http.Client{},
	}
}

type weatherClient struct {
	client *http.Client
}

func (wc weatherClient) Get(city, key string) ([]byte, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, key)
	resp, err := wc.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytesResp, _ := ioutil.ReadAll(resp.Body)
	return bytesResp, nil
}
