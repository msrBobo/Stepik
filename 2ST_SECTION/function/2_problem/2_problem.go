package main

import (
	"fmt"
)

func main() {

	p := minimumFromFour()
	fmt.Println(p)
}

func minimumFromFour()int{
    var nums int
	min := 10000
	for i := 0; i < 4; i++ {
		fmt.Scan(&nums)
		
		if min > nums {
            min = nums
        }
	}
	return min
}