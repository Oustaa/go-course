package main

import (
	_ "embed"
	"fmt"
)

//go:embed message.txt
var message string

func masaisn() {
	fmt.Println(message)
}
