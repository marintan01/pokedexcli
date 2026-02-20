package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()

		result := cleanInput(input)[0]

		fmt.Printf("Your command was: %s\n", result)

	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splitted := strings.Fields(text)
	return splitted

}
