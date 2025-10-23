package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherResponce struct {
	Current struct {
		Time        string  `json:"time"`
		Temperature float64 `json:"temperature_2m"`
	}
}

type OMClient struct {
	httpClient *http.Client
}

func NewOMClient(httpClient *http.Client) *OMClient {
	return &OMClient{
		httpClient: httpClient,
	}
}

func (o *OMClient) GetTemperature(lat, long float64) (WeatherResponce, error) {
	res, err := o.httpClient.Get(
		fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m",
			lat,
			long,
		),
	)

	if err != nil {
		fmt.Println("err:", err)
		return WeatherResponce{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return WeatherResponce{}, fmt.Errorf("status code %d", res.StatusCode)
	}

	var result WeatherResponce

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("err:", err)
		return WeatherResponce{}, err
	}

	return result, nil
}
