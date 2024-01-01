package main
import "fmt"

func main() {

   var num1,num2 int
   sum := 0
   fmt.Scan(&num1, &num2)

   for i := num1; i <= num2; i++{
        sum += i
   }
   fmt.Println(sum)
   
}