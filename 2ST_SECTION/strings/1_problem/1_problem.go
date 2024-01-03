package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	if isValidString(text) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

func isValidString(s string) bool {
	s = strings.TrimRight(s, "\n")

	rs := []rune(s)

	if len(rs) == 0 {
		return false
	}

	if !unicode.IsUpper(rs[0]) {
		return false
	}

	if rs[len(rs)-1] != '.' {
		return false
	}

	return true
}