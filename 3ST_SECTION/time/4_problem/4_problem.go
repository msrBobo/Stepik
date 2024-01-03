package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {
	var mins, secs int
	fmt.Scanf("%d мин. %d сек.", &mins, &secs)

	duration := time.Duration(mins)*time.Minute + time.Duration(secs)*time.Second

	baseTime := time.Unix(now, 0).UTC()

	resultTime := baseTime.Add(duration)

	fmt.Println(resultTime.Format(time.UnixDate))
}