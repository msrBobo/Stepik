package main
import "fmt"

func main() {
    var arr []int
    for i := 1; i != 0;i++{
        var nums int
        fmt.Scan(&nums)
        if nums < 10 {
            continue
        }else if nums > 100 {
            break
        }
        arr = append(arr, nums)
    }
    for i := range arr {
        fmt.Println(arr[i])
    }
}