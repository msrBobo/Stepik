package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "os"
)

func main() {
	var data map[string]interface{}

	inputData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
	}

	err = json.Unmarshal(inputData, &data)
	if err != nil {
	}

	students := data["Students"].([]interface{})
	totalRatings := 0
	totalStudents := len(students)

	for _, studentData := range students {
	student := studentData.(map[string]interface{})
	ratings := student["Rating"].([]interface{})
	totalRatings += len(ratings)
	}

	averageRatings := float64(totalRatings) / float64(totalStudents)

	result := map[string]float64{
	"Average": averageRatings,
	}

	jsonResult, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
	}

 fmt.Println(string(jsonResult))
}
