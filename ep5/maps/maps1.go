package main

import "fmt"

func main() {
	table := map[string]int{
		"Jan": 100, "Feb": 445, "Mar": 514, "Apr": 233,
		"May": 321, "Jun": 644, "Jul": 113, "Aug": 734,
		"Sep": 553, "Oct": 344, "Nov": 831, "Dec": 312,
	}

	q2 := []int{table["Apr"], table["May"], table["Jun"]}

	fmt.Println(q2)
}
