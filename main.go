package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	asciiOffset = 32
	asciiHeight = 8
)

var asciiMap map[byte][]string

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

// Load the ASCII art banner into a map
func loadBanner(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}
	strData := strings.ReplaceAll(string(data), "\r\n", "\n")
	// Check if the first character in the data is a newline and skip it if necessary
	if len(strData) > 0 && strData[0] == '\n' {
		strData = strData[1:]
	}
	asciiLines := strings.Split(strData, "\n\n")
	asciiMap = make(map[byte][]string)
	// Create a map for ASCII characters (starting from the space character)
	for i, art := range asciiLines {
		char := byte(asciiOffset + i) // ASCII code starts at 32 for space
		asciiMap[char] = strings.Split(art, "\n")
	}
}

// Process the input string and print the ASCII art
func processString(input string) {
	// Replace literal "\n" with actual newlines and split into lines
	input = strings.ReplaceAll(input, `\n`, "\n")
	inputLines := strings.Split(input, "\n")
	// Process each line separately
	for _, line := range inputLines {
		if line == "" {
			fmt.Println() // Handle empty lines (newline in the ASCII art)
			continue
		}
		// Build and print the ASCII art for the line
		printAsciiArt(buildAsciiArt(line))
	}
}

// Build the ASCII art for a given line of input
func buildAsciiArt(line string) [][]string {
	var asciiChars [][]string
	for _, char := range []byte(line) { // Use byte slice to process the input
		asciiChars = append(asciiChars, asciiMap[char])
	}
	return asciiChars
}

// Print the ASCII art for the given characters
func printAsciiArt(asciiChars [][]string) {
	for i := 0; i < asciiHeight; i++ {
		for _, charLines := range asciiChars {
			fmt.Print(charLines[i]) // Print each line of each character
			// Uncomment the next line if you want spaces between characters
			// fmt.Print(" ")
		}
		fmt.Println() // Move to the next line of the ASCII art
	}
}

func main() {
	// Check the validity of the input
	if !checkValidity() {
		return
	}
	// Load the banner (ASCII art)
	loadBanner("thinkertoy.txt")
	// Process the input string
	processString(os.Args[1])
}
