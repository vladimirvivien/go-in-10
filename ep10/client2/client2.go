package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// A simple HTTP client that uses user-defined client instance
func main() {
	// declare a client
	client := &http.Client{}
	// Use default client package to retrieve http payload
	resp, err := http.Get("http://en.wikipedia.org/wiki/Jean-Jacques_Dessalines")
	if err != nil {
		log.Fatalf("http failed: %s", err)
	}
	defer resp.Body.Close() // practice good hygiene, and close resource

	// extract text from response body
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print
	fmt.Println(string(text))
}
