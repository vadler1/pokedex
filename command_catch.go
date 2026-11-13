package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	if _, exists := cfg.caughtPokemon[pokemonName]; exists {
		return errors.New(pokemonName + " has already been caught")
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

	pokemonData, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	probability := rand.Intn(pokemonData.BaseExperience)

	if probability > 40 {
		fmt.Println(pokemonName + " escaped!")
		return nil
	}

	fmt.Println(pokemonName + " was caught!")
	cfg.caughtPokemon[pokemonName] = pokemonData

	return nil
}
