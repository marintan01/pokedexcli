package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("Command inspect needs a single pokemon argument")
	}

	pokemonName := args[0]

	pokemon, present := cfg.pokedex[pokemonName]

	if !present {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println(pokemon)
	return nil

}
