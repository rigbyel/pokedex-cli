package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetExploreAreaResponse (client *HttpClient, areaName string) (ExploreAreaResponse, error) {
	totalUrl := baseUrl + "/location-area/" + areaName

	if data, ok := client.cache.Get(totalUrl); ok {
		areaResponse := ExploreAreaResponse{}
		err := json.Unmarshal(data, &areaResponse)
		return areaResponse, err
	}

	req, err := http.NewRequest("GET", totalUrl, nil)
	if err != nil {
		return ExploreAreaResponse{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return ExploreAreaResponse{}, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploreAreaResponse{}, err
	}

	areaResponse := ExploreAreaResponse{}
	err = json.Unmarshal(data, &areaResponse)
	if err != nil {
		return ExploreAreaResponse{}, err
	}

	client.cache.Add(totalUrl, data)
	return areaResponse, nil
}