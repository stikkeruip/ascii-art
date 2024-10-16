package main

import (
	"os"
)

var asciiLines []string

func main() {

	if !checkValidity() {
		return
	}

	// Then we load the appropriate banner to memory, in the form of lines

	loadBanner("standard.txt")

	// Then we print
	processString(os.Args[1])
}
