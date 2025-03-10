package main

// import (
// 	"fmt"
// )

// func main() {
// 	a, b := 5, 10
// 	const x int = 12

// 	// x = 223

// 	swap(&a, &b)

// 	fmt.Println(a == 10, b == 5)
// }

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
