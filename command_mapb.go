package main

import "fmt"

func callbackMapb(cfg *Config, params ...string) error {
	if cfg.prevLocationUrl == nil {
		return fmt.Errorf("unable to move back, you are on the first page")
	}
	
	locations, err := cfg.pokeapiClient.GetLocationAreasResponse(cfg.prevLocationUrl)
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