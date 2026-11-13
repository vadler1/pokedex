package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	MIN_BASE_EXPERIENCE = 36
	MAX_BASE_EXPERIENCE = 608
)

func catchProbability(baseExperience int) float64 {
	adjustedExperience := baseExperience - MIN_BASE_EXPERIENCE

	experienceRatio := float64(adjustedExperience) / float64(MAX_BASE_EXPERIENCE-MIN_BASE_EXPERIENCE)
	probability := 1.0 - experienceRatio
	if probability < 0 {
		probability = 0
	}
	return probability
}

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

	baseExperience := pokemonData.BaseExperience

	probability := catchProbability(baseExperience)
	if rand.Float64() < probability {
		fmt.Println(pokemonName + " was caught!")
		cfg.caughtPokemon[pokemonName] = pokemonData
	} else {
		fmt.Println(pokemonName + " escaped!")
	}

	return nil
}
