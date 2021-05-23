package main

import (
	"fmt"
)

type Student struct {
	Name string
	GPA  grade
}

type grade float64
type history map[Student]int
type True bool
type description string

func main() {
	s1 := Student{Name: "Luke S.", GPA: 4.0}
	fmt.Printf("Name: %s, Grade :%f", s1.Name, s1.GPA)
}
