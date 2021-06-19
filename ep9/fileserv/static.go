package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver
	http.Handle("/", http.FileServer(http.Dir(".")))

	// start server (stop and log error if any)
	log.Println("serving static files on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
