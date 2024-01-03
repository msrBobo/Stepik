package main

import (
	"fmt"
	"math"
)

func isRightTriangle(a, b, c float64) bool {
	return math.Pow(c, 2) == math.Pow(a, 2)+math.Pow(b, 2)
}

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if isRightTriangle(a, b, c) {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}