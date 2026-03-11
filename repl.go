package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anand-anshul/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()

		cleanUserInput := cleanInput(userInput)

		if len(cleanUserInput) == 0 {
			continue
		}

		command := cleanUserInput[0]
		args := []string{}
		if len(cleanUserInput) > 1 {
			args = cleanUserInput[1:]
		}

		commandStruct, exists := getCommands()[command]
		if !exists {
			fmt.Println("Unknown Command")
			continue
		}

		err := commandStruct.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(trimmedText)
	textSlice := strings.Fields(loweredText)
	return textSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next page",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "exit",
			description: "Display previous page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemon in area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
	}
}
