package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no location specified")
	}

	locationName := args[0]
	fmt.Println("Exploring " + locationName + "...")

	locationData, err := cfg.pokeapiClient.Explore(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationData.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}
