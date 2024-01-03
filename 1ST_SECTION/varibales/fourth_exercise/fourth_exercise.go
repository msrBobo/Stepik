package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    lastDigit := n % 10

    fmt.Println(lastDigit)
}