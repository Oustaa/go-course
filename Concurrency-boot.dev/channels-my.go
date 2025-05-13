package main

import "fmt"

func send69(ch chan int) {
	fmt.Println("DURRING SENDING")
	ch <- 69
	fmt.Println("AFTER DURRING SENDING")
}

func main() {
	ch := make(chan int)

	fmt.Println("BEFFOR SENDING")
	go send69(ch)
	fmt.Println("AFTER SENDING")

	fmt.Println(<-ch)
}
