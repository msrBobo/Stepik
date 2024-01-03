package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var minElement, countMin int
    fmt.Scan(&minElement)
    countMin = 1 

    for i := 1; i < n; i++ {
        var num int
        fmt.Scan(&num)

        if num < minElement {
            minElement = num
            countMin = 1
        } else if num == minElement {
            countMin++
        }
    }

    fmt.Println(countMin)
}