package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsStruct, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsStruct.Next
	cfg.prevLocationsURL = locationsStruct.Previous

	locations := locationsStruct.Results

	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsStruct, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsStruct.Next
	cfg.prevLocationsURL = locationsStruct.Previous

	locations := locationsStruct.Results

	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}
