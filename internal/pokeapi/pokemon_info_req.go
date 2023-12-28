package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (client *HttpClient) GetPokemonInfo(name string) (Pokemon, error) {
	totalUrl := baseUrl + "/pokemon/" + name

	if data, ok := client.cache.Get(totalUrl); ok {
		pokemonInfoResponse := Pokemon{}
		err := json.Unmarshal(data, &pokemonInfoResponse)
		return pokemonInfoResponse, err
	}

	req, err := http.NewRequest("GET", totalUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}
	
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonInfoResponse := Pokemon{}
	err = json.Unmarshal(data, &pokemonInfoResponse)
	if err != nil {
		return Pokemon{}, err
	} 
	
	client.cache.Add(totalUrl, data)
	return pokemonInfoResponse, nil
}