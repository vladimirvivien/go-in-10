package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// handler function handles incoming request
func hello(w http.ResponseWriter, r *http.Request) {
	// configure http response
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	// write response
	fmt.Fprintf(w, "Hello, I love Go!")
}

func main() {
	http.HandleFunc("/hello", hello)
	// explicitly declare server to control
	// server behavior
	server := http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 3,
	}

	// start server, (or stop and log error if any)
	log.Fatal(server.ListenAndServe())
}
