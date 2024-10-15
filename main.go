package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("Filed to open file")
	}

	// strData := string(data)

	// lines := strings.Split(strData, "\n\n")

	// var givenChars []string

	// for _, b := range input {
	// 	givenChars = append(givenChars, lines[b-32])
	// }

	// for i := 0; i < len(givenChars); i++ {
	// 	for j := 0; j < len(givenChars[i]); j++ {
	// 		if givenChars[i][j] == '\n' {
	// 			fmt.Print()
	// 			break
	// 		}
	// 		fmt.Print(string(givenChars[i][j]))
	// 	}
	// }
}

func AsciiWord(word string) {
	for i := range word {
		AddLine(word[i])
	}
}

func AddLine(byte) {
}
