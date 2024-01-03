package main

import (
	"fmt"
	"time"
)

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	result := make(chan int)
    c := 0
    go func(){
        defer close(result)
    for {     
        select {
            
            case v := <-arguments:
                c += v
            case <-done:
                result <-c
                return
        }
                
    }
    }()
	return result
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})

	resultChan := calculator(arguments, done)

	go func() {
		for i := 1; i <= 5; i++ {
			arguments <- i
			time.Sleep(time.Second)
		}

		close(done)
	}()

	result := <-resultChan
	fmt.Println("Final Result:", result)
}
