package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *Config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("no pokemon name provided")
	}
	pokemonName := params[0]
	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	const maxBaseExperience = 650
	randomNum := rand.Intn(maxBaseExperience)

	// the lower pokemon's base experience, the higher chances to catch it
	if randomNum > pokemonInfo.BaseExperience {
		fmt.Printf("Congrats! You've caught %s!\n", pokemonName)
		cfg.caughtPokemons[pokemonName] = pokemonInfo
	} else {
		fmt.Printf("Oops... %s escaped\n", pokemonName)
	}

	return nil
}