package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	formats = map[string]string{
		"time-12":     "3:04PM MST",
		"time-24":     "15:04",
		"date-short":  "02-Jan-06",
		"datetime-12": "02-Jan-06 03:04PM MST",
		"datetime-24": "02-Jan-06 15:04 MST",
	}
)

func timeServ(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dt := formatTime("date-short")
	tm := formatTime("time-24")
	fmt.Fprintf(w, `{"date":"%s", "time":"%s"}`, dt, tm)
}

func formatTime(formatKey string) string {
	return time.Now().Format(formats[formatKey])
}

func main() {
	log.Println("starting time service on port 8080...")
	http.HandleFunc("/time", timeServ)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
