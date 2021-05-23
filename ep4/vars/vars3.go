package main

import "fmt"

func main() {
	firstName := "Vladimir"
	lastName, title := "Vivien", "Software Engineer"
	age := 40
	count := 1

	fmt.Printf("%s %s, %s, age %d\n", firstName, lastName, title, age)
	fmt.Println("Received:", count)
}
