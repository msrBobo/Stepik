package main

import "fmt"

func digitalRoot(n int) int {
    for n >= 10 {
        sum := 0
        for n > 0 {
            sum += n % 10
            n /= 10
        }
        n = sum
    }
    return n
}

func main() {
    var n int
    fmt.Scan(&n)

    result := digitalRoot(n)
    fmt.Println(result)
}