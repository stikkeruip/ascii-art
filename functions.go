package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var asciiMap map[rune][]string

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

// Load the ASCII art banner into a map
func loadBanner(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	asciiLines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n\n")
	asciiMap = make(map[rune][]string)

	// Manually handle the space character to be 8 empty lines
	asciiMap[' '] = []string{
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
	}

	// Create a map for ASCII characters (starting after the space character)
	for i, art := range asciiLines[1:] { // Start from index 1 to skip space
		char := rune(33 + i) // ASCII starts at 33 after space
		asciiMap[char] = strings.Split(art, "\n")
	}
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

// Process the input string and print the ASCII art
func processString(input string) {
	// Replace literal "\n" with newlines and split into lines
	input = strings.ReplaceAll(input, `\n`, "\n")
	inputLines := strings.Split(input, "\n")

	// Process each line separately
	for idx, line := range inputLines {
		if line == "" {
			fmt.Println() // Handle empty lines (newline in the ASCII art)
			continue
		}

		// Build and print the ASCII art for the line
		printAsciiArt(buildAsciiArt(line))

		// Add a newline between input lines
		if idx < len(inputLines)-1 {
			fmt.Println()
		}
	}
}

// Build the ASCII art for a given line of input
func buildAsciiArt(line string) [][]string {
	var asciiChars [][]string
	for _, char := range line {
		// Handle space character
		if char == ' ' {
			asciiChars = append(asciiChars, asciiMap[' '])
			continue
		}

		// Skip characters not in the ASCII range
		if char < 32 || char > 126 {
			continue
		}
		asciiChars = append(asciiChars, asciiMap[char])
	}
	return asciiChars
}

// Print the ASCII art for the given characters
func printAsciiArt(asciiChars [][]string) {
	maxLines := 8 // Assuming each character has 8 lines
	for i := 0; i < maxLines; i++ {
		for _, charLines := range asciiChars {
			fmt.Print(charLines[i]) // Print each line of each character
			// fmt.Print(" ")          // Add a space between characters
		}
		fmt.Println() // Move to the next line of the ASCII art
	}
}
