package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	if isPalindrom(text) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func isPalindrom(s string) bool {
	s = strings.TrimRight(s, "\n")

	newRs := reverseString(s)
	p := []rune(s)
	rs := []rune(newRs)
	for i := 0; i < len(p); i++ {
        if p[i]!= rs[i] {
            return false
        }
    }
	return true
}

func reverseString(input string) string {
    runes := []rune(input)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}