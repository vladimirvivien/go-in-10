package main

import (
	"fmt"
)

func main() {
	// buffered channel of size 4
	ch := make(chan int)
	//go func() {
	ch <- 2
	//}()

	// read channel sequentially
	fmt.Println(<-ch)

}
