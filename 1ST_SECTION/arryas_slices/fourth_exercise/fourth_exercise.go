package main
import "fmt"

func main() {
	
	var N,p int
	var slices []int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&p)
		slices = append(slices, p)
	}
	for i := range(slices) {
		if i % 2 == 0 {
			fmt.Print(" ",slices[i])
		}
	}
}