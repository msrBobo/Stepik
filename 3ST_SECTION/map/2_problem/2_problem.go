package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   {"CityA", "CityB"},  
		100:  {"CityX", "CityY"},   
		1000: {"CityM", "CityN"},
	}

	cityPopulation := map[string]int{
		"CityA": 50,
		"CityB": 80,
		"CityX": 500,
		"CityY": 800,
		"CityM": 1200,
		"CityN": 1500,
	}

	for populationRange, cities := range groupCity {
		if populationRange != 100 {
			for _, city := range cities {
				delete(cityPopulation, city)
			}
		}
	}

	fmt.Println(cityPopulation)
}
