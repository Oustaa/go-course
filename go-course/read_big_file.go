package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	content, err := os.ReadFile("numbers.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(content))

	duration := time.Since(start)
	fmt.Printf("Performance: %v\n", duration)
}
