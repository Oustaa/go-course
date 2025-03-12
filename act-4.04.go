package main

import "fmt"

type muInt int16

func main() {
	words := []string{"Good", "Good", "bad", "Good", "Good"}
	words = append(words[:2], words[3:]...)

	fmt.Println(words)
}
