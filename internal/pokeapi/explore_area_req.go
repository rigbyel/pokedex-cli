package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type ExploreAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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