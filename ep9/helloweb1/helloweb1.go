package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler function handles incoming request
// and returns response
func hello(w http.ResponseWriter, r *http.Request) {
	// write response once (any subsequent call will overwrite response)
	fmt.Fprintf(w, "Hello, I love Go!")
}

func main() {
	// define the route of the request and
	// the handler to handle the response
	http.HandleFunc("/hello", hello)

	// start the server to listen to request on port 8080,
	// then wait for error (if any, log it and stop)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
