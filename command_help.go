package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, commandStruct := range getCommands() {
		fmt.Printf("%s: %s\n", commandStruct.name, commandStruct.description)
	}

	return nil
}
