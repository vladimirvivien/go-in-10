package main

import "fmt"

func main() {
	table := make(map[string]int)
	table["Apr"] = 100
	table["May"] = 445
	table["Jun"] = 514

	q2 := []int{table["Apr"], table["May"], table["Jun"]}

	fmt.Println(q2)
}
