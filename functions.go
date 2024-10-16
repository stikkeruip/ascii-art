package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

//func CheckValidity checks that we have exactly one, properly formed argument.

func checkValidity() bool {
	if len(os.Args) != 2 {
		fmt.Println("Error: Please provide exactly one argument. Refer to README.md for assistance.")
		return false
	}
	if isNotASCII(os.Args[1]) {
		fmt.Println("Error: Only ASCII characters or newline symbols (\\n) are allowed. Refer to README.md for more information.")
		return false
	}
	return true
}

func loadBanner(filename string) {
	contentEntire, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Then we split it to lines

	asciiLines = strings.Split(string(contentEntire), "\n")
}

func drawWord(word string) {

	for _, char := range word {
		asciiValue := int(char)
		startIndex := (asciiValue-32)*9 + 1

		// If output is empty, initialize it with 8 empty strings
		if len(output) == 0 {
			output = make([]string, 8)
		}

		for i := 0; i < 8; i++ {
			output[i] += asciiLines[startIndex+i]
		}
	}

	// Print the resulting ASCII art line by line
	for _, line := range output {
		fmt.Println(line)
	}

	fmt.Println(contentEntire)
}

// func isNotASCII checks if every rune of an arguments is within the ASCII range.

func isNotASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}

func processString(input string) {
	// Split the string by newlines, but retain them in the result
	parts := strings.Split(input, "\n")

	for i, part := range parts {
		if part != "" {
			// If it's a non-empty part, it's a regular string (call Function A)
			drawWord(part)
		}
		if i < len(parts)-1 {
			// If we're not at the last element, there's a newline between parts
			fmt.Println()
		}
	}
}
