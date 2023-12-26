package main

import (
	"fmt"

	"github.com/rigbyel/pokedex-cli/internal/pokeapi"
)

func callbackExplore(cfg *Config, params ...string) error{
	if len(params) == 0 {
		return fmt.Errorf("no location area provided")
	}
	areaName := params[0]
	areaInfo, err := pokeapi.GetExploreAreaResponse(&cfg.pokeapiClient, areaName)
	if err != nil {
		return err
	}

	for _, pokemon := range areaInfo.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}