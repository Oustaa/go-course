package main

import "fmt"

func readThem(in, out chan string) {
	for i := range in {
		fmt.Print(i)
	}

	out <- "notification: We just done"
}

func main5() {
	in := make(chan string)
	out := make(chan string)

	go readThem(in, out)

	in <- "Oussama "
	in <- "Is "
	in <- "The "
	in <- "Fucking "
	in <- "Goat"
	in <- "!!\n"
	close(in)

	fmt.Println(<-out)
}
