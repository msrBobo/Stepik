package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	
	go func (){
        defer close(ch)
			select {
				case c := <-firstChan:
					ch <- c * c
				case c := <-secondChan:
					ch <- c * 3
				case <-stopChan:
					break
			}
	}()

	return ch
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	resultChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		for i := 1; i <= 5; i++ {
			firstChan <- i
			secondChan <- i
			time.Sleep(time.Second)
		}

		close(stopChan)
	}()

	for result := range resultChan {
		fmt.Println("Result:", result)
	}

	fmt.Println("Main: Calculator goroutine has stopped.")
}
