package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any Pokemon yet.")
		return nil
	}

	for name, _ := range cfg.caughtPokemon {
		fmt.Println(" - " + name)
	}

	return nil
}
