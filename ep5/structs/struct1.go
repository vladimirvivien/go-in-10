package main

import (
	"fmt"
)

type planet struct {
	diam int
	name string
	desc string
}

func main() {
	p0 := planet{
		diam: 7926,
		name: "Earth",
		desc: "Third rock from the Sun",
	}

	var p1 planet
	p1.name = "Mars"
	p1.diam = 5332

	fmt.Print(p0.name)
}
