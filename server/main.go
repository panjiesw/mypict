package main

import (
	"fmt"
	"time"
)

func now(t **time.Time) {
	tt := time.Now()
	*t = &tt
}

func main() {
	var t *time.Time
	now(&t)
	fmt.Printf("time: %v", t)
}
