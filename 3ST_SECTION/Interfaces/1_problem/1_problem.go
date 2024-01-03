package main

import (
	"fmt"
)

func readTask() (value1, value2, operation interface{}) {

	return 5.0, 5.6, "/"
	
	}

func main() {
	a, b, op := readTask()

	aFloat, aFloatOK := a.(float64)
	if !aFloatOK {
		fmt.Printf("value=%v: %T\n", a, a)
		return
	}

	bFloat, bFloatOK := b.(float64)
	if !bFloatOK {
		fmt.Printf("value=%v: %T\n", b, b)
		return
	}

	opString, opStringOK := op.(string)
	if !opStringOK {
		fmt.Printf("value=%v: %T\n", op, op)
		return
	}

	switch opString {
	case "+", "-", "*", "/":
		result := 0.0
		switch opString {
		case "+":
			result = aFloat + bFloat
		case "-":
			result = aFloat - bFloat
		case "*":
			result = aFloat * bFloat
		case "/":
			if bFloat == 0 {
				fmt.Println("деление на ноль")
				return
			}
			result = aFloat / bFloat
		}
		fmt.Printf("%.4f\n", result)
	default:
		fmt.Println("неизвестная операция")
		return
	}
}
