package main

import (
	"fmt"
	"math"
)

func findHypotenuse(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	c := findHypotenuse(a, b)
	fmt.Printf("%.0f\n", c)
}