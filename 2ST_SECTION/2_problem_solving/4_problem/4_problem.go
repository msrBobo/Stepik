package main

import (
	"fmt"
	"strconv"
	
)

func main() {

	var nums int
	fmt.Scan(&nums)
	str := strconv.Itoa(nums)
	var d uint8
	for i := range str {
		d = (str[i]-48) * (str[i]-48)
		fmt.Print(d)
	}
}