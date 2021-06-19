package main

import (
	"fmt"
)

func main() {
	// buffered channel of size 4
	workResult := work(2, 4, 6, 8, 10)

	// wait for result
	res := <-workResult
	fmt.Println(res)

}

func work(vals ...int) chan int {
	resChan := make(chan int)
	go func() {
		defer close(resChan)
		total := 0
		for _, val := range vals {
			total = total + val
		}
		resChan <- total
	}()
	return resChan
}
