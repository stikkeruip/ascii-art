package main

import (
	"fmt"
	"os"
	"unicode"
)

// Check the validity of the input
func checkValidity() bool {
	if len(os.Args) != 2 {
		fmt.Println("Error: Please provide exactly one argument.")
		return false
	}
	if isNotASCII(os.Args[1]) {
		fmt.Println("Error: Only ASCII characters or newline symbols (\\n) are allowed.")
		return false
	}
	return true
}

// Check if the input contains non-ASCII characters
func isNotASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}
