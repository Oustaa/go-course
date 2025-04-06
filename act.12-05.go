package main

import (
	"fmt"
	"time"
)

func maiwe2n() {
	current := time.Now()
	nyTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	laTime, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The local current time is:", current.Format("2006-01-02 15:04:05"))
	fmt.Println("The time in New York is: ", current.In(nyTime).Format("2006-01-02 15:04:05"))
	fmt.Println("The time in Los Angeles is: ", current.In(laTime).Format("2006-01-02 15:04:05"))
}
