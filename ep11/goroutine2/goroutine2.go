package main

import (
	"fmt"
)

// start of main routine (function)
func main() {
	// followings run serially within same
	// execution thread as main
	go count(10, 50, 10)
	go count(60, 100, 10)
	go count(110, 200, 20)
	fmt.Scanln() // blocks (for keyboard input)
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
