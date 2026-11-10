package main

import "strings"

func cleanInput(text string) []string {
	cleaned := []string{}

	words := strings.Fields(text)
	for _, word := range words {
		cleaned = append(cleaned, strings.ToLower(word))
	}

	return cleaned
}
