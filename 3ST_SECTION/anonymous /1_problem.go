package main

import "fmt"

func main() {
	result := uint(0)
	fn := func(num uint) uint {
		
		multiplier := uint(1)

		for num > 0 {
			digit := num % 10
			num /= 10

			if digit%2 == 0 && digit != 0 {
				result += digit * multiplier
				multiplier *= 10
			}
		}

		if result == 0 {
			return 100
		}

		return result
	}

	_ = fn(727178)

	fmt.Println(result)
}