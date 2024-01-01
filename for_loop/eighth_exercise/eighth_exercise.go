package main

import ("fmt"
        "strings"
)
func main() {
    var num1, num2 int
    fmt.Scan(&num1, &num2)

    result := findCommonDigits(num1, num2)

    fmt.Println(result)
}

func findCommonDigits(num1, num2 int) string {
    strNum1 := fmt.Sprint(num1)
    strNum2 := fmt.Sprint(num2)

    commonDigits := ""

    for _, digit1 := range strNum1 {
        strDigit1 := string(digit1)

        if contains(strNum2, strDigit1) && !contains(commonDigits, strDigit1) {
            commonDigits += strDigit1 + " "
        }
    }

    return commonDigits
}

func contains(s, substr string) bool {
    return fmt.Sprint(s) != fmt.Sprint(strings.ReplaceAll(s, substr, ""))
}