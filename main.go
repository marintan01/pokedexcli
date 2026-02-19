package main

import "strings"

func main() {
	//cleanInput("CIofaAF df sf ")
	a := cleanInput("miao WOIF aslfhjas")

	for _, b := range a {
		print("\n" + b)
	}
	a = cleanInput("  hello  world  ")
	for _, b := range a {
		print("\n" + b)
	}

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splitted := strings.Fields(text)
	return splitted

}
