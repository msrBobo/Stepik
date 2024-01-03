package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	o := bufio.NewScanner(os.Stdin)
	o.Scan()
	input := o.Text()

	a := strings.Split(input, ",")

	p := strings.TrimSpace(a[0])
	m := strings.TrimSpace(a[1])


	past1, err := time.Parse("02.01.2006 15:04:05", p)
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	future2, err := time.Parse("02.01.2006 15:04:05", m)
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	if past1.After(future2) {
		duration := past1.Sub(future2)
		fmt.Println(duration)
	}
	if past1.Before(future2) {
		duration := future2.Sub(past1)
		fmt.Println(duration)
	}
}

