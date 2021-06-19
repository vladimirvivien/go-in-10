package main

import (
	"fmt"
)

func main() {
	waitCh := make(chan bool)

	// do work here
	go func() {
		total := 0
		for _, val := range []int{2, 4, 6, 8} {
			total *= val
		}
		close(waitCh)
	}()

	// wait here
	<-waitCh
	fmt.Println("work done")
}
