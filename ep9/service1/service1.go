package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	formats = map[string]string{
		"time-12":     "3:04PM",
		"time-24":     "15:04",
		"date-short":  "02-Jan-06",
		"datetime-12": "02-Jan-06 03:04PM MST",
		"datetime-24": "02-Jan-06 15:04 MST",
	}
)

func timeServ(w http.ResponseWriter, r *http.Request) {
	format := "datetime-12"

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `<!DOCTYPE html><html><head>
	  <title>Time Server</title>
	</head><body><h1>%s</h1></body></html>
	`, formatTime(format))
}

func formatTime(formatKey string) string {
	return time.Now().Format(formats[formatKey])
}

func main() {
	log.Println("starting time service on port 8080...")
	http.HandleFunc("/time", timeServ)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
