package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client HttpClient) GetLocationAreasResponse() (LocationAreasResponse, error) {
	endpoint := "/location-area?offset=0&limit=20"
	totalUrl := baseUrl + endpoint

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
	
	return locationAreasResp, err
}