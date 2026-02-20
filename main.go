package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()

		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback()
	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splitted := strings.Fields(text)
	return splitted

}
