package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
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

		commandStruct, exists := getCommands()[command]
		if !exists {
			fmt.Println("Unknown Command")
			continue
		}

		err := commandStruct.callback()
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
	callback    func() error
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
	}
}
