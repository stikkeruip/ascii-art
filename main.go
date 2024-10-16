package main

import (
	"os"
)

func main() {
	// Check validity
	if !checkValidity() {
		return
	}

	// Load the banner (ASCII art)
	loadBanner("standard.txt")

	// Process the input string
	processString(os.Args[1])
}
