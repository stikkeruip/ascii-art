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

	lines := strings.Split(strData, "\r\n")

	var givenChars []string

	for _, b := range input {
		givenChars = append(givenChars, lines[b-32])
	}

	for i := 0; i < len(givenChars); i++ {
		for j := 0; j < len(givenChars[i]); j++ {
			if givenChars[i][j] == '\n' {
				fmt.Print()
				break
			}
			fmt.Print(string(givenChars[i][j]))
		}
	}
}
