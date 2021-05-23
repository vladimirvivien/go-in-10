package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Stream content of a file
// and inspect each line as it is
// being read.
func main() {
	file, err := os.Open("irise.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	count := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(line)
			break
		}
		if count == 10 {
			fmt.Println(line)
			break
		}
		fmt.Print(strings.ToUpper(line))
		count++
	}
}
