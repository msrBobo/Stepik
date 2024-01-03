package main

import (
	"fmt"
)

func main() {
	a,b := 2,4
	test(&a,&b)
}

func test(x1 *int, x2 *int) {
	p := *x1
	*x1 = *x2
	*x2 = p
	fmt.Println(*x1, *x2)
    
}