package main

import (
	"fmt"
	"unicode"
)

func isValidPassword(password string) bool {
	if len(password) < 5 {
		return false
	}

	for _, char := range password {
		if !(unicode.IsDigit(char) || (unicode.IsLetter(char) && (unicode.Is(unicode.Latin, char)))) {
			return false
		}
	}

	return true
}

func main() {
	var password string
	fmt.Scan(&password)

	if isValidPassword(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}