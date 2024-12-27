package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const string = "data"

	for i := 0; i < len(string); i++ {
		fmt.Printf("%x", string[i])
	}

	fmt.Println()

	fmt.Println("Rune count :", utf8.RuneCountInString(string))

	for i, runValue := range string {

		fmt.Printf("%#U", i, runValue)
	}
}
