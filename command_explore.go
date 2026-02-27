package main

import (
	"fmt"
	"github.com/marintan01/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("Command explore needs a single location argument")
	}

	location, err := pokeapi.GetPokemonInLocation(cfg.cache, args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *location.Name)
	for _, pokemon := range location.PokemonEncaunters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil

}
