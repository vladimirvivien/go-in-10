package lib

import "time"

func LongTime() string {
	return time.Now().String()
}

func ShortTime() string {
	return time.Now().Format(time.Kitchen)
}
