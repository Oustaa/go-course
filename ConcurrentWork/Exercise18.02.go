package main

import (
	"log"
	"sync"
)

func _sum_(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0

	for i := from; i <= to; i++ {
		*res += i
	}

	wg.Done()
}

func masin() {
	s1 := 0
	s2 := 0
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go _sum_(1, 100, wg, &s1)
	go _sum_(1, 1000, wg, &s2)
	wg.Wait()

	log.Println(s2)
	log.Println(s1)
}
