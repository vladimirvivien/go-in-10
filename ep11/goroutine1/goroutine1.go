package main

import (
	"fmt"
)

// start of main routine (function)
func main() {
	// followings run concurrently
	go count(10, 50, 10)
	go count(60, 100, 10)
	go count(110, 200, 20)
}

func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}
