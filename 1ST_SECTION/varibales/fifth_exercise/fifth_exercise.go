package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    tens := (n / 10) % 10

    fmt.Println(tens)
}