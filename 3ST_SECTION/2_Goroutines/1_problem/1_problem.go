package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	
	time.Sleep(2 * time.Second)
	fmt.Println("Work completed")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		work()
	}()

	wg.Wait()

	fmt.Println("Main goroutine: Waiting for the goroutine to complete...")
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine: Done")
}
