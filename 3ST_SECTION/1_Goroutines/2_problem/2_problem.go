package main

import (
	"fmt"
)

func task2(ch chan string, N string) {
    for i := 0;i != 5;i++{
        ch <- N + " "
    }
   close(ch)
}

func main() {
	ch := make(chan string)

	go task2(ch, "Hello")

	for result := range ch {
		fmt.Print(result)
	}
}