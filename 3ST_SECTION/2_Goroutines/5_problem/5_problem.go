package main

import (
	"fmt"
	"sync"
)


func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var (
		slice1   = make([]int, n)
		slice2   = make([]int, n)
		ch sync.WaitGroup
		p sync.Mutex
	)
	ch.Add(n * 2)
	go func() {
		for i := 0; i < n; i++ {
			p.Lock()
			slice1[i] = <-in1
			p.Unlock()
			ch.Done()
		}
	}()
	go func() {
		for i := 0; i < n; i++ {
			p.Lock()
			slice2[i] = <-in2
			p.Unlock()
			ch.Done()
		}
	}()
	go func() {
		for i, v := range slice1 {
			slice1[i] = fn(v)
		}
	}()
	go func() {
		for i, v := range slice2 {
			slice2[i] = fn(v)
		}
	}()
	go func() {
		defer close(out)
		ch.Wait()
		for i := 0; i < n; i++ {
			out <- slice1[i] + slice2[i]
		}
	}()
}

func main() {
	in1 := make(chan int, 6)
	in2 := make(chan int, 6)
	out := make(chan int, 6)

	for i := 0; i <= 5; i++ {
		in1 <- i
		in2 <- i
	}

	fn := func(n int) int {
		return n
	}

	merge2Channels(fn, in1, in2, out, 6)

	for i := 0; i <= 5; i++ {
		fmt.Println(<-out)
	}
}