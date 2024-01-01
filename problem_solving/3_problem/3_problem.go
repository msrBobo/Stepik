package main

import "fmt"

func main() {
    var degrees int
    fmt.Scan(&degrees)

    hours := degrees / 3600

    minutes := (degrees % 3600) / 60

    fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}