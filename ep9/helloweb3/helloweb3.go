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

	message := "Hello, from Go!!"

	fmt.Fprintf(w, `<!DOCTYPE html><html>
	<head>
	  <title>Hello</title>
	</head>
	
	<body>
		<h1>%s</h1>
	</body></html>	
	`, message)
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
