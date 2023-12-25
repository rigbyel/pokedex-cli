package pokeapi

import (
	"net/http"
	"time"
)

type HttpClient struct {
	httpClient http.Client
}

func NewHttpClient() HttpClient {
	return HttpClient{
		httpClient: http.Client { 
			Timeout: 5*time.Second * time.Duration(1),
		},
	}
}

var baseUrl = "https://pokeapi.co/api/v2"

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
