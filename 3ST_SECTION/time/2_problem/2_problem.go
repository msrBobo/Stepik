package main

import (
	"fmt"
	 "time"
	"bufio"
	"os"
)
func main() {
    o := bufio.NewScanner(os.Stdin)
	o.Scan()
	input := o.Text()


	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, input)
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	if t.Hour() >= 13 {
		t = t.AddDate(0, 0, 1)
	}

	// Natijani chiqarish
	fmt.Println(t.Format(layout))
}
