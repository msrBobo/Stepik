package main

import (
	"fmt"
	"time"
)

func work(n int) int {
	if n > 3 {
	   time.Sleep(time.Millisecond * 500)
	   return n + 1
	} else {
	   time.Sleep(time.Millisecond * 500)
	   return n - 1
	}
 }

 func main() {

	m := make(map[int]int)
	var a int

	for i := 0;i < 10;i++{
		fmt.Scan(&a)
		if m[a] == 0{
			m[a] = work(a)
			
		}
		fmt.Print(m[a]," ")	
	}
	fmt.Print("time limit ok\n")
}