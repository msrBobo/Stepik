package main

import (
	"fmt"
)

func removeDuplicateChars(input string) string {
	charCount := make(map[rune]int)
	result := ""

	for _, char := range input {
		charCount[char]++
	}

	for _, char := range input {
		if charCount[char] == 1 {
			result += string(char)
		}
	}

	return result
}

func main() {
	var input string
	fmt.Scan(&input)
	result := removeDuplicateChars(input)
	fmt.Println(result)
}
