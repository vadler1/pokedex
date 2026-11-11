package main

import (
	"time"

	"github.com/vadler1/pokedex/internal/pokeapi"
	"github.com/vadler1/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		cache:         pokecache.NewCache(5 * time.Minute),
	}

	startRepl(cfg)
}
