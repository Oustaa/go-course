package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello David"
}

func main1() {
	ch := make(chan string)
	go greet(ch)

	ch <- "Hello John"

	log.Println(<-ch)
	log.Println(<-ch)
}
