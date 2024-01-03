package main

import (
	"fmt"
	"time"
)

func main() {
	var input string
	fmt.Scan(&input)

	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	fmt.Println(parsedTime.Format(time.UnixDate))
}
