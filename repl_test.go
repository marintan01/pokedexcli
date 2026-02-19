package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "miao WOIF aslfhjas",
			expected: []string{"miao", "woif", "aslfhjas"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for _, acc := range actual {
			print("INIZIO " + acc)
		}
		if len(actual) != len(c.expected) {
			print("Mi trovo al" + actual[0] + c.expected[0])
			t.Errorf("ERROR - numero dei risultati non coerenti \nactual: %d \nexpected %d", len(actual), len(c.expected))
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("ERROR - non matching word \nactual: %s \nexpected %s", word, expectedWord)
			}
		}
	}

}
