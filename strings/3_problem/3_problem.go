package main

import (
	"fmt"
	"strings"
)

func findSubstring(x, s string) int {
	index := strings.Index(x, s)
	return index
}

func main() {
	var x, s string
	fmt.Scan(&x)
	fmt.Scan(&s)

	result := findSubstring(x, s)
	fmt.Println(result)
}
