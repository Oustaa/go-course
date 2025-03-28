package main

import (
	"fmt"

	"github.com/oustaa/printer"
)

func main() {
	msg := printer.PrintNewUUID()

	fmt.Println(msg)
}
