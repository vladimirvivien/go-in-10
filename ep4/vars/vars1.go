package main

import "fmt"

var firstName string
var lastName, title string
var age int

func main() {
	var count int

	fmt.Printf("%s %s %s %d\n", title, firstName, lastName, age)
	fmt.Println("Received:", count)
}
