package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    maxNumber := -1
    for i := b; i >= a; i-- {
        if i%7 == 0 {
            maxNumber = i
            break
        }
    }

    if maxNumber != -1 {
        fmt.Println(maxNumber)
    } else {
        fmt.Println("NO")
    }
}