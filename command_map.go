package main

import "fmt"

func callbackMap(cfg *Config, params ...string) error {
	locations, err := cfg.pokeapiClient.GetLocationAreasResponse(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, loc := range locations.Results {
		fmt.Printf("- %s\n", loc.Name)
	}

	cfg.nextLocationUrl = locations.Next
	cfg.prevLocationUrl = locations.Previous

	return nil
}