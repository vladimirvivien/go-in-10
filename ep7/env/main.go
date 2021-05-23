package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("You are watching episode:", os.Getenv("EP_VALUE"))
}
