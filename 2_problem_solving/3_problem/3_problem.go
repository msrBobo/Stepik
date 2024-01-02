package main

import (
	"fmt"

)

func findMaxDigit(s string) rune {
	maxDigit := '0'

	for _, char := range s {
		if char > maxDigit {
			maxDigit = char
		}
	}

	return maxDigit
}

func main() {
	var input string
	fmt.Scan(&input)

	maxDigit := findMaxDigit(input)
	fmt.Println(string(maxDigit))
}