package main

import (
	"fmt"
)

type direction string

type vehicle struct {
	make  string
	model string
}

func (v *vehicle) startEngine() {
	fmt.Println("Engine started.")
}

func (v *vehicle) drive(dist float32, dir direction) {
	fmt.Printf("Drive: %0.2f miles, %s\n", dist, dir)
}

func main() {
	truck := &vehicle{make: "ford", model: "f150"}
	truck.startEngine()
	truck.drive(10, direction("north"))
}
