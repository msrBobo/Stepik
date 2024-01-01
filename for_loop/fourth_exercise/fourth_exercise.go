package main

import "fmt"

func main() {
    var num, max, count int

    fmt.Scan(&num)
    max = num

    for num != 0 {
        if num > max {
            max = num
            count = 1 
        } else if num == max {
            count++ 
        }

        fmt.Scan(&num)
    }

    fmt.Println(count)
}