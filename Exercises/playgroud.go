package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetMessage(in chan string) {
	for message := range in {
		fmt.Println(message)
	}
}

func main7() {
	ch := make(chan string)

	go GetMessage(ch)

	// ch <- "Hello Oussama Tailba"
	// ch <- "Hello Oussama Tailba 2"
	// ch <- "Hello Oussama Tailba 3"
	// ch <- "Hello Oussama Tailba 4"

	fmt.Printf("%s", "ousta6>")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Printf("%s", "ousta6>")
		line := scanner.Text()
		ch <- line
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
