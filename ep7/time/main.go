package main

import (
	"flag"
	"fmt"
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

func main() {
	// default format
	format := "datetime-12"

	// setup program flag
	flag.StringVar(&format, "format", format, "The name of the time format (i.e. time-12, time-24, datetime-12, datetime-24")
	flag.Parse()

	// retrieve actual format from map
	dateFmt := formats[format]

	fmt.Println(formatTime(dateFmt))
}

func formatTime(timeFmt string) string {
	return time.Now().Format(timeFmt)
}
