package main

import "log"

func maisddfn() {
	ch := make(chan int)
	ch <- 1
	i := <-ch
	log.Println(i)
}
