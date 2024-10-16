package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("Filed to open file")
	}
	strData := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(strData, "\n\n")

	var givenChars [][]string

	for _, b := range input {
		charLines := strings.Split(lines[b-32], "\n")
		givenChars = append(givenChars, charLines)
	}

	maxLines := 0
	for _, charLines := range givenChars {
		if len(charLines) > maxLines {
			maxLines = len(charLines)
		}
	}

	for lineIndex := 0; lineIndex < maxLines; lineIndex++ {
		for _, charLines := range givenChars {
			if lineIndex < len(charLines) {
				fmt.Print(charLines[lineIndex])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
