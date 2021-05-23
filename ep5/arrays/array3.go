package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{4, 5, 6}
	a1 = a2
	a1[0] = 4

	fmt.Println(a1)
}
