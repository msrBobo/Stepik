package main

import "fmt"

func main() {
    var n, num, sum int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Scan(&num)
        if num >= 10 && num <= 99 && num%8 == 0 {
            sum += num
        }
    }

    fmt.Println(sum)
}