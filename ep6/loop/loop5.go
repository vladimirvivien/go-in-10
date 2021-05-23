package main

import (
	"fmt"
)

type Histogram map[string]int

func main() {
	table := Histogram{
		"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
		"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
		"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312,
	}

	for month, value := range table {
		fmt.Printf("Month: %s, Sales: %d\n", month, value)
	}
}
