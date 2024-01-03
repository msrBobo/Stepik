package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    var word string
    if n%10 == 1 && n%100 != 11 {
        word = "korova"
    } else if (n%10 == 2 || n%10 == 3 || n%10 == 4) && (n%100 < 10 || n%100 > 20) {
        word = "korovy"
    } else {
        word = "korov"
    }

    fmt.Printf("%d %s\n", n, word)
}