package main
import "fmt"

func main() {
	
	var N,p,count int

	var slices []int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&p)
		slices = append(slices, p)
	}
	for i := range(slices) {
		if slices[i] >= 0 {
			count++
		}
	}
	fmt.Println(count)
}