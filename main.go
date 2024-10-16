package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	asciiStart = 32
	asciiEnd   = 126
)

func main() {
	// Ensure the user provides an input string
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input string.")
		return
	}

	input := os.Args[1]

	// Check if the input is empty or only whitespace
	if len(strings.TrimSpace(input)) == 0 {
		fmt.Println("Input is empty. Nothing to display.")
		return
	}

	// Replace '\n' with actual newline characters
	input = strings.ReplaceAll(input, `\n`, "\n")

	// Load the ascii font from the file
	asciiChars, err := loadFont("standard.txt")
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}

	// Create a map of ascii characters to their corresponding art
	asciiMap := createAsciiMap(asciiChars)

	// Split the input string into lines based on newlines
	inputLines := strings.Split(input, "\n")
	for idx, inputLine := range inputLines {
		if inputLine == "" {
			// Empty line due to consecutive '\n', print a newline
			fmt.Println()
			continue
		}

		// find what art is needed from input
		givenChars := buildAsciiArt(inputLine, asciiMap)
		//print ascii art
		printAsciiArt(givenChars)

		// Print a newline after each line
		if idx < len(inputLines)-1 {
			fmt.Println()
		}
	}
}

// loads the ascii art from the file path
func loadFont(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	strData := strings.ReplaceAll(string(data), "\r\n", "\n")
	asciiChars := strings.Split(strData, "\n\n")
	return asciiChars, nil
}

// creates a map of ascii characters to their corresponding art
func createAsciiMap(asciiChars []string) map[rune][]string {
	asciiMap := make(map[rune][]string)
	for i, art := range asciiChars {
		char := rune(asciiStart + i)
		asciiMap[char] = strings.Split(art, "\n")
	}
	return asciiMap
}

// creates a string of the ascii art needed to print based off the input
func buildAsciiArt(line string, asciiMap map[rune][]string) [][]string {
	var givenChars [][]string
	for _, b := range line {
		// Check if an ascii value in the txt file
		if b < asciiStart || b > asciiEnd {
			continue
		}
		charLines, exists := asciiMap[b]
		if !exists {
			fmt.Printf("Warning: Character '%c' not found in font data.\n", b)
			continue
		}
		givenChars = append(givenChars, charLines)
	}
	return givenChars
}

// printAsciiArt prints the ascii art for the given characters line by line
func printAsciiArt(givenChars [][]string) {
	maxLines := getMaxLines(givenChars)
	for lineIndex := 0; lineIndex < maxLines; lineIndex++ {
		for _, charLines := range givenChars {
			// Print the current line of each character if it exists
			if lineIndex < len(charLines) {
				fmt.Print(charLines[lineIndex])
			}
			fmt.Print(" ") // Adjust spacing
		}
		fmt.Println()
	}
}

// calculates the maximum number of lines in the given characters
func getMaxLines(givenChars [][]string) int {
	maxLines := 0
	for _, charLines := range givenChars {
		if len(charLines) > maxLines {
			maxLines = len(charLines)
		}
	}
	return maxLines
}
