package main

import (
	"fmt"
	"github.com/rigbyel/pokedex-cli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.HttpClient
	prevLocationUrl *string
	nextLocationUrl *string
}

func main() {
	cfg := Config {
		pokeapiClient : pokeapi.NewHttpClient(),
		prevLocationUrl: nil,
		nextLocationUrl: nil,
	}
	
	fmt.Println("Hello! This is Pokedex. Enter 'help' command to see more information.")
	startREPL(&cfg)
}