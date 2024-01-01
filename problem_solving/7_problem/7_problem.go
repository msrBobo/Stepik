package main
import "fmt"

func main() {

    var N,P int
    fmt.Scan(&N)
    var count int
    for i := 0; i < N; i++ {
        fmt.Scan(&P)
        if P == 0 {
            count++
        }
    }
    fmt.Println(count)
    
}