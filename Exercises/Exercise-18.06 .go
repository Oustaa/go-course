package main

import (
	"log"
	"sync"
	"time"
)

func _push(from, to int, out chan int, wg *sync.WaitGroup) {
	for i := from; i <= to; i++ {
		out <- i
		time.Sleep(time.Microsecond)
	}

	wg.Done()
}

func main2() {
	s1 := 0
	ch := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(4)

	go _push(1, 25, ch, &wg)
	go _push(26, 50, ch, &wg)
	go _push(51, 75, ch, &wg)
	go _push(76, 100, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		log.Println(i)
		s1 += i
	}
	log.Println(s1)
}
