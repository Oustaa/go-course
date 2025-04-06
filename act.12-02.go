package main

import (
	"fmt"
	"time"
)

func main212() {
	date := time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println("Date:", date.Format("2006-01-02"))
}
