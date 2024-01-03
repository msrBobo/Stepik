package main

import (
	"fmt"
	"errors"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	result, err := divide(a, b)

	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(result)
	}
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}