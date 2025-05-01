package main

import (
	"fmt"
	"time"
)

func _sum(from, to int) int {
	res := 0
	for i := from; i <= to; i++ {
		res += i
	}
	return res
}

func mainws() {
	var s1, s2 int
	go func() {
		s1 = _sum(1, 100)
	}()
	s2 = _sum(1, 10)
	time.Sleep(time.Second)
	fmt.Println(s1, s2)
}
