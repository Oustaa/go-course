package main

import "fmt"

type Person struct {
	name         string
	age          int
	is_he_or_she string
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 12
	p.is_he_or_she = "he"

	return &p
}

func add(a, b int16) int16 {
	return a + b

}

func main() {
	const p = 3.14
	result := add(12, 34)
	// var a, b int16 = 34, 55
	// var is_ok bool = false

	var s string
	fmt.Println(s)

	var e int
	fmt.Println(e)

	fmt.Printf("%d \n", result)

	fmt.Println("Hello I am Oussama Tailba")

	for i := range 3000 {
		fmt.Println(i)
	}

}
