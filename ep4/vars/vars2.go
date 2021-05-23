package main

import "fmt"

var firstName = "Vladimir"
var lastName, title = "Vivien", "Software Engineer"
var age = 40

func main() {
	var count = 1

	fmt.Printf("%s %s, %s, age %d\n", firstName, lastName, title, age)
	fmt.Println("Received:", count)
}
