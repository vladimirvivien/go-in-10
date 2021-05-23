package main

import (
	"fmt"
)

type Employee struct {
	Name  string
	Title string
}

func main() {
	emps := []Employee{
		{Name: "Vladimir", Title: "SE"},
		{Name: "Markus", Title: "Accountant"},
		{Name: "Selena", Title: "DE"},
	}

	for i, emp := range emps {
		fmt.Printf("%d. Name: %s, Title: %s\n", i, emp.Name, emp.Title)
	}
}
