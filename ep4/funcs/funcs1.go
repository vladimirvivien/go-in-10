package main

import "fmt"

func main() {
	fmt.Print(message("vladimir"))
}

func message(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
