package main

import (
	"fmt"
	"strings"
	"time"
)

type Histogram map[string]int

func main() {
	table := Histogram{
		"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
		"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
		"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312,
	}

	for month, value := range filterByMonth(table, "Aug") {
		fmt.Printf("Month: %s, Sales: %d\n", month, value)
	}
}

func filterByMonth(histo Histogram, mo string) Histogram {
	defer fmt.Println(time.Now())
	result := make(Histogram)
	for month, val := range histo {
		if strings.HasPrefix(month, mo) {
			result[month] = val
		}
	}
	return result
}
