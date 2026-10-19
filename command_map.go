package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	cfg.pageLocation += 1
	offset := (cfg.pageLocation - 1) * 20
	fmt.Printf("You are on page: %d\n", cfg.pageLocation)
	for i, loc := range locationsResp.Results {
		fmt.Printf("%d - %s\n", i+1+(offset), loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	cfg.pageLocation -= 1
	offset := (cfg.pageLocation - 1) * 20
	fmt.Printf("You are on page: %d\n", cfg.pageLocation)
	for i, loc := range locationResp.Results {
		fmt.Printf("%d - %s\n", i+1+(offset), loc.Name)
	}
	return nil
}
