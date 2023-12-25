package main

import (
	"fmt"
	"log"
	"github.com/rigbyel/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewHttpClient()
	locations, err := pokeapiClient.GetLocationAreasResponse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(locations)

	//fmt.Println("Hello! This is Pokedex. Enter 'help' command to see more information.")
	//startREPL()
}