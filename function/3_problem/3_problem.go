package main

import (
	"fmt"
)

func main() {

	results := vote(0,0,1)

	fmt.Println(results)
}

func vote(x int, y int, z int)int{
	var zero = 0
	var one = 0
	var slices []int
	slices = append(slices, x, y, z)
	for _,i := range(slices) {
		if i == 0{
			zero++
		}else{
			one++
		}
		if one < zero{
			return 0
		}
		

	}
	
	return 1
}