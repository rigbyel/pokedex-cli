package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *HttpClient) GetLocationAreasResponse(pageUrl *string) (LocationAreasResponse, error) {
	var totalUrl string
	if pageUrl == nil {
		endpoint := "/location-area?offset=0&limit=20"
		totalUrl = baseUrl + endpoint
	} else {
		totalUrl = *pageUrl
	} 

	if data, ok := client.cache.Get(totalUrl); ok {
		locationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResp)
		return locationAreasResp, err
	}

	req, err := http.NewRequest("GET", totalUrl, nil)
	if err != nil {
		fmt.Println(err)
		return LocationAreasResponse{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return LocationAreasResponse{}, err
	}
	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	client.cache.Add(totalUrl, data)
	return locationAreasResp, err
}