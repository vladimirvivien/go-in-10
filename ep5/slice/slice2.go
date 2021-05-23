package main

import "fmt"

func main() {
	half := make([]string, 3)
	half[0] = "Jan"
	half[1] = "Feb"
	half[2] = "Mar"
	// append
	half = append(half, "May", "Jun", "Jul")

	fmt.Printf("Months: %d\n", len(half))
	fmt.Println(half[3])
}
