package main

import (
	"flag"
	"fmt"
)

func main() {
	var username string
	var awesome bool

	flag.StringVar(&username, "username", "Vladimir", "the username")
	flag.BoolVar(&awesome, "isawesome", true, "is this an awesome user?")
	flag.Parse()

	if awesome {
		fmt.Printf("Hello %s, you are so awesome", username)
	} else {
		fmt.Println("Sorry only awesome users allowed.")
	}
}
