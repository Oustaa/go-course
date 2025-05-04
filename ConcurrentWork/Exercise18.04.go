// package main

// import (
// 	"log"
// )

// func greet(ch chan string) {
// 	ch <- "Hello"
// }

// func main() {
// 	ch := make(chan string)
// 	go greet(ch)

// 	log.Println(<-ch)
// }

package main

import "log"

func greeting(ch chan string) {
	log.Println("greeting: ", <-ch)
	ch <- "Hello Me, I am him"
}

func mai2qn() {
	ch := make(chan string)
	ch <- "I am the main func"
	go greeting(ch)

	log.Println(<-ch)
}
