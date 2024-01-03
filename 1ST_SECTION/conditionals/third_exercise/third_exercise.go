package main
import "fmt"

func main() {

    var num int
    var sum int
    fmt.Scan(&num)
    for num > 0 {
        sum = num % 10
        num = num / 10
    }
    fmt.Println(sum)
}