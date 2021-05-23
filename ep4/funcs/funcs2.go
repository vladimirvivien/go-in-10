package main

import "fmt"

func main() {
	fmt.Print(message("Vladimir"))
}

func message(name string) string {
	filter := func(name string) string {
		if name == "Vladimir" {
			return "Vladimir Vivien"
		}
		return name
	}
	return fmt.Sprintf("Hello, %s!", filter(name))
}
