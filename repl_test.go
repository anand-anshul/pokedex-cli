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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "SINGLE",
			expected: []string{"single"},
		},
		{
			input:    "a  b   c",
			expected: []string{"a", "b", "c"},
		},
		{
			input:    "   leading space",
			expected: []string{"leading", "space"},
		},
		{
			input:    "trailing space   ",
			expected: []string{"trailing", "space"},
		},
		{
			input:    "  MIXED  Case WoRds ",
			expected: []string{"mixed", "case", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("slices not same size")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("words not same in slice")
			}
		}
	}
}
