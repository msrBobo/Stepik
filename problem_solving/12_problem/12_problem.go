package main

import "fmt"

func main() {
    
    var n int
    fmt.Scan(&n)

    powerOfTwo := 1
    for powerOfTwo <= n {
        fmt.Print(powerOfTwo, " ")
        powerOfTwo *= 2
    }

    fmt.Println()
}