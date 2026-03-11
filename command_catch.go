package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("not valid Pokemon")
	}
	name := args[0]
	pokemonStruct, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	res := rand.Intn(pokemonStruct.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonStruct.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemonStruct.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonStruct.Name)
	cfg.caughtPokemon[pokemonStruct.Name] = pokemonStruct
	return nil

}
