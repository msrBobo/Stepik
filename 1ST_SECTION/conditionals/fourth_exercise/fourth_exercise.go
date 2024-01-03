package main
import "fmt"

func main() {

    var nums int
    fmt.Scan(&nums)
    digit1 := nums % 10
    digit2 := (nums / 10) % 10
    digit3 := (nums / 100) % 10
    digit4 := (nums / 1000) % 10
    digit5 := (nums / 10000) % 10
    digit6 := (nums / 100000)
    if digit1 + digit2 + digit3 == digit4 + digit5 + digit6{
        fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}