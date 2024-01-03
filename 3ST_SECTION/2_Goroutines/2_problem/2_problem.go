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

	numRoutines := 10

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}

	wg.Wait()

	fmt.Println("All goroutines have completed their work.")
}
