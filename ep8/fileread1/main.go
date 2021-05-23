package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("irise.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	file2, err := os.Create("irise2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file2.Close()

	// copy file content to in-mem buffer
	//buf := new(bytes.Buffer)
	if _, err := io.Copy(file2, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print buff
	//fmt.Println(buf.String())
}
