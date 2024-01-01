package main

import (
	"fmt"
)

func main() {

	var N []int
	var p int
	for i := 0; i <= 4; i++ {
		fmt.Scan(&p)
		N = append(N, p)
	}
	fmt.Println(N[4])

}