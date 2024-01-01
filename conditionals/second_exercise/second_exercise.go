package main

import "fmt"

func main() {
    var number int
    fmt.Scan(&number)

    digit1 := number % 10
    digit2 := (number / 10) % 10
    digit3 := number / 100

    if digit1 != digit2 && digit2 != digit3 && digit1 != digit3 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}