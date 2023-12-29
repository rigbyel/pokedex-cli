package main

import "fmt"

func callbackInspect(cfg *Config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("no pokemon name providen")
	}
	pokemonName := params[0]
	pokemonInfo, ok := cfg.caughtPokemons[pokemonName]
	if !ok {
		return fmt.Errorf("%s hasn't been caught yet", pokemonName)
	}
	
	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Printf("Stats:\n")
	for _, item := range pokemonInfo.Stats {
		fmt.Printf("\t-%s: %d\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, item := range pokemonInfo.Types {
		fmt.Printf("\t-%s\n", item.Type.Name)
	}

	return nil
}