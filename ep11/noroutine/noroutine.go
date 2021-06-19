package main

import (
	"fmt"
)

// start of main routine (function)
func main() {
	// followings run serially within same
	// execution thread as main
	count(10, 50, 10)
	count(60, 100, 10)
	count(110, 200, 20)
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
