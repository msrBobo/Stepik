package main

import "fmt"

func removeDigit(number, digitToRemove int) int {
	result := 0
	position := 1

	for number > 0 {
		digit := number % 10

		if digit != digitToRemove {
			result += digit * position
			position *= 10
		}

		number /= 10
	}

	return result
}

func main() {
	var number, digitToRemove int

	fmt.Scan(&number)
	fmt.Scan(&digitToRemove)

	result := removeDigit(number, digitToRemove)
	fmt.Println(result)
}
