package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Responce struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetCoordinates(city string) (Responce, error) {
	r, err := c.httpClient.Get(
		fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json",
			city,
		),
	)
	if err != nil {
		log.Println("err:", err)
		return Responce{}, err
	}

	if r.StatusCode != http.StatusOK {
		return Responce{}, fmt.Errorf("status code %d", r.StatusCode)
	}

	// этот чисто чтобы распарсить нормально потому что там джисон плохо приходит
	var geoResponse struct {
		Result []Responce `json:"results"`
	}

	if err := json.NewDecoder(r.Body).Decode(&geoResponse); err != nil {
		log.Println("err:", err)
		return Responce{}, err
	}

	return geoResponse.Result[0], nil
}
