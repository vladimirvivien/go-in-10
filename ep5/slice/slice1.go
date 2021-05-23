package main

import "fmt"

func main() {
	months := []string{
		"Jan", "Feb", "Mar", "Apr",
		"May", "Jun", "Jul", "Aug",
		"Sep", "Oct", "Nov", "Dec",
	}

	months = append(months, "mymonth")

	fmt.Println(len(months))

	fmt.Print(months[3])
}
