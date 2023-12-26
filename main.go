package main

import (
	"fmt"
	"time"

	"github.com/rigbyel/pokedex-cli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.HttpClient
	prevLocationUrl *string
	nextLocationUrl *string
}

func main() {
	cfg := Config {
		pokeapiClient : pokeapi.NewHttpClient(5*time.Minute),
		prevLocationUrl: nil,
		nextLocationUrl: nil,
	}
	
	fmt.Println("Hello! This is Pokedex. Enter 'help' command to see more information.")
	startREPL(&cfg)
}