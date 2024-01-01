package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)

    result := (a*a) + (b*b)

    fmt.Println(result)
}