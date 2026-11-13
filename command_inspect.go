package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	if _, exists := cfg.caughtPokemon[pokemonName]; !exists {
		return errors.New("you have not caught that pokemon")
	}

	pokemonData := cfg.caughtPokemon[pokemonName]

	fmt.Println("Name: " + pokemonData.Name)
	fmt.Println("Height: ", pokemonData.Height)
	fmt.Println("Weight: ", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Println("  -"+stat.Stat.Name+": ", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonData.Types {
		fmt.Println("  - " + t.Type.Name)
	}

	return nil
}
