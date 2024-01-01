package main
import ("fmt"
		"slices"
	)

func main()  {
	var array []int
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array = append(array, a)
		
	}
	max := slices.Max(array)
	fmt.Println(max)
	
}