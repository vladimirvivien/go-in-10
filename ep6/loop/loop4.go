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

	for _, emp := range emps {
		fmt.Printf("Name: %s, Title: %s\n", emp.Name, emp.Title)
	}
}
