package main

import (
	"fmt"
)

func main() {

	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(nums ...int) (int, int) {
	sum := 0
	count := len(nums)

	for _, num := range nums {
		sum += num
	}

	return count, sum
}