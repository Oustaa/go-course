package main

import "fmt"

type Game struct {
	name string
	age  int
}

func madsd3in() {

	myGame := Game{"Counter sctike: 2", 34}

	fmt.Println(myGame)

	fmt.Printf("%b\n", 10)
	fmt.Printf("%x\n", 10)
	fmt.Printf("float is %22.0f\n", 9.909)

	fmt.Printf("%T\n", myGame)
	fmt.Printf("%#v\n", myGame)

	g := Game{name: "Counter sctike: 2", age: 34}

	fmt.Print(g)
}

func (g Game) String() string {
	return fmt.Sprint("I am just a game")
}
