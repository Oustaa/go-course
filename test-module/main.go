package main

import (
	"fmt"

	person "something.com/Person"
)

func main() {

	me := person.New("Oussama Tailba", 27, "Bab ghmat Syba", "Marrakech", "Morocco")

	fmt.Printf("%#v\n", me)

	me.Name = ""
}
