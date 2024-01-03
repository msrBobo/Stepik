package main

import (
	"fmt"
	
)

func main() {
	var x string
	fmt.Scan(&x)
	rs := []rune(x)
	var new []rune
	for i := range rs {
        if i % 2 != 0 {
			new = append(new, rs[i])
		}
    }
	fmt.Println(string(new))

}