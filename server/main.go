package main

import (
	"fmt"
	"time"

	"github.com/ventu-io/go-shortid"
)

func now(t **time.Time) {
	tt := time.Now()
	*t = &tt
}

func main() {
	i := 1

	for i < 54 {
		fmt.Println(shortid.MustGenerate())
		i++
	}
}
