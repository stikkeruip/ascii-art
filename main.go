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
	charHeight = 8
)

func main() {
	// Ensure the user provides an input string
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input string.")
		return
	}

	input := os.Args[1]

	// replace '\n' with actual newline characters
	input = strings.ReplaceAll(input, `\n`, "\n")

	// load the ascii font from the file
	asciiChars, err := loadFont("standard.txt")
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}

	// create a map of ascii characters to their corresponding art
	asciiMap := createAsciiMap(asciiChars)

	// split the input string into lines based on newlines
	inputLines := strings.Split(input, "\n")
	for idx, inputLine := range inputLines {
		if inputLine == "" {
			// empty line due to consecutive '\n', print a newline
			fmt.Println()
			continue
		}

		// find what art is needed from input
		givenChars := buildAsciiArt(inputLine, asciiMap)
		// print ascii art
		printAsciiArt(givenChars)

		// print a newline after each line, except the last one
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

	// check if the first character in the file is a newline and skip it if necessary
	if len(strData) > 0 && strData[0] == '\n' {
		strData = strData[1:]
	}

	// now split by double newline to separate each character representation
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
		// check if an ascii value in the txt file
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

// prints the ascii art for the given characters line by line
func printAsciiArt(givenChars [][]string) {
	for lineIndex := 0; lineIndex < charHeight; lineIndex++ {
		lineContent := ""

		for _, charLines := range givenChars {
			// append the current line of each character if it exists, else add empty line
			if lineIndex < len(charLines) {
				lineContent += charLines[lineIndex]
			} else {
				lineContent += "    " // handle space or missing lines
			}
			lineContent += " " // adjust spacing
		}

		fmt.Println(lineContent)
	}
}
