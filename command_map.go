package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ptr(s string) *string {
	return &s
}

type locationEntry struct {
	Name string
	URL  string
}

type locationResults struct {
	Count    int
	Next     *string
	Previous *string
	Results  []locationEntry
}

func commandMap(cfg *config) error {

	var url string
	if cfg.next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var results locationResults
	if err := decoder.Decode(&results); err != nil {
		return err
	}

	cfg.next = results.Next
	cfg.prev = results.Previous

	for _, entry := range results.Results {
		fmt.Println(entry.Name)
	}

	return nil
}
