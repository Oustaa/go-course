package main

import "fmt"

func main() {
	result := add(add(1.66,2.55), 0.79)

	fmt.Println(result)
}

func add[T string | int | float64 | float32](a, b T) T {
	return a + b
}