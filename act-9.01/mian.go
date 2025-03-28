package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	myId := uuid.New()
	q := quote.Glass()

	fmt.Println(myId)
	fmt.Println(q)
}
