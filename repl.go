package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

func cleanInput(text string) []string {
	cleaned := []string{}

	words := strings.Fields(text)
	for _, word := range words {
		cleaned = append(cleaned, strings.ToLower(word))
	}

	return cleaned
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)

			fmt.Printf("Your command was: %v\n", cleanedInput[0])
		} else {
			break
		}
	}
}
