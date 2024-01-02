package main

import (
	"fmt"
	"strings"
)

func addStarsBetweenLetters(s string) string {
	return strings.Join(strings.Split(s, ""), "*")
}

func main() {
	var input string
	fmt.Scan(&input)

	result := addStarsBetweenLetters(input)
	fmt.Println(result)
}