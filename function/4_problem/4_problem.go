package main

import (
	"fmt"
)

func main() {

	results := fibonacci(4)

	fmt.Println(results)
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}