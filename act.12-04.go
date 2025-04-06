package main

import (
	"fmt"
	"time"
)

func maind332() {
	now := time.Now()

	_666FromNow := time.Duration(6*time.Hour + 6*time.Minute + 6*time.Second)

	future := now.Add(_666FromNow)

	fmt.Printf("6 hours, 6 minutes, and 6 seconds from now is: %s\n", future.Format(time.ANSIC))

}
