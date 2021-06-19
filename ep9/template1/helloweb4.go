package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// type use for templated page data
type PageData struct {
	Message string
	Time    string
}

// handler function handles incoming request
func hello(w http.ResponseWriter, r *http.Request) {
	// configure http response
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	pd := PageData{
		Message: "Go is awesome!",
		Time:    time.Now().Format(time.Kitchen),
	}

	// parse templated HTML file with data
	t, err := template.ParseFiles("hello.html")

	// if error, stop and return error immediately
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// execute and return the template
	err = t.Execute(w, pd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
