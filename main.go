package main

import (
	"time"

	"github.com/vadler1/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
