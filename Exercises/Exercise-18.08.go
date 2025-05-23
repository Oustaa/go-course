package main

import (
	"fmt"
	"log"
)

func worker(in chan int, out chan int, index int) {
	sum := 0
	for i := range in {
		fmt.Printf("worker %d got number %d\n", index, i)
		sum += i
	}
	out <- sum
}

func sum(workers, from, to int) int {
	out := make(chan int, workers)
	in := make(chan int, 20)
	for i := 0; i < workers; i++ {
		go worker(in, out, i)
	}

	for i := from; i <= to; i++ {
		in <- i
	}

	close(in)

	sum := 0
	for i := 0; i < workers; i++ {
		sum += <-out
	}

	close(out)
	return sum
}

func main4() {
	res := sum(20, 1, 100)
	log.Println(res)
}
