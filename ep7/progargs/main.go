package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Prog name:", os.Args[0])
	fmt.Println("Othre args:", os.Args[1:])
}
