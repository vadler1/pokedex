package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapb(cfg *config) error {

	var url string

	if cfg.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *cfg.prev
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
