package main

import "fmt"

func isTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && a+c > b
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if isTriangle(a, b, c) {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}