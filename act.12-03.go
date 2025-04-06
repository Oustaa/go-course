package main

import (
	"fmt"
	"time"
)

func m34122ain() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()
	diff := end.Sub(start)
	duration := diff.Seconds()

	fmt.Printf("The execution took exactly %f seconds!\n", duration)
}
