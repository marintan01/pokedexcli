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

	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()

		command, ok := getCommands()[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback()
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
	}
}
