package main

import "fmt"

func main() {
    var degrees int
    fmt.Scan(&degrees)

    hours := degrees / 30

    minutes := (degrees % 30) * 2

    fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}