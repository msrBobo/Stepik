package main

import (
	"fmt"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var temp string
	for {
		value, ok := <-inputStream
		if !ok {
			break
		}
		if value != temp {
			outputStream <- value
			temp = value
		}
	}
}

func main() {
	
	input := []string{"apple", "apple", "orange", "orange", "banana", "banana", "banana", "grape"}
	inputStream := make(chan string, len(input))
	outputStream := make(chan string)

	go func() {
		for _, item := range input {
			inputStream <- item
		}
		close(inputStream)
	}()

	go removeDuplicates(inputStream, outputStream)

	for result := range outputStream {
		fmt.Println(result)
	}
}
