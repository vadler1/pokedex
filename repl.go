package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]

		cmd, exists := cfg.commands[command]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	cleaned := []string{}

	words := strings.Fields(text)
	for _, word := range words {
		cleaned = append(cleaned, strings.ToLower(word))
	}

	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type commandRegistry map[string]cliCommand

type config struct {
	commands commandRegistry
	next     *string
	prev     *string
}

func getCommands() commandRegistry {

	return commandRegistry{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous map page",
			callback:    commandMapb,
		},
	}
}
