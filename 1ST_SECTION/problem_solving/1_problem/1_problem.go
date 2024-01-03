package main
import "fmt"

func main() {

    var nums int
    fmt.Scan(&nums)

    frist := nums % 10
    second := (nums / 10) % 10
    there := nums / 100
    result := frist + second + there
    fmt.Println(result)
}