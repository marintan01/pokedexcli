package main

import (
	"fmt"
	"math/rand"

	"github.com/marintan01/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("Command catch needs a single pokemon argument")
	}

	pokemon, err := pokeapi.GetPokemon(cfg.cache, args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemon.Name)
	chance := rand.Intn(pokemon.BaseExperience)
	if chance > 50 {
		fmt.Printf("%s escaped!\n", *pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", *pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[*pokemon.Name] = pokemon
	}

	return nil

}
