package main

import (
	"fmt"
	"time"
)

func mainwew21() {
	current := time.Now()

	d, m, y, h, mi, s := current.Day(), current.Month(), current.Year(), current.Hour(), current.Minute(), current.Second()
	// 15:32:30 2023/10/17
	fmt.Printf("%d:%d:%d %d/%d/%d\n", h, mi, s, y, m, d)
}
