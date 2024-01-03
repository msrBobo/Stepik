package main
import "fmt"

func main() {

    var nums int
    fmt.Scan(&nums)

    frist := nums % 10
    second := (nums / 10) % 10
    there := nums / 100
    fmt.Printf("%d%d%d\n",frist, second, there)
}