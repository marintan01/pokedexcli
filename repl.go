package main

import (
	"bufio"
	"fmt"
	"github.com/marintan01/pokedexcli/internal/pokecache"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	cache    *pokecache.Cache
	next     *string
	previous *string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := pokecache.NewCache(time.Minute * 5)
	cfg := &config{
		cache: &c,
	}

	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()

		clearedInput := cleanInput(input)

		command, ok := getCommands()[clearedInput[0]]
		args := clearedInput[1:]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args)
		if err != nil {
			fmt.Printf("An error as accoured: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splitted := strings.Fields(text)
	return splitted

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 pokemon locations",
			callback:    commandMapNext,
		},
		"explore": {
			name:        "explore",
			description: "Given the location name it return the list of pokemons",
			callback:    commandExplore,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 pokemon locations",
			callback:    commandMapPrevious,
		},
	}
}
