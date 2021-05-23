package main

import (
	"fmt"
)

type Employee struct {
	Name  string
	Title string
}

func main() {
	var emps []Employee
	emps = []Employee{
		Employee{Name: "Vladimir", Title: "SE"},
		Employee{Name: "Markus", Title: "Accountant"},
		Employee{Name: "Selena", Title: "DE"},
	}

	for i := 0; i < len(emps); i++ {
		fmt.Printf("Name: %s, Title: %s\n", emps[i].Name, emps[i].Title)
	}
}
