package main

import (
	"fmt"
	"github.com/marintan01/pokedexcli/internal/pokeapi"
)

func commandMapNext(cfg *config) error {

	locations, err := pokeapi.GetLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil

}

func commandMapPrevious(cfg *config) error {

	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := pokeapi.GetLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil

}
