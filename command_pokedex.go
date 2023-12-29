package main

import "fmt"

func callbackPokedex(cfg *Config, params ...string) error {
	if len(cfg.caughtPokemons) == 0 {
		fmt.Println("You haven't caught any pokemons yet")
		return nil
	}
	
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf("-%s\n", pokemon.Name)
	}
	return nil
}