package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemonStruct := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonStruct.Name)
	}

	return nil
}
