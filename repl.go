package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vadler1/pokedex/internal/pokeapi"
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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := cfg.commands[command]
		if exists {
			err := cmd.callback(cfg, args...)
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
	callback    func(*config, ...string) error
}

type commandRegistry map[string]cliCommand

type config struct {
	commands         commandRegistry
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
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
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous map page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
	}
}
