package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func sum(from int, to int, wg *sync.WaitGroup, result *int32) {
	for i := from; i <= to; i++ {
		atomic.AddInt32(result, int32(i))
		// *result += int32(i)

	}
	wg.Done()
}

func mawin() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}

	wg.Add(4)

	go sum(1, 25, wg, &s1)
	go sum(26, 50, wg, &s1)
	go sum(51, 75, wg, &s1)
	go sum(76, 100, wg, &s1)

	wg.Wait()
	log.Println(s1)
}
