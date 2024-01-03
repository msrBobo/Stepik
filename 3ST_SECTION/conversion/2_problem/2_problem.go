package main

import (
	"fmt"
	"unicode"
	"strconv"

)

func adding(s1, s2 string) int64 {
	
	cleanAndConvert := func(s string) int64 {
		var result string
		for _, char := range s {
			if unicode.IsDigit(char) {
				result += string(char)
			}
		}

		num, _ := strconv.ParseInt(result, 10, 64)
		return num
	}

	num1 := cleanAndConvert(s1)
	num2 := cleanAndConvert(s2)

	return num1 + num2
}

func main() {

	results := adding("%^80","hhhhh20&&&&nd")

	fmt.Println(results)

}
