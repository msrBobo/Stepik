package main

import (
	"fmt"
)

func convert(x int64)uint16{
	result := uint16(x)
	return result
}

func main() {

	var x int64
	results := convert(x)

	fmt.Printf("%T\n",results)
}